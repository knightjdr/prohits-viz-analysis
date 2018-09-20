package svg

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// HeatmapColumnNames writes column names for an SVG
func HeatmapColumnNames(dims HDimensions, columns []string) string {
	svg := make([]string, 0)

	xOffset := dims.fontSize / 2
	yOffset := dims.topMargin - 2
	svg = append(svg, fmt.Sprintf("\t<g transform=\"translate(%d)\">\n", dims.leftMargin))
	for i, colName := range columns {
		xPos := (i * dims.cellSize) + xOffset
		text := fmt.Sprintf(
			"\t\t<text y=\"%d\" x=\"%d\" font-size=\"%d\""+
				" text-anchor=\"end\" transform=\"rotate(90, %d, %d)\">%s</text>\n",
			yOffset, xPos, dims.fontSize, xPos, yOffset, colName,
		)
		svg = append(svg, text)
	}
	svg = append(svg, "\t</g>\n")
	return helper.StringConcat(svg)
}
