package heatmap

import (
	"fmt"
	"math"
)

func writeAnnotations(h *Heatmap, writeString func(string)) {
	if h.Annotations.List == nil || len(h.Annotations.List) == 0 {
		return
	}

	writeString(fmt.Sprintf("\t<g transform=\"translate(%d, %d)\">\n", h.LeftMargin, h.TopMargin))
	for _, annotation := range h.Annotations.List {
		x := int(math.Round(annotation.Position.X * float64(h.PlotWidth)))
		y := int(math.Round(annotation.Position.Y * float64(h.PlotHeight)))
		text := fmt.Sprintf(
			"\t\t<text y=\"%d\" x=\"%d\" font-size=\"%d\" text-anchor=\"middle\">%s</text>\n",
			y, x, h.Annotations.FontSize, annotation.Text,
		)
		writeString(text)
	}
	writeString("\t</g>\n")
}
