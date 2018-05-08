package dotplot

import (
	"io/ioutil"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Hierarchical clusters dataset hierarchically and outputs files.
func Hierarchical(dataset typedef.Dataset) {
	// Generate bait-prey table.
	data := BaitPreyMatrix(dataset.Data, dataset.Params.ScoreType)

	// Generate bait and prey distance matrices.
	baitDist := hclust.Distance(data.Abundance, dataset.Params.Distance, true)
	preyDist := hclust.Distance(data.Abundance, dataset.Params.Distance, false)

	// Bait and prey clustering.
	baitClust, err := hclust.Cluster(baitDist, dataset.Params.ClusteringMethod)
	logmessage.CheckError(err, true)
	preyClust, err := hclust.Cluster(preyDist, dataset.Params.ClusteringMethod)
	logmessage.CheckError(err, true)

	// Optimize clustering.
	baitClust = hclust.Optimize(baitClust, baitDist)
	preyClust = hclust.Optimize(preyClust, preyDist)

	// Create tree and get clustering order.
	baitTree, err := hclust.Tree(baitClust, data.Baits)
	logmessage.CheckError(err, true)
	preyTree, err := hclust.Tree(preyClust, data.Preys)
	logmessage.CheckError(err, true)

	// Create matrix with normalized rows for bait-prey abundances.
	ratios := NormalizeMatrix(data.Abundance)

	// Output svgs.

	// Output bait-bait svg.
	SvgBB(baitDist, data.Baits, baitTree.Order, dataset.Params.ColorSpace)

	// Output bait-prey dotplot.
	SvgDotplot(
		data.Abundance,
		ratios,
		data.Score,
		data.Baits,
		data.Preys,
		baitTree.Order,
		preyTree.Order,
		dataset.Params,
	)

	// Output bait-prey heatmap.
	SvgHeatmap(
		data.Abundance,
		data.Baits,
		data.Preys,
		baitTree.Order,
		preyTree.Order,
		dataset.Params.ColorSpace,
		dataset.Params.MaximumAbundance,
	)

	// Output prey-prey svg.
	SvgPP(preyDist, data.Preys, preyTree.Order, dataset.Params.ColorSpace)

	// Output cytoscape files.

	// Write bait-bait cytoscape file.
	WriteBBCytoscape(baitDist, data.Baits)

	// Write bait-prey cytoscape file.
	WriteBPCytoscape(dataset)

	// Write prey-prey cytoscape file.
	WritePPCytoscape(preyDist, data.Preys)

	// Output other files.

	// Write transformed data matrix to file.
	WriteMatrix(data.Abundance, data.Baits, data.Preys, "other/data-transformed.txt")

	// Write transformed data matrix to file but as ratios instead of absolutes.
	WriteMatrix(ratios, data.Baits, data.Preys, "other/data-transformed-ratios.txt")

	// Write newick trees to files.
	ioutil.WriteFile("other/bait-dendrogram.txt", []byte(baitTree.Newick), 0644)
	ioutil.WriteFile("other/prey-dendrogram.txt", []byte(preyTree.Newick), 0644)
	return
}
