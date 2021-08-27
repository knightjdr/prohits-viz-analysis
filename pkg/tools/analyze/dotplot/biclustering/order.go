package biclustering

import (
	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/dotplot/hierarchical"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func order(matrices *types.Matrices, minAbundance float64) hierarchical.HclustData {
	singles := generateMatrix(matrices, minAbundance)

	order := nestedClustering()

	orderedData := hierarchical.HclustData{
		Dendrogram: map[string][]hclust.SubCluster{
			"condition": {},
			"readout":   {},
		},
		Distance: map[string][][]float64{
			"condition": {},
			"readout":   {},
		},
		NormalizedDistance: map[string][][]float64{
			"condition": {},
			"readout":   {},
		},
		Tree: map[string]hclust.TreeLayout{
			"condition": {
				Order: order["condition"],
			},
			"readout": {
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
