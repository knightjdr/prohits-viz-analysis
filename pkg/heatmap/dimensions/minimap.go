package dimensions

func setMinimapDimensions(dims *Heatmap) {
	dims.SvgHeight = dims.RowNumber * dims.CellSize
	dims.SvgWidth = dims.ColumnNumber * dims.CellSize
	dims.PlotHeight = dims.SvgHeight
	dims.PlotWidth = dims.SvgWidth
}
