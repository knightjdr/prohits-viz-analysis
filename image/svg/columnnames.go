package svg

import (
	"fmt"
)

// ColumnNames writes column names for an SVG
func ColumnNames(dims HeatmapDimensions, columns []string, writeString func(string)) {
	xOffset := dims.fontSize / 2
	yOffset := dims.topMargin - 2
	writeString(fmt.Sprintf("\t<g transform=\"translate(%d)\">\n", dims.leftMargin))
	for i, colName := range columns {
		xPos := (i * dims.cellSize) + xOffset
		text := fmt.Sprintf(
			"\t\t<text y=\"%d\" x=\"%d\" font-size=\"%d\""+
				" text-anchor=\"end\" transform=\"rotate(90, %d, %d)\">%s</text>\n",
			yOffset, xPos, dims.fontSize, xPos, yOffset, colName,
		)
		writeString(text)
	}
	writeString("\t</g>\n")
	return
}
