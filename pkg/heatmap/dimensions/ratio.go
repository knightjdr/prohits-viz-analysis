package dimensions

import (
	"math"
)

func calculateRatio(dims *Heatmap, matrix [][]float64) {
	dims.ColumnNumber = len(matrix[0])
	colRatio := float64(dims.ColumnNumber*idealCellSize) / float64(maxImageWidth)
	dims.RowNumber = len(matrix)
	rowRatio := float64(dims.RowNumber*idealCellSize) / float64(maxImageHeight)

	dims.Ratio = float64(1)
	if colRatio > 1 || rowRatio > 1 {
		dims.Ratio = math.Max(colRatio, rowRatio)
		dims.Ratio = 1 / dims.Ratio
	}
	if dims.Ratio < minRatio {
		dims.Ratio = minRatio
	}
}
