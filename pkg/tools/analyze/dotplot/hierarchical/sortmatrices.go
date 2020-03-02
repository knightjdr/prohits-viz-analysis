package hierarchical

import (
	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/pkg/normalize"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// SortedData from clustering.
type SortedData struct {
	ConditionDist [][]float64
	Matrices      *types.Matrices
	ReadoutDist   [][]float64
}

func sortMatrices(matrices *types.Matrices, clusteredData HclustData) *SortedData {
	sorted := &SortedData{
		Matrices: &types.Matrices{
			Conditions: clusteredData.Tree["condition"].Order,
			Readouts:   clusteredData.Tree["readout"].Order,
		},
	}

	sorted.Matrices.Abundance, _ = hclust.Sort(matrices.Abundance, matrices.Conditions, sorted.Matrices.Conditions, "column")
	sorted.Matrices.Abundance, _ = hclust.Sort(sorted.Matrices.Abundance, matrices.Readouts, sorted.Matrices.Readouts, "row")
	sorted.ConditionDist, _ = hclust.Sort(clusteredData.NormalizedDistance["condition"], matrices.Conditions, sorted.Matrices.Conditions, "column")
	sorted.ConditionDist, _ = hclust.Sort(sorted.ConditionDist, matrices.Conditions, sorted.Matrices.Conditions, "row")
	sorted.ReadoutDist, _ = hclust.Sort(clusteredData.NormalizedDistance["readout"], matrices.Readouts, sorted.Matrices.Readouts, "column")
	sorted.ReadoutDist, _ = hclust.Sort(sorted.ReadoutDist, matrices.Readouts, sorted.Matrices.Readouts, "row")
	sorted.Matrices.Score, _ = hclust.Sort(matrices.Score, matrices.Conditions, sorted.Matrices.Conditions, "column")
	sorted.Matrices.Score, _ = hclust.Sort(sorted.Matrices.Score, matrices.Readouts, sorted.Matrices.Readouts, "row")

	sorted.Matrices.Ratio = normalize.Matrix(sorted.Matrices.Abundance)

	return sorted

}
