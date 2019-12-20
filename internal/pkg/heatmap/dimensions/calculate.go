// Package dimensions calculates the dimensions for a heat map
package dimensions

// Heatmap contains heatmap plot dimensions.
type Heatmap struct {
	CellSize     int
	ColumnNumber int
	FontSize     int
	LeftMargin   int
	PlotHeight   int
	PlotWidth    int
	Ratio        float64
	RowNumber    int
	SvgHeight    int
	SvgWidth     int
	TopMargin    int
}

// Calculate heatmap/dotplot dimensions
func Calculate(matrix [][]float64, columns, rows []string, isMinimap bool) *Heatmap {
	dims := &Heatmap{}
	calculateRatio(dims, matrix)
	calculateCellSize(dims)

	if isMinimap {
		setMinimapDimensions(dims)
	} else {
		calculateMarginsAndSize(dims, columns, rows)
	}

	return dims
}
