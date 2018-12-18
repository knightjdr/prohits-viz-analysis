package svg

import (
	"fmt"
	"math"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Annotations adds annotations to an SVG
func Annotations(annotations typedef.Annotations, dims HeatmapDimensions, writeString func(string)) {
	if annotations.List == nil && len(annotations.List) == 0 {
		return
	}
	writeString(fmt.Sprintf("\t<g transform=\"translate(%d, %d)\">\n", dims.leftMargin, dims.topMargin))
	for _, annotation := range annotations.List {
		x := int(math.Round(annotation.X * float64(dims.plotWidth)))
		y := int(math.Round(annotation.Y * float64(dims.plotHeight)))
		text := fmt.Sprintf(
			"\t\t<text y=\"%d\" x=\"%d\" font-size=\"%d\""+
				" text-anchor=\"middle\">%s</text>\n",
			y, x, annotations.FontSize, annotation.Text,
		)
		writeString(text)
	}
	writeString("\t</g>\n")
}
