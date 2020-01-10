package hierarchical

import (
	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/normalize"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

const ignoreNodes = 250000

type hclustData struct {
	clustering         map[string][]hclust.SubCluster
	distance           map[string][][]float64
	normalizedDistance map[string][][]float64
	tree               map[string]hclust.TreeLayout
}

func cluster(matrices *types.Matrices, settings types.Settings) hclustData {
	data := hclustData{
		clustering:         make(map[string][]hclust.SubCluster, 2),
		distance:           make(map[string][][]float64, 2),
		normalizedDistance: make(map[string][][]float64, 2),
		tree:               make(map[string]hclust.TreeLayout, 2),
	}

	// Generate condition and readout distance matrices.
	data.distance["condition"] = hclust.Distance(matrices.Abundance, settings.Distance, true)
	data.distance["readout"] = hclust.Distance(matrices.Abundance, settings.Distance, false)

	var err error

	// Condition and readout clustering.
	data.clustering["condition"], err = hclust.Cluster(data.distance["condition"], settings.ClusteringMethod)
	log.CheckError(err, true)
	data.clustering["readout"], err = hclust.Cluster(data.distance["readout"], settings.ClusteringMethod)
	log.CheckError(err, true)

	// Optimize clustering.
	if settings.ClusteringOptimize {
		data.clustering["condition"] = hclust.Optimize(data.clustering["condition"], data.distance["condition"], ignoreNodes)
		data.clustering["readout"] = hclust.Optimize(data.clustering["readout"], data.distance["readout"], ignoreNodes)
	}

	// Create tree and get clustering order.
	data.tree["condition"], err = hclust.Tree(data.clustering["condition"], matrices.Conditions)
	log.CheckError(err, true)
	data.tree["readout"], err = hclust.Tree(data.clustering["readout"], matrices.Readouts)
	log.CheckError(err, true)

	// Normalize distance matrices to 1.
	data.normalizedDistance["condition"] = normalize.Matrix(data.distance["condition"])
	data.normalizedDistance["readout"] = normalize.Matrix(data.distance["readout"])

	return data
}
