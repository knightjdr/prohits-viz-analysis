package dotplot

import "fmt"

func writeBoundingBox(d *Dotplot, writeString func(string)) {
	if d.BoundingBox {
		str := fmt.Sprintf(
			"\t<rect fill=\"none\" y=\"%d\" x=\"%d\" width=\"%d\" height=\"%d\" stroke=\"#000000\" stroke-width=\"0.5\" />\n",
			d.TopMargin, d.LeftMargin, d.SvgWidth-d.LeftMargin, d.SvgHeight-d.TopMargin,
		)
		writeString(str)
	}
}
