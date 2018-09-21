package svg

import (
	"fmt"
	"math"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// HeatmapAnnotations adds annotations to an SVG
func HeatmapAnnotations(annotations typedef.Annotations, dims HDimensions) string {
	if annotations.List == nil && len(annotations.List) == 0 {
		return ""
	}
	svg := make([]string, 0)
	svg = append(svg, fmt.Sprintf("\t<g transform=\"translate(%d, %d)\">\n", dims.leftMargin, dims.topMargin))
	for _, annotation := range annotations.List {
		x := int(math.Round(annotation.X * float64(dims.plotWidth)))
		y := int(math.Round(annotation.Y * float64(dims.plotHeight)))
		text := fmt.Sprintf(
			"\t\t<text y=\"%d\" x=\"%d\" font-size=\"%d\""+
				" text-anchor=\"middle\">%s</text>\n",
			y, x, annotations.FontSize, annotation.Text,
		)
		svg = append(svg, text)

	}
	svg = append(svg, "\t</g>\n")
	return helper.StringConcat(svg)
}
