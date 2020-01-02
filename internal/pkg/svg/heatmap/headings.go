package heatmap

import "fmt"

func writeHeadings(h *Heatmap, writeString func(string)) {
	if h.XLabel != "" {
		xOffset := h.LeftMargin + ((h.SvgWidth - h.LeftMargin) / 2)
		text := fmt.Sprintf(
			"\t<text y=\"10\" x=\"%d\" font-size=\"12\" text-anchor=\"middle\">%s</text>\n",
			xOffset, h.XLabel,
		)
		writeString(text)
	}

	if h.YLabel != "" {
		yOffset := h.TopMargin + ((h.SvgHeight - h.TopMargin) / 2)
		text := fmt.Sprintf(
			"\t<text y=\"%d\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, %d)\">%s</text>\n",
			yOffset, yOffset, h.YLabel,
		)
		writeString(text)
	}
}
