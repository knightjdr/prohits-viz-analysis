package svg

import "fmt"

// BoundingBox draws a bounding box around a dotplot/heatmap
func BoundingBox(dims HDimensions) string {
	return fmt.Sprintf(
		"\t<rect fill=\"none\" y=\"%d\" x=\"%d\" width=\"%d\" height=\"%d\""+
			" stroke=\"#000000\" stroke-width=\"0.5\" />\n",
		dims.topMargin, dims.leftMargin, dims.svgWidth-dims.leftMargin, dims.svgHeight-dims.topMargin,
	)
}
