package dotplot

import (
	"fmt"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
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

	// Sort abundance matrix.
	sortedAbundance, _ := hclust.Sort(data.Abundance, data.Baits, baitTree.Order, "column")
	sortedAbundance, _ = hclust.Sort(sortedAbundance, data.Preys, preyTree.Order, "row")
	sortedRatios := NormalizeMatrix(sortedAbundance)
	sortedScores, _ := hclust.Sort(data.Score, data.Baits, baitTree.Order, "column")
	sortedScores, _ = hclust.Sort(sortedScores, data.Preys, preyTree.Order, "row")

	// Write log.
	LogParams(dataset.Params)

	// Output svgs.

	// Write bait-bait svg.
	SvgBB(baitDist, data.Baits, baitTree.Order, dataset.Params.ColorSpace)

	// Write bait-prey dotplot.
	SvgDotplot(
		sortedAbundance,
		sortedRatios,
		sortedScores,
		baitTree.Order,
		preyTree.Order,
		dataset.Params,
	)

	// Write bait-prey heatmap.
	SvgHeatmap(
		sortedAbundance,
		baitTree.Order,
		preyTree.Order,
		dataset.Params.ColorSpace,
		dataset.Params.MaximumAbundance,
	)

	// Write prey-prey svg.
	SvgPP(preyDist, data.Preys, preyTree.Order, dataset.Params.ColorSpace)

	// Write dotplot legend.
	legendTitle := fmt.Sprintf("Dotplot - %s", dataset.Params.Abundance)
	dotplotLegend := svg.DotplotLegend(
		dataset.Params.ColorSpace,
		legendTitle,
		101,
		0,
		dataset.Params.MaximumAbundance,
		dataset.Params.PrimaryFilter,
		dataset.Params.SecondaryFilter,
		dataset.Params.Score,
		dataset.Params.ScoreType,
	)
	afero.WriteFile(fs.Instance, "svg/dotplot-legend.svg", []byte(dotplotLegend), 0644)

	// Write distance legend.
	legendTitle = fmt.Sprintf("Distance - %s", dataset.Params.Abundance)
	distanceLegend := svg.Gradient(dataset.Params.ColorSpace, legendTitle, 101, 0, dataset.Params.MaximumAbundance)
	afero.WriteFile(fs.Instance, "svg/distance-legend.svg", []byte(distanceLegend), 0644)

	// Create pdfs and pngs from svg.
	svgList := []string{
		"bait-bait.svg",
		"distance-legend.svg",
		"dotplot-legend.svg",
		"dotplot.svg",
		"heatmap.svg",
		// "prey-prey.svg",
	}
	svg.ConvertPdf(svgList)
	svg.ConvertPng(svgList)

	// Output cytoscape files.

	// Write bait-bait cytoscape file.
	WriteBBCytoscape(baitDist, data.Baits)

	// Write bait-prey cytoscape file.
	WriteBPCytoscape(dataset)

	// Write prey-prey cytoscape file.
	WritePPCytoscape(preyDist, data.Preys)

	// Output other files.

	// Write transformed data matrix to file.
	WriteMatrix(sortedAbundance, baitTree.Order, preyTree.Order, "other/data-transformed.txt")

	// Write transformed data matrix to file but as ratios instead of absolutes.
	WriteMatrix(sortedRatios, baitTree.Order, preyTree.Order, "other/data-transformed-ratios.txt")

	// Write newick trees to files.
	afero.WriteFile(fs.Instance, "other/bait-dendrogram.txt", []byte(baitTree.Newick), 0644)
	afero.WriteFile(fs.Instance, "other/prey-dendrogram.txt", []byte(preyTree.Newick), 0644)

	// Create interactive files.
	return
}
