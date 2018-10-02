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

// NoCluster creates a dotplot using a list of baits and preys.
func NoCluster(dataset typedef.Dataset) {
	// Write log.
	LogParams(dataset.Parameters)

	// Generate bait-prey table.
	data := BaitPreyMatrix(dataset.Data, dataset.Parameters.ScoreType)

	// Cluster baits.
	var baitOrder []string
	if dataset.Parameters.BaitClustering == "none" {
		// Generate distance matrix.
		baitDist := hclust.Distance(data.Abundance, dataset.Parameters.Distance, true)

		// Cluster.
		baitClust, err := hclust.Cluster(baitDist, dataset.Parameters.ClusteringMethod)
		logmessage.CheckError(err, true)

		// Optimize clustering.
		baitClust = hclust.Optimize(baitClust, baitDist)

		// Create tree and get clustering order.
		baitTree, err := hclust.Tree(baitClust, data.Baits)
		logmessage.CheckError(err, true)

		// Set bait order.
		baitOrder = baitTree.Order

		// Normalize distance matrix to 1.
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

		// Sort distance matrix.
		sortedBaitDist, _ := hclust.Sort(normalizedBaitDist, data.Baits, baitOrder, "column")
		sortedBaitDist, _ = hclust.Sort(sortedBaitDist, data.Baits, baitOrder, "row")

		// Write bait-bait svg.
		if dataset.Parameters.WriteDistance {
			SvgBB(sortedBaitDist, baitOrder, dataset.Parameters.FillColor)

			// Generate pdfs and pngs.
			if dataset.Parameters.Pdf {
				svg.ConvertPdf([]string{"bait-bait.svg"})
			}
			if dataset.Parameters.Png {
				svg.ConvertPng([]string{"bait-bait.svg"})
			}

			// Create  minimaps from svg.
			svg.ConvertMap([]string{"bait-bait.svg"})

			// Create interactive files.
			json := InteractiveHeatmap(
				normalizedBaitDist,
				baitOrder,
				baitOrder,
				true,
				dataset.Parameters,
				"minimap/bait-bait.png",
			)
			afero.WriteFile(fs.Instance, "interactive/bait-bait.json", []byte(json), 0644)
		}

		// Write bait-bait cytoscape file.
		WriteBBCytoscape(baitDist, data.Baits)

		// Write newick tree to file.
		afero.WriteFile(fs.Instance, "other/bait-dendrogram.txt", []byte(baitTree.Newick), 0644)
	} else {
		baitOrder = dataset.Parameters.BaitList
	}

	// Cluster preys.
	var preyOrder []string
	if dataset.Parameters.PreyClustering == "none" {
		// Generate distance matrix.
		preyDist := hclust.Distance(data.Abundance, dataset.Parameters.Distance, false)

		// Cluster.
		preyClust, err := hclust.Cluster(preyDist, dataset.Parameters.ClusteringMethod)
		logmessage.CheckError(err, true)

		// Optimize clustering.
		preyClust = hclust.Optimize(preyClust, preyDist)

		// Create tree and get clustering order.
		preyTree, err := hclust.Tree(preyClust, data.Preys)
		logmessage.CheckError(err, true)

		// Set prey order.
		preyOrder = preyTree.Order

		// Normalize distance matrix to 1.
		maxDist := float64(0)
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

		// Sort distance matrix.
		sortedPreyDist, _ := hclust.Sort(normalizedPreyDist, data.Preys, preyOrder, "column")
		sortedPreyDist, _ = hclust.Sort(sortedPreyDist, data.Preys, preyOrder, "row")

		// Write prey-prey svg.
		if dataset.Parameters.WriteDistance {
			SvgPP(sortedPreyDist, preyOrder, dataset.Parameters.FillColor)

			// Generate pdfs and pngs.
			if dataset.Parameters.Pdf {
				svg.ConvertPdf([]string{"prey-prey.svg"})
			}
			if dataset.Parameters.Png {
				svg.ConvertPng([]string{"prey-prey.svg"})
			}

			// Create  minimaps from svg.
			svg.ConvertMap([]string{"prey-prey.svg"})

			// Create interactive files.
			json := InteractiveHeatmap(
				normalizedPreyDist,
				preyOrder,
				preyOrder,
				true,
				dataset.Parameters,
				"minimap/prey-prey.png",
			)
			afero.WriteFile(fs.Instance, "interactive/prey-prey.json", []byte(json), 0644)
		}

		// Write prey-prey cytoscape file.
		WritePPCytoscape(preyDist, data.Preys)

		// Write newick tree to file.
		afero.WriteFile(fs.Instance, "other/prey-dendrogram.txt", []byte(preyTree.Newick), 0644)
	} else {
		preyOrder = dataset.Parameters.PreyList
	}

	// Sort matrices.
	sortedAbundance, _ := hclust.Sort(data.Abundance, data.Baits, baitOrder, "column")
	sortedAbundance, _ = hclust.Sort(sortedAbundance, data.Preys, preyOrder, "row")
	sortedRatios := NormalizeMatrix(sortedAbundance)
	sortedScores, _ := hclust.Sort(data.Score, data.Baits, baitOrder, "column")
	sortedScores, _ = hclust.Sort(sortedScores, data.Preys, preyOrder, "row")

	// Write bait-prey dotplot.
	if dataset.Parameters.WriteDotplot {
		SvgDotplot(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			baitOrder,
			preyOrder,
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

		// Create pdfs and pngs.
		if dataset.Parameters.Pdf {
			svg.ConvertPdf([]string{"dotplot.svg", "dotplot-legend.svg"})
		}
		if dataset.Parameters.Png {
			svg.ConvertPng([]string{"dotplot.svg", "dotplot-legend.svg"})
		}

		// Create minimaps from svg.
		svg.ConvertMap([]string{"dotplot.svg"})

		// Create interactive file.
		json := InteractiveDotplot(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			baitOrder,
			preyOrder,
			dataset.Parameters,
			"minimap/dotplot.png",
		)
		afero.WriteFile(fs.Instance, "interactive/dotplot.json", []byte(json), 0644)
	}

	// Write bait-prey heatmap.
	if dataset.Parameters.WriteHeatmap {
		SvgHeatmap(
			sortedAbundance,
			baitOrder,
			preyOrder,
			dataset.Parameters.FillColor,
			dataset.Parameters.AbundanceCap,
			false,
		)

		// Create pdfs and pngs.
		if dataset.Parameters.Pdf {
			svg.ConvertPdf([]string{"heatmap.svg"})
		}
		if dataset.Parameters.Png {
			svg.ConvertPng([]string{"heatmap.svg"})
		}
	}

	// Write distance legend.
	if dataset.Parameters.WriteDistance {
		// Write distance legend.
		legendTitle := fmt.Sprintf("Distance - %s", dataset.Parameters.Abundance)
		distanceLegend := svg.Gradient(dataset.Parameters.FillColor, legendTitle, 101, 0, dataset.Parameters.AbundanceCap)
		afero.WriteFile(fs.Instance, "svg/distance-legend.svg", []byte(distanceLegend), 0644)

		// Generate pdfs and pngs.
		if dataset.Parameters.Pdf {
			svg.ConvertPdf([]string{"distance-legend.svg"})
		}
		if dataset.Parameters.Png {
			svg.ConvertPng([]string{"distance-legend.svg"})
		}
	}

	// Output other files.

	// Write bait-prey cytoscape file.
	WriteBPCytoscape(dataset)

	// Write transformed data matrix to file.
	WriteMatrix(sortedAbundance, baitOrder, preyOrder, "other/data-transformed.txt")

	// Write transformed data matrix to file but as ratios instead of absolutes.
	WriteMatrix(sortedRatios, baitOrder, preyOrder, "other/data-transformed-ratios.txt")
	return
}
