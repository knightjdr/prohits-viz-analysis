package heatmap

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
)

func writeCells(h *Heatmap, writeString func(string)) {
	colorGradient := createGradient(h)

	convertValueToIndex := float.GetRange(h.FillMin, h.FillMax, 0, 100)

	writeString(fmt.Sprintf("\t<g id=\"minimap\" transform=\"translate(%d, %d)\">\n", h.LeftMargin, h.TopMargin))

	for i, row := range h.Matrix {
		y := i * h.CellSize
		for j, value := range row {
			index := int(convertValueToIndex(value))
			fill := colorGradient[index].Hex

			cell := fmt.Sprintf(
				"\t\t<rect fill=\"%s\" y=\"%d\" x=\"%d\" width=\"%d\" height=\"%d\" />\n",
				fill, y, j*h.CellSize, h.CellSize, h.CellSize,
			)
			writeString(cell)
		}
	}
	writeString("\t</g>\n")
}
