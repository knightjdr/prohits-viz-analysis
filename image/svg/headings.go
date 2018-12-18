package svg

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Headings adds column and row headings
func Headings(parameters typedef.Parameters, dims HeatmapDimensions, writeString func(string)) {
	if parameters.XLabel == "" && parameters.YLabel == "" {
		return
	}

	// Add column label.
	if parameters.XLabel != "" {
		xOffset := dims.leftMargin + ((dims.svgWidth - dims.leftMargin) / 2)
		text := fmt.Sprintf(
			"\t<text y=\"10\" x=\"%d\" font-size=\"12\""+
				" text-anchor=\"middle\">%s</text>\n",
			xOffset, parameters.XLabel,
		)
		writeString(text)
	}

	// Add row label.
	if parameters.YLabel != "" {
		yOffset := dims.topMargin + ((dims.svgHeight - dims.topMargin) / 2)
		text := fmt.Sprintf(
			"\t<text y=\"%d\" x=\"10\" font-size=\"12\""+
				" text-anchor=\"middle\" transform=\"rotate(-90, 10, %d)\">%s</text>\n",
			yOffset, yOffset, parameters.YLabel,
		)
		writeString(text)
	}
}
