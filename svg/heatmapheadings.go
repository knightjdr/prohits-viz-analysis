package svg

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// HeatmapHeadings adds column and row headings
func HeatmapHeadings(dims HDimensions, options map[string]interface{}) string {
	svg := make([]string, 0)

	// Add column label.
	xOffset := dims.leftMargin + ((dims.svgWidth - dims.leftMargin) / 2)
	text := fmt.Sprintf(
		"\t<text y=\"10\" x=\"%d\" font-size=\"12\""+
			" text-anchor=\"middle\">%s</text>\n",
		xOffset, options["colLabel"],
	)
	svg = append(svg, text)

	// Add row label.
	yOffset := dims.topMargin + ((dims.svgHeight - dims.topMargin) / 2)
	text = fmt.Sprintf(
		"\t<text y=\"%d\" x=\"10\" font-size=\"12\""+
			" text-anchor=\"middle\" transform=\"rotate(-90, 10, %d)\">%s</text>\n",
		yOffset, yOffset, options["rowLabel"],
	)
	svg = append(svg, text)
	return helper.StringConcat(svg)
}
