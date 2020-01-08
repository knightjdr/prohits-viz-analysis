package sync

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/matrix/frontend"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func createMatrices(data *minimap) *types.Matrices {

	if data.ImageType == "heatmap" {
		return &types.Matrices{
			Abundance: frontend.CreateHeatmapMatrix(data.RowDB, map[string][]int{"columns": data.ColumnOrder, "rows": data.RowOrder}),
		}
	}

	return frontend.CreateDotplotMatrices(data.RowDB, map[string][]int{"columns": data.ColumnOrder, "rows": data.RowOrder})
}
