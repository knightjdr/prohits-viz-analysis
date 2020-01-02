package heatmap

import "fmt"

func writeRowNames(h *Heatmap, writeString func(string)) {
	xOffset := h.LeftMargin - 2
	yOffset := (h.CellSize + h.FontSize - 2) / 2

	writeString(fmt.Sprintf("\t<g transform=\"translate(0, %d)\">\n", h.TopMargin))

	for i, name := range h.Rows {
		y := (i * h.CellSize) + yOffset
		text := fmt.Sprintf(
			"\t\t<text y=\"%d\" x=\"%d\" font-size=\"%d\" text-anchor=\"end\">%s</text>\n",
			y, xOffset, h.FontSize, name,
		)
		writeString(text)
	}
	writeString("\t</g>\n")
}
