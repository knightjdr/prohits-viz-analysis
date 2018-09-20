package svg

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// HeatmapRowNames writes rows names for an SVG
func HeatmapRowNames(dims HDimensions, rows []string) string {
	svg := make([]string, 0)

	xOffset := dims.leftMargin - 2
	yOffset := (dims.cellSize + dims.fontSize - 2) / 2
	svg = append(svg, fmt.Sprintf("\t<g transform=\"translate(0, %d)\">\n", dims.topMargin))
	for i, rowName := range rows {
		yPos := (i * dims.cellSize) + yOffset
		text := fmt.Sprintf(
			"\t\t<text y=\"%d\" x=\"%d\" font-size=\"%d\" text-anchor=\"end\">%s</text>\n",
			yPos, xOffset, dims.fontSize, rowName,
		)
		svg = append(svg, text)
	}
	svg = append(svg, "\t</g>\n")
	return helper.StringConcat(svg)
}
