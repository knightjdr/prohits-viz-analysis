package parse

import (
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// FormatMatrix converts json data slice to matrices for svg creation.
func FormatMatrix(data *Data) *typedef.Matrices {
	// Define matrix dimensions.
	colNum := len(data.Rows[0].Data)
	rowNum := len(data.Rows)

	// Init matrices.
	heatmap := typedef.Matrices{
		Abundance: make([][]float64, rowNum),
		Ratio:     make([][]float64, rowNum),
		Score:     make([][]float64, rowNum),
	}

	// Define matrices.
	for i, row := range data.Rows {
		heatmap.Abundance[i] = make([]float64, colNum)
		heatmap.Ratio[i] = make([]float64, colNum)
		heatmap.Score[i] = make([]float64, colNum)
		for j, cell := range row.Data {
			heatmap.Abundance[i][j] = cell.Value
			heatmap.Ratio[i][j] = cell.Ratio
			heatmap.Score[i][j] = cell.Score
		}
	}
	return &heatmap
}
