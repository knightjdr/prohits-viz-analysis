package svg

import (
	"fmt"
)

// RowNames writes rows names for an SVG
func RowNames(dims HeatmapDimensions, rows []string, writeString func(string)) {
	xOffset := dims.leftMargin - 2
	yOffset := (dims.cellSize + dims.fontSize - 2) / 2
	writeString(fmt.Sprintf("\t<g transform=\"translate(0, %d)\">\n", dims.topMargin))
	for i, rowName := range rows {
		yPos := (i * dims.cellSize) + yOffset
		text := fmt.Sprintf(
			"\t\t<text y=\"%d\" x=\"%d\" font-size=\"%d\" text-anchor=\"end\">%s</text>\n",
			yPos, xOffset, dims.fontSize, rowName,
		)
		writeString(text)
	}
	writeString("\t</g>\n")
	return
}
