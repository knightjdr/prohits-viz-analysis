package svg

import (
	"fmt"
	"math"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// scoreColorFunc returns a function for determining the gradient index to use
// for the score color.
func scoreColorFunc(scoretype string, primary, secondary float64, numColors int) func(score float64) int {
	if scoretype == "gte" {
		return func(score float64) int {
			if score >= primary {
				return numColors
			} else if score < primary && score >= secondary {
				return numColors / 2
			}
			return numColors / 4
		}
	}
	return func(score float64) int {
		if score <= primary {
			return numColors
		} else if score > primary && score <= secondary {
			return numColors / 2
		}
		return numColors / 4
	}
}

// DotplotRows generates the dotplot element for an SVG.
func DotplotRows(
	matrix, ratios, scores [][]float64,
	dims HDimensions,
	dotplotparameters DParameters,
	parameters typedef.Parameters,
) string {
	svg := make([]string, 0)

	// Get color gradients.
	edgeColorGradient := colorGradient(parameters.EdgeColor, 101, false)
	fillColorGradient := colorGradient(parameters.FillColor, 101, parameters.InvertColor)

	// Get function for determining score edge color.
	edgeColorFunc := scoreColorFunc(parameters.ScoreType, parameters.PrimaryFilter, parameters.SecondaryFilter, 100)

	// Write rows.
	svg = append(svg, fmt.Sprintf("\t<g id=\"minimap\" transform=\"translate(%d, %d)\">\n", dims.leftMargin, dims.topMargin))
	for i, row := range matrix {
		// Set x position.
		iPos := (i * dims.cellSize) + dotplotparameters.cellSizeHalf

		// Draw dots.
		for j, value := range row {
			if value > 0 {
				// Get fill color.
				var fill string
				if value > parameters.AbundanceCap {
					fill = fillColorGradient[100]
				} else {
					index := int(math.Round(value / parameters.AbundanceCap * 100))
					fill = fillColorGradient[index]
				}

				// Edge color
				edgeColorIndex := edgeColorFunc(scores[i][j])
				edgeColor := edgeColorGradient[edgeColorIndex]

				// Get circle radius.
				radius := helper.Round(ratios[i][j]*dotplotparameters.maxRadius, 0.01)

				// Draw circle.
				circle := fmt.Sprintf(
					"\t\t<circle fill=\"%s\" cy=\"%d\" cx=\"%d\" r=\"%f\""+
						" stroke=\"%s\" stroke-width=\"%f\"/>\n",
					fill, iPos, (j*dims.cellSize)+dotplotparameters.cellSizeHalf, radius, edgeColor, dotplotparameters.edgeWidth,
				)
				svg = append(svg, circle)
			}
		}
	}
	svg = append(svg, "\t</g>\n")
	return helper.StringConcat(svg)
}
