package svg

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// HeatmapMarkers add rectangular boxes to and SVG
func HeatmapMarkers(markers typedef.Markers, dims HDimensions) string {
	if markers.List == nil && len(markers.List) == 0 {
		return ""
	}
	svg := make([]string, 0)
	svg = append(svg, fmt.Sprintf("\t<g transform=\"translate(%d, %d)\">\n", dims.leftMargin, dims.topMargin))
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
		svg = append(svg, rect)

	}
	svg = append(svg, "\t</g>\n")
	return helper.StringConcat(svg)
}