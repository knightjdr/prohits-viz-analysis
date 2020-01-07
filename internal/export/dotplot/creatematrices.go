package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/export/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func createMatrices(data *heatmap.Heatmap) *types.Matrices {
	matrices := &types.Matrices{
		Abundance: make([][]float64, len(data.RowOrder)),
		Ratio:     make([][]float64, len(data.RowOrder)),
		Score:     make([][]float64, len(data.RowOrder)),
	}

	for i, rowIndex := range data.RowOrder {
		matrices.Abundance[i] = make([]float64, len(data.ColumnOrder))
		matrices.Ratio[i] = make([]float64, len(data.ColumnOrder))
		matrices.Score[i] = make([]float64, len(data.ColumnOrder))
		for j, columnIndex := range data.ColumnOrder {
			matrices.Abundance[i][j] = data.RowDB[rowIndex].Data[columnIndex].Value
			matrices.Ratio[i][j] = data.RowDB[rowIndex].Data[columnIndex].Ratio
			matrices.Score[i][j] = data.RowDB[rowIndex].Data[columnIndex].Score
		}
	}

	return matrices
}
