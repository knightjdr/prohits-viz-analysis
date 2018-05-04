package dotplot

import (
	"io/ioutil"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/types"
)

// Hierarchical clusters dataset hierarchically and outputs files.
func Hierarchical(dataset types.Dataset) {
	// Generate bait-prey table.
	matrix, baitList, preyList := BaitPreyMatrix(dataset.Data)

	// Generate bait and prey distance matrices.
	baitDist := hclust.Distance(matrix, dataset.Params.Distance, true)
	preyDist := hclust.Distance(matrix, dataset.Params.Distance, false)

	// Bait and prey clustering.
	baitClust, err := hclust.Cluster(baitDist, dataset.Params.ClusteringMethod)
	logmessage.CheckError(err, true)
	preyClust, err := hclust.Cluster(preyDist, dataset.Params.ClusteringMethod)
	logmessage.CheckError(err, true)

	// Optimize clustering.
	baitClust = hclust.Optimize(baitClust, baitDist)
	preyClust = hclust.Optimize(preyClust, preyDist)

	// Create tree and get clustering order.
	baitTree, err := hclust.Tree(baitClust, baitList)
	logmessage.CheckError(err, true)
	preyTree, err := hclust.Tree(preyClust, preyList)
	logmessage.CheckError(err, true)

	// Write transformed data matrix to file.
	WriteMatrix(matrix, baitList, preyList)

	// Write transformed data matrix to file but as ratios instead of absolutes.
	WriteRatios(matrix, baitList, preyList)

	// Write bait-prey cytoscape file.
	WriteBPCytoscape(dataset)

	// Write bait-bait cytoscape file.
	WriteBBCytoscape(baitDist, baitList)

	// Write prey-prey cytoscape file.
	WritePPCytoscape(preyDist, preyList)

	// Write newick trees to files.
	ioutil.WriteFile("other/bait-dendrogram.txt", []byte(baitTree.Newick), 0644)
	ioutil.WriteFile("other/prey-dendrogram.txt", []byte(preyTree.Newick), 0644)
	return
}
