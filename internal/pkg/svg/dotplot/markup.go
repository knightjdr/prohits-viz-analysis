package dotplot

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/heatmap"

func writeMarkup(d *Dotplot, writeString func(string)) {
	h := heatmap.Initialize()
	h.Annotations = d.Annotations
	h.CellSize = d.CellSize
	h.LeftMargin = d.LeftMargin
	h.Markers = d.Markers
	h.PlotHeight = d.PlotHeight
	h.PlotWidth = d.PlotWidth
	h.SvgHeight = d.SvgHeight
	h.SvgWidth = d.SvgWidth
	h.TopMargin = d.TopMargin
	h.XLabel = d.XLabel
	h.YLabel = d.YLabel
	heatmap.WriteMarkup(h, writeString)
}
