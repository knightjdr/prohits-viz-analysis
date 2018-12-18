package svg

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// HeatmapRows generates the heatmap element for an SVG.
func HeatmapRows(
	matrix [][]float64,
	dims HeatmapDimensions,
	parameters typedef.Parameters,
	writeString func(string),
) {
	// Get color gradient.
	colorGradient := colorGradient(parameters.FillColor, 101, parameters.InvertColor)

	// Get range function
	getIndex := helper.SetRange(parameters.MinAbundance, parameters.AbundanceCap, 0, 100)

	// Write rows.
	writeString(fmt.Sprintf("\t<g id=\"minimap\" transform=\"translate(%d, %d)\">\n", dims.leftMargin, dims.topMargin))
	for i, row := range matrix {
		iPos := i * dims.cellSize
		for j, value := range row {
			var fill string
			if value > parameters.AbundanceCap {
				fill = colorGradient[100]
			} else {
				index := int(getIndex(value))
				fill = colorGradient[index]
			}
			rect := fmt.Sprintf(
				"\t\t<rect fill=\"%s\" y=\"%d\" x=\"%d\" width=\"%d\" height=\"%d\" />\n",
				fill, iPos, j*dims.cellSize, dims.cellSize, dims.cellSize,
			)
			writeString(rect)
		}
	}
	writeString("\t</g>\n")
}
