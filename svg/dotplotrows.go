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
	parameters DParameters,
	options map[string]interface{},
) string {
	svg := make([]string, 0)

	// Get color gradients.
	edgeColorGradient := colorGradient(options["edgeColor"].(string), 101, false)
	fillColorGradient := colorGradient(options["fillColor"].(string), 101, options["invertColor"].(bool))

	// Get function for determining score edge color.
	edgeColorFunc := ScoreColorFunc(options["scoreType"].(string), options["primary"].(float64), options["secondary"].(float64), 100)

	// Write rows.
	svg = append(svg, fmt.Sprintf("\t<g id=\"minimap\" transform=\"translate(%d, %d)\">\n", dims.leftMargin, dims.topMargin))
	for i, row := range matrix {
		// Set x position.
		iPos := (i * dims.cellSize) + parameters.cellSizeHalf

		// Draw dots.
		for j, value := range row {
			if value > 0 {
				// Get fill color.
				var fill string
				if value > options["abundanceCap"].(float64) {
					fill = fillColorGradient[100]
				} else {
					index := int(math.Round(value / options["abundanceCap"].(float64) * 100))
					fill = fillColorGradient[index]
				}

				// Edge color
				edgeColorIndex := edgeColorFunc(scores[i][j])
				edgeColor := edgeColorGradient[edgeColorIndex]

				// Get circle radius.
				radius := helper.Round(ratios[i][j]*parameters.maxRadius, 0.01)

				// Draw circle.
				circle := fmt.Sprintf(
					"\t\t<circle fill=\"%s\" cy=\"%d\" cx=\"%d\" r=\"%f\""+
						" stroke=\"%s\" stroke-width=\"%f\"/>\n",
					fill, iPos, (j*dims.cellSize)+parameters.cellSizeHalf, radius, edgeColor, parameters.edgeWidth,
				)
				svg = append(svg, circle)
			}
		}
	}
	svg = append(svg, "\t</g>\n")
	return helper.StringConcat(svg)
}
