package svg

import (
	"fmt"
	"math"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// DotplotRows generates the dotplot element for an SVG.
func DotplotRows(
	matrix, ratios, scores [][]float64,
	dims HDimensions,
	params DParameters,
	options map[string]interface{},
) string {
	svg := make([]string, 0)

	// Get color gradients.
	edgeColorGradient := ColorGradient(options["edgeColor"].(string), 101, false)
	fillColorGradient := ColorGradient(options["fillColor"].(string), 101, options["invert"].(bool))

	// Get function for determining score edge color.
	edgeColorFunc := ScoreColorFunc(options["scoreType"].(string), options["primary"].(float64), options["secondary"].(float64), 100)

	// Write rows.
	svg = append(svg, fmt.Sprintf("\t<g id=\"minimap\" transform=\"translate(%d, %d)\">\n", dims.leftMargin, dims.topMargin))
	for i, row := range matrix {
		// Set x position.
		iPos := (i * dims.cellSize) + params.cellSizeHalf

		// Draw dots.
		for j, value := range row {
			if value > 0 {
				// Get fill color.
				var fill string
				if value > options["maximumAbundance"].(float64) {
					fill = fillColorGradient[100]
				} else {
					index := int(math.Round(value / options["maximumAbundance"].(float64) * 100))
					fill = fillColorGradient[index]
				}

				// Edge color
				edgeColorIndex := edgeColorFunc(scores[i][j])
				edgeColor := edgeColorGradient[edgeColorIndex]

				// Get circle radius.
				radius := helper.Round(ratios[i][j]*params.maxRadius, 0.01)

				// Draw circle.
				circle := fmt.Sprintf(
					"\t\t<circle fill=\"%s\" cy=\"%d\" cx=\"%d\" r=\"%f\""+
						" stroke=\"%s\" stroke-width=\"%f\"/>\n",
					fill, iPos, (j*dims.cellSize)+params.cellSizeHalf, radius, edgeColor, params.edgeWidth,
				)
				svg = append(svg, circle)
			}
		}
	}
	svg = append(svg, "\t</g>\n")
	return helper.StringConcat(svg)
}
