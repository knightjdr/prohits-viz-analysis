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

	// Normalize distance matrices to 1.
	maxDist := float64(0)
	normalizedBaitDist := baitDist
	for _, row := range normalizedBaitDist {
		for _, dist := range row {
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	for i, row := range normalizedBaitDist {
		for j, dist := range row {
			normalizedBaitDist[i][j] = dist / maxDist
		}
	}
	maxDist = float64(0)
	normalizedPreyDist := preyDist
	for _, row := range normalizedPreyDist {
		for _, dist := range row {
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	for i, row := range normalizedPreyDist {
		for j, dist := range row {
			normalizedPreyDist[i][j] = dist / maxDist
		}
	}

	// Sort matrices.
	sortedAbundance, _ := hclust.Sort(data.Abundance, data.Baits, baitTree.Order, "column")
	sortedAbundance, _ = hclust.Sort(sortedAbundance, data.Preys, preyTree.Order, "row")
	sortedBaitDist, _ := hclust.Sort(normalizedBaitDist, data.Baits, baitTree.Order, "column")
	sortedBaitDist, _ = hclust.Sort(sortedBaitDist, data.Baits, baitTree.Order, "row")
	sortedPreyDist, _ := hclust.Sort(normalizedPreyDist, data.Preys, preyTree.Order, "column")
	sortedPreyDist, _ = hclust.Sort(sortedPreyDist, data.Preys, preyTree.Order, "row")
	sortedRatios := NormalizeMatrix(sortedAbundance)
	sortedScores, _ := hclust.Sort(data.Score, data.Baits, baitTree.Order, "column")
	sortedScores, _ = hclust.Sort(sortedScores, data.Preys, preyTree.Order, "row")

	// Write log.
	LogParams(dataset.Params)

	// Output svgs.

	// Write bait-bait svg.
	if dataset.Params.WriteDistance {
		SvgBB(sortedBaitDist, baitTree.Order, dataset.Params.ColorSpace)

		// Write distance legend.
		legendTitle := fmt.Sprintf("Distance - %s", dataset.Params.Abundance)
		distanceLegend := svg.Gradient(dataset.Params.ColorSpace, legendTitle, 101, 0, dataset.Params.MaximumAbundance)
		afero.WriteFile(fs.Instance, "svg/distance-legend.svg", []byte(distanceLegend), 0644)
	}

	// Write bait-prey dotplot.
	if dataset.Params.WriteDotplot {
		SvgDotplot(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			baitTree.Order,
			preyTree.Order,
			dataset.Params,
		)

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
	}

	// Write bait-prey heatmap.
	if dataset.Params.WriteHeatmap {
		SvgHeatmap(
			sortedAbundance,
			baitTree.Order,
			preyTree.Order,
			dataset.Params.ColorSpace,
			dataset.Params.MaximumAbundance,
		)
	}

	// Write prey-prey svg.
	if dataset.Params.WriteDistance {
		SvgPP(sortedPreyDist, preyTree.Order, dataset.Params.ColorSpace)
	}

	// Create pdfs from svg.
	svgList := make([]string, 0)
	svgMiniList := make([]string, 0)
	if dataset.Params.WriteDistance {
		svgList = append(svgList, "bait-bait.svg")
		svgList = append(svgList, "distance-legend.svg")
		svgList = append(svgList, "prey-prey.svg")
		svgMiniList = append(svgMiniList, "bait-bait.svg")
		svgMiniList = append(svgMiniList, "prey-prey.svg")
	}
	if dataset.Params.WriteDotplot {
		svgList = append(svgList, "dotplot.svg")
		svgList = append(svgList, "dotplot-legend.svg")
		svgMiniList = append(svgMiniList, "dotplot.svg")
	}
	if dataset.Params.WriteHeatmap {
		svgList = append(svgList, "heatmap.svg")
	}
	if dataset.Params.Pdf {
		svg.ConvertPdf(svgList)
	}
	if dataset.Params.Png {
		svg.ConvertPng(svgList)
	}

	// Create  minimaps from svg.
	svg.ConvertMap(svgMiniList)

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
	if dataset.Params.WriteDistance {
		json := InteractiveHeatmap(
			normalizedBaitDist,
			baitTree.Order,
			baitTree.Order,
			dataset.Params,
			"minimap/bait-bait.png",
		)
		afero.WriteFile(fs.Instance, "interactive/bait-bait.txt", []byte(json), 0644)
		json = InteractiveHeatmap(
			normalizedPreyDist,
			preyTree.Order,
			preyTree.Order,
			dataset.Params,
			"minimap/prey-prey.png",
		)
		afero.WriteFile(fs.Instance, "interactive/prey-prey.txt", []byte(json), 0644)
	}
	if dataset.Params.WriteDotplot {
		json := InteractiveDotplot(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			baitTree.Order,
			preyTree.Order,
			dataset.Params,
			"minimap/dotplot.png",
		)
		afero.WriteFile(fs.Instance, "interactive/dotplot.txt", []byte(json), 0644)
	}

	return
}
