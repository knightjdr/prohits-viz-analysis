package dotplot

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/heatmap"

func writeLabels(d *Dotplot, writeString func(string)) {
	h := heatmap.Initialize()
	h.CellSize = d.CellSize
	h.Columns = d.Columns
	h.FontSize = d.FontSize
	h.LeftMargin = d.LeftMargin
	h.Rows = d.Rows
	h.TopMargin = d.TopMargin
	heatmap.WriteLabels(h, writeString)
}
