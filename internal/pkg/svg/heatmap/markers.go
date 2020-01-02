package heatmap

import (
	"fmt"
)

func writeMarkers(h *Heatmap, writeString func(string)) {
	if h.Markers.List == nil || len(h.Markers.List) == 0 {
		return
	}

	writeString(fmt.Sprintf("\t<g transform=\"translate(%d, %d)\">\n", h.LeftMargin, h.TopMargin))
	for _, marker := range h.Markers.List {
		height := marker.Height * h.CellSize
		width := marker.Width * h.CellSize
		x := marker.X * h.CellSize
		y := marker.Y * h.CellSize

		rect := fmt.Sprintf(
			"\t\t<rect y=\"%d\" x=\"%d\" width=\"%d\" height=\"%d\""+
				" stroke=\"%s\" stroke-width=\"1\" fill=\"none\"/>\n",
			y, x, width, height, h.Markers.Color,
		)
		writeString(rect)
	}
	writeString("\t</g>\n")
}
