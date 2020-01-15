package nocluster

import (
	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/internal/analyze/dotplot/hierarchical"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/normalize"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
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
	sorted.Matrices.Score, _ = hclust.Sort(matrices.Score, matrices.Conditions, sorted.Matrices.Conditions, "column")
	sorted.Matrices.Score, _ = hclust.Sort(sorted.Matrices.Score, matrices.Readouts, sorted.Matrices.Readouts, "row")
	sorted.Matrices.Ratio = normalize.Matrix(sorted.Matrices.Abundance)

	if len(orderedData.NormalizedDistance["condition"]) > 0 {
		sorted.ConditionDist, _ = hclust.Sort(orderedData.NormalizedDistance["condition"], matrices.Conditions, sorted.Matrices.Conditions, "column")
		sorted.ConditionDist, _ = hclust.Sort(sorted.ConditionDist, matrices.Conditions, sorted.Matrices.Conditions, "row")
	}
	if len(orderedData.NormalizedDistance["readout"]) > 0 {
		sorted.ReadoutDist, _ = hclust.Sort(orderedData.NormalizedDistance["readout"], matrices.Readouts, sorted.Matrices.Readouts, "column")
		sorted.ReadoutDist, _ = hclust.Sort(sorted.ReadoutDist, matrices.Readouts, sorted.Matrices.Readouts, "row")
	}

	return sorted
}
