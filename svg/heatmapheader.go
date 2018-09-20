package svg

import "fmt"

// HeatmapHeader creates the beginning of an SVG.
func HeatmapHeader(dims HDimensions) string {
	return fmt.Sprintf(
		"<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\""+
			" xml:space=\"preserve\" width=\"%d\" height=\"%d\" viewBox=\"0 0 %d %d\">\n",
		dims.svgWidth, dims.svgHeight, dims.svgWidth, dims.svgHeight,
	)
}
