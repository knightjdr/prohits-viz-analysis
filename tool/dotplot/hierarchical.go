package dotplot

import (
	"io/ioutil"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Hierarchical clusters dataset hierarchically and outputs files.
func Hierarchical(dataset typedef.Dataset) {
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

	// Output svg.
	params := map[string]interface{}{
		"colorSpace":       dataset.Params.ColorSpace,
		"maximumAbundance": dataset.Params.MaximumAbundance,
	}
	heatmap := svg.Heatmap(matrix, baitList, preyList, params)
	ioutil.WriteFile("svg/heatmap.svg", []byte(heatmap), 0644)

	// Output cytoscape files.

	// Write bait-bait cytoscape file.
	WriteBBCytoscape(baitDist, baitList)

	// Write bait-prey cytoscape file.
	WriteBPCytoscape(dataset)

	// Write prey-prey cytoscape file.
	WritePPCytoscape(preyDist, preyList)

	// Output other files.

	// Write transformed data matrix to file.
	WriteMatrix(matrix, baitList, preyList)

	// Write transformed data matrix to file but as ratios instead of absolutes.
	WriteRatios(matrix, baitList, preyList)

	// Write newick trees to files.
	ioutil.WriteFile("other/bait-dendrogram.txt", []byte(baitTree.Newick), 0644)
	ioutil.WriteFile("other/prey-dendrogram.txt", []byte(preyTree.Newick), 0644)
	return
}
