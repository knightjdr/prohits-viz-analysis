package biclustering

import (
	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/internal/analyze/dotplot/hierarchical"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func order(matrices *types.Matrices, minAbundance float64) hierarchical.HclustData {
	singles := generateMatrix(matrices, minAbundance)

	order := nestedClustering()

	orderedData := hierarchical.HclustData{
		Dendrogram: map[string][]hclust.SubCluster{
			"condition": []hclust.SubCluster{},
			"readout":   []hclust.SubCluster{},
		},
		Distance: map[string][][]float64{
			"condition": [][]float64{},
			"readout":   [][]float64{},
		},
		NormalizedDistance: map[string][][]float64{
			"condition": [][]float64{},
			"readout":   [][]float64{},
		},
		Tree: map[string]hclust.TreeLayout{
			"condition": hclust.TreeLayout{
				Order: order["condition"],
			},
			"readout": hclust.TreeLayout{
				Order: append(order["readout"], singles...),
			},
		},
		UnsortedNames: map[string][]string{
			"condition": matrices.Conditions,
			"readout":   matrices.Readouts,
		},
	}

	return orderedData
}
