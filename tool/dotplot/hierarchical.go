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
	// Write log.
	LogParams(dataset.Parameters)

	// Generate bait-prey table.
	data := BaitPreyMatrix(dataset.Data, dataset.Parameters.ScoreType)

	// Generate bait and prey distance matrices.
	baitDist := hclust.Distance(data.Abundance, dataset.Parameters.Distance, true)
	preyDist := hclust.Distance(data.Abundance, dataset.Parameters.Distance, false)

	// Bait and prey clustering.
	baitClust, err := hclust.Cluster(baitDist, dataset.Parameters.ClusteringMethod)
	logmessage.CheckError(err, true)
	preyClust, err := hclust.Cluster(preyDist, dataset.Parameters.ClusteringMethod)
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

	// Output svgs.

	// Write bait-bait svg.
	if dataset.Parameters.WriteDistance {
		SvgBB(sortedBaitDist, baitTree.Order, dataset.Parameters.FillColor)

		// Write distance legend.
		legendTitle := fmt.Sprintf("Distance - %s", dataset.Parameters.Abundance)
		distanceLegend := svg.Gradient(dataset.Parameters.FillColor, legendTitle, 101, 0, dataset.Parameters.AbundanceCap)
		afero.WriteFile(fs.Instance, "svg/distance-legend.svg", []byte(distanceLegend), 0644)
	}

	// Write bait-prey dotplot.
	if dataset.Parameters.WriteDotplot {
		SvgDotplot(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			baitTree.Order,
			preyTree.Order,
			false,
			dataset.Parameters,
		)

		// Write dotplot legend.
		legendTitle := fmt.Sprintf("Dotplot - %s", dataset.Parameters.Abundance)
		dotplotLegend := svg.DotplotLegend(
			dataset.Parameters.FillColor,
			legendTitle,
			101,
			0,
			dataset.Parameters.AbundanceCap,
			dataset.Parameters.PrimaryFilter,
			dataset.Parameters.SecondaryFilter,
			dataset.Parameters.Score,
			dataset.Parameters.ScoreType,
		)
		afero.WriteFile(fs.Instance, "svg/dotplot-legend.svg", []byte(dotplotLegend), 0644)
	}

	// Write bait-prey heatmap.
	if dataset.Parameters.WriteHeatmap {
		SvgHeatmap(
			sortedAbundance,
			baitTree.Order,
			preyTree.Order,
			dataset.Parameters.FillColor,
			dataset.Parameters.AbundanceCap,
			false,
		)
	}

	// Write prey-prey svg.
	if dataset.Parameters.WriteDistance {
		SvgPP(sortedPreyDist, preyTree.Order, dataset.Parameters.FillColor)
	}

	// Create pdfs from svg.
	svgList := make([]string, 0)
	svgMiniList := make([]string, 0)
	if dataset.Parameters.WriteDistance {
		svgList = append(svgList, "bait-bait.svg")
		svgList = append(svgList, "distance-legend.svg")
		svgList = append(svgList, "prey-prey.svg")
		svgMiniList = append(svgMiniList, "bait-bait.svg")
		svgMiniList = append(svgMiniList, "prey-prey.svg")
	}
	if dataset.Parameters.WriteDotplot {
		svgList = append(svgList, "dotplot.svg")
		svgList = append(svgList, "dotplot-legend.svg")
		svgMiniList = append(svgMiniList, "dotplot.svg")
	}
	if dataset.Parameters.WriteHeatmap {
		svgList = append(svgList, "heatmap.svg")
	}
	if dataset.Parameters.Pdf {
		svg.ConvertPdf(svgList)
	}
	if dataset.Parameters.Png {
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
	if dataset.Parameters.WriteDistance {
		json := InteractiveHeatmap(
			normalizedBaitDist,
			baitTree.Order,
			baitTree.Order,
			true,
			dataset.Parameters,
			"minimap/bait-bait.png",
		)
		afero.WriteFile(fs.Instance, "interactive/bait-bait.json", []byte(json), 0644)
		json = InteractiveHeatmap(
			normalizedPreyDist,
			preyTree.Order,
			preyTree.Order,
			true,
			dataset.Parameters,
			"minimap/prey-prey.png",
		)
		afero.WriteFile(fs.Instance, "interactive/prey-prey.json", []byte(json), 0644)
	}
	if dataset.Parameters.WriteDotplot {
		json := InteractiveDotplot(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			baitTree.Order,
			preyTree.Order,
			dataset.Parameters,
			"minimap/dotplot.png",
		)
		afero.WriteFile(fs.Instance, "interactive/dotplot.json", []byte(json), 0644)
	}

	return
}
