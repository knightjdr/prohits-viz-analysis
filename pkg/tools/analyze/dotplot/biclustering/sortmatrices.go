package biclustering

import (
	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/dotplot/hierarchical"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func sortMatrices(matrices *types.Matrices, orderedData hierarchical.HclustData) *hierarchical.SortedData {
	sorted := &hierarchical.SortedData{
		Matrices: &types.Matrices{
			Conditions: orderedData.Tree["condition"].Order,
			Readouts:   orderedData.Tree["readout"].Order,
		},
	}

	sorted.Matrices.Abundance, _ = hclust.Sort(matrices.Abundance, matrices.Conditions, sorted.Matrices.Conditions, "column")
	sorted.Matrices.Abundance, _ = hclust.Sort(sorted.Matrices.Abundance, matrices.Readouts, sorted.Matrices.Readouts, "row")
	sorted.Matrices.Ratio, _ = hclust.Sort(matrices.Ratio, matrices.Conditions, sorted.Matrices.Conditions, "column")
	sorted.Matrices.Ratio, _ = hclust.Sort(sorted.Matrices.Ratio, matrices.Readouts, sorted.Matrices.Readouts, "row")
	sorted.Matrices.Score, _ = hclust.Sort(matrices.Score, matrices.Conditions, sorted.Matrices.Conditions, "column")
	sorted.Matrices.Score, _ = hclust.Sort(sorted.Matrices.Score, matrices.Readouts, sorted.Matrices.Readouts, "row")

	return sorted
}
