package svg

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Markers add rectangular boxes to and SVG
func Markers(markers typedef.Markers, dims HeatmapDimensions, writeString func(string)) {
	if markers.List == nil && len(markers.List) == 0 {
		return
	}
	writeString(fmt.Sprintf("\t<g transform=\"translate(%d, %d)\">\n", dims.leftMargin, dims.topMargin))
	for _, marker := range markers.List {
		height := marker.Height * dims.cellSize
		width := marker.Width * dims.cellSize
		x := marker.X * dims.cellSize
		y := marker.Y * dims.cellSize
		rect := fmt.Sprintf(
			"\t\t<rect y=\"%d\" x=\"%d\" width=\"%d\" height=\"%d\""+
				" stroke=\"%s\" stroke-width=\"1\" fill=\"none\"/>\n",
			y, x, width, height, markers.Color,
		)
		writeString(rect)
	}
	writeString("\t</g>\n")
}
