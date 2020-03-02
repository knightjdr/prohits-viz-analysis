package correlation

import (
	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

const ignoreNodes = types.HclustIgnore

func cluster(corrData *correlationData, settings types.Settings) {
	var err error

	distance := hclust.Distance(corrData.matrix, settings.Distance, false)
	corrData.dendrogram, err = hclust.Cluster(distance, settings.ClusteringMethod)
	log.CheckError(err, true)

	if settings.ClusteringOptimize {
		corrData.dendrogram = hclust.Optimize(corrData.dendrogram, distance, ignoreNodes)
	}

	corrData.tree, err = hclust.Tree(corrData.dendrogram, corrData.labels)
	log.CheckError(err, true)

	corrData.matrix, _ = hclust.Sort(corrData.matrix, corrData.labels, corrData.tree.Order, "column")
	corrData.matrix, _ = hclust.Sort(corrData.matrix, corrData.labels, corrData.tree.Order, "row")

	corrData.sortedLabels = corrData.tree.Order
}
