package dotplot

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/heatmap"

func writeHeader(d *Dotplot, writeString func(string)) {
	h := heatmap.Initialize()
	h.SvgHeight = d.SvgHeight
	h.SvgWidth = d.SvgWidth
	heatmap.WriteHeader(h, writeString)
}
