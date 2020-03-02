package hierarchical

import (
	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/normalize"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

const ignoreNodes = types.HclustIgnore

// HclustData contains clustering information.
type HclustData struct {
	Dendrogram         map[string][]hclust.SubCluster
	Distance           map[string][][]float64
	NormalizedDistance map[string][][]float64
	Tree               map[string]hclust.TreeLayout
	UnsortedNames      map[string][]string
}

func cluster(matrices *types.Matrices, settings types.Settings) HclustData {
	data := HclustData{
		Dendrogram:         make(map[string][]hclust.SubCluster, 2),
		Distance:           make(map[string][][]float64, 2),
		NormalizedDistance: make(map[string][][]float64, 2),
		Tree:               make(map[string]hclust.TreeLayout, 2),
		UnsortedNames: map[string][]string{
			"condition": matrices.Conditions,
			"readout":   matrices.Readouts,
		},
	}

	// Generate condition and readout distance matrices.
	data.Distance["condition"] = hclust.Distance(matrices.Abundance, settings.Distance, true)
	data.Distance["readout"] = hclust.Distance(matrices.Abundance, settings.Distance, false)

	var err error

	// Condition and readout clustering.
	data.Dendrogram["condition"], err = hclust.Cluster(data.Distance["condition"], settings.ClusteringMethod)
	log.CheckError(err, true)
	data.Dendrogram["readout"], err = hclust.Cluster(data.Distance["readout"], settings.ClusteringMethod)
	log.CheckError(err, true)

	// Optimize clustering.
	if settings.ClusteringOptimize {
		data.Dendrogram["condition"] = hclust.Optimize(data.Dendrogram["condition"], data.Distance["condition"], ignoreNodes)
		data.Dendrogram["readout"] = hclust.Optimize(data.Dendrogram["readout"], data.Distance["readout"], ignoreNodes)
	}

	// Create tree and get clustering order.
	data.Tree["condition"], err = hclust.Tree(data.Dendrogram["condition"], matrices.Conditions)
	log.CheckError(err, true)
	data.Tree["readout"], err = hclust.Tree(data.Dendrogram["readout"], matrices.Readouts)
	log.CheckError(err, true)

	// Normalize distance matrices to 1.
	data.NormalizedDistance["condition"] = normalize.Matrix(data.Distance["condition"])
	data.NormalizedDistance["readout"] = normalize.Matrix(data.Distance["readout"])

	return data
}
