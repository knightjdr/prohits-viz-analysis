package heatmap

import (
	"fmt"
)

func writeColumnNames(h *Heatmap, writeString func(string)) {
	xOffset := h.FontSize / 2
	yOffset := h.TopMargin - 2

	writeString(fmt.Sprintf("\t<g transform=\"translate(%d)\">\n", h.LeftMargin))

	for i, name := range h.Columns {
		x := (i * h.CellSize) + xOffset
		text := fmt.Sprintf(
			"\t\t<text y=\"%d\" x=\"%d\" font-size=\"%d\" text-anchor=\"end\" transform=\"rotate(90, %d, %d)\">%s</text>\n",
			yOffset, x, h.FontSize, x, yOffset, name,
		)
		writeString(text)
	}
	writeString("\t</g>\n")
}
