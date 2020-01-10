package hierarchical

import (
	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/normalize"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

type sortedData struct {
	conditionDist [][]float64
	matrices      *types.Matrices
	readoutDist   [][]float64
}

func sortMatrices(matrices *types.Matrices, clusteredData hclustData) *sortedData {
	sorted := &sortedData{
		matrices: &types.Matrices{
			Conditions: clusteredData.tree["condition"].Order,
			Readouts:   clusteredData.tree["readout"].Order,
		},
	}

	sorted.matrices.Abundance, _ = hclust.Sort(matrices.Abundance, matrices.Conditions, sorted.matrices.Conditions, "column")
	sorted.matrices.Abundance, _ = hclust.Sort(sorted.matrices.Abundance, matrices.Readouts, sorted.matrices.Readouts, "row")
	sorted.conditionDist, _ = hclust.Sort(clusteredData.normalizedDistance["condition"], matrices.Conditions, sorted.matrices.Conditions, "column")
	sorted.conditionDist, _ = hclust.Sort(sorted.conditionDist, matrices.Conditions, sorted.matrices.Conditions, "row")
	sorted.readoutDist, _ = hclust.Sort(clusteredData.normalizedDistance["readout"], matrices.Readouts, sorted.matrices.Readouts, "column")
	sorted.readoutDist, _ = hclust.Sort(sorted.readoutDist, matrices.Readouts, sorted.matrices.Readouts, "row")
	sorted.matrices.Score, _ = hclust.Sort(matrices.Score, matrices.Conditions, sorted.matrices.Conditions, "column")
	sorted.matrices.Score, _ = hclust.Sort(sorted.matrices.Score, matrices.Readouts, sorted.matrices.Readouts, "row")

	sorted.matrices.Ratio = normalize.Matrix(sorted.matrices.Abundance)

	return sorted

}
