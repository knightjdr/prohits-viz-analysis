package svg

import (
	"fmt"
	"math"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// HeatmapRows generates the heatmap element for an SVG.
func HeatmapRows(matrix [][]float64, dims HDimensions, parameters typedef.Parameters) string {
	svg := make([]string, 0)

	// Get color gradient.
	colorGradient := colorGradient(parameters.FillColor, 101, parameters.InvertColor)

	// Write rows.
	svg = append(svg, fmt.Sprintf("\t<g id=\"minimap\" transform=\"translate(%d, %d)\">\n", dims.leftMargin, dims.topMargin))
	for i, row := range matrix {
		iPos := i * dims.cellSize
		for j, value := range row {
			var fill string
			if value > parameters.AbundanceCap {
				fill = colorGradient[100]
			} else {
				index := int(math.Round(value / parameters.AbundanceCap * 100))
				fill = colorGradient[index]
			}
			rect := fmt.Sprintf(
				"\t\t<rect fill=\"%s\" y=\"%d\" x=\"%d\" width=\"%d\" height=\"%d\" />\n",
				fill, iPos, j*dims.cellSize, dims.cellSize, dims.cellSize,
			)
			svg = append(svg, rect)
		}
	}
	svg = append(svg, "\t</g>\n")
	return helper.StringConcat(svg)
}
