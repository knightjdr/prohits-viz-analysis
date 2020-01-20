package correlation

import (
	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
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

	tree, err := hclust.Tree(corrData.dendrogram, corrData.labels)
	log.CheckError(err, true)

	corrData.matrix, _ = hclust.Sort(corrData.matrix, corrData.labels, tree.Order, "column")
	corrData.matrix, _ = hclust.Sort(corrData.matrix, corrData.labels, tree.Order, "row")

	corrData.sortedLabels = tree.Order
}
