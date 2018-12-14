package dotplot

import (
	"fmt"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/interactive"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// NoCluster creates a dotplot using a list of conditions and readouts.
func NoCluster(dataset typedef.Dataset) {
	// Write log.
	LogParams(dataset.Parameters)

	// Generate condition-readout table.
	data := helper.ConditionReadoutMatrix(dataset.FileData, dataset.Parameters.ScoreType, true, false)

	// Cluster conditions.
	var conditionOrder []string
	if dataset.Parameters.ConditionClustering != "none" {
		// Generate distance matrix.
		conditionDist := hclust.Distance(data.Abundance, dataset.Parameters.Distance, true)

		// Cluster.
		conditionClust, err := hclust.Cluster(conditionDist, dataset.Parameters.ClusteringMethod)
		logmessage.CheckError(err, true)

		// Optimize clustering.
		if dataset.Parameters.ClusteringOptimize {
			conditionClust = hclust.Optimize(conditionClust, conditionDist, ignore)
		}

		// Create tree and get clustering order.
		conditionTree, err := hclust.Tree(conditionClust, data.Conditions)
		logmessage.CheckError(err, true)

		// Set condition order.
		conditionOrder = conditionTree.Order

		// Normalize distance matrix to 1.
		maxDist := float64(0)
		normalizedConditionDist := conditionDist
		for _, row := range normalizedConditionDist {
			for _, dist := range row {
				if dist > maxDist {
					maxDist = dist
				}
			}
		}
		for i, row := range normalizedConditionDist {
			for j, dist := range row {
				normalizedConditionDist[i][j] = dist / maxDist
			}
		}

		// Sort distance matrix.
		sortedConditionDist, _ := hclust.Sort(normalizedConditionDist, data.Conditions, conditionOrder, "column")
		sortedConditionDist, _ = hclust.Sort(sortedConditionDist, data.Conditions, conditionOrder, "row")

		// Write condition-condition svg.
		if dataset.Parameters.WriteDistance {
			SvgCC(sortedConditionDist, conditionOrder, dataset.Parameters)

			// Generate pdfs and pngs.
			if dataset.Parameters.Pdf {
				svg.ConvertPdf([]string{"condition-condition.svg"})
			}
			if dataset.Parameters.Png {
				svg.ConvertPng([]string{"condition-condition.svg"})
			}

			// Write minimap.
			MinimapDistance(
				sortedConditionDist,
				conditionOrder,
				dataset.Parameters.FillColor,
				"condition-condition",
			)

			// Create minimaps from svg.
			svg.ConvertMap([]string{"condition-condition.svg"})

			// Create interactive files.
			distanceParams := dataset.Parameters
			distanceParams.AbundanceCap = 1
			distanceParams.MinAbundance = 0
			json := interactive.ParseHeatmap(
				"heatmap",
				sortedConditionDist,
				[][]float64{},
				[][]float64{},
				conditionOrder,
				conditionOrder,
				true,
				distanceParams,
				"minimap/condition-condition.png",
				dataset.Parameters.Condition,
				dataset.Parameters.Condition,
			)
			afero.WriteFile(fs.Instance, "interactive/condition-condition.json", []byte(json), 0644)
		}

		// Write condition-condition cytoscape file.
		WriteBBCytoscape(conditionDist, data.Conditions)

		// Write newick tree to file.
		afero.WriteFile(fs.Instance, "other/condition-dendrogram.txt", []byte(conditionTree.Newick), 0644)
	} else {
		// Only keep user specified conditions that are actually in the file
		conditionOrder = checkList(dataset.FileData, "condition", dataset.Parameters.ConditionList)
	}

	// Cluster readouts.
	var readoutOrder []string
	if dataset.Parameters.ReadoutClustering != "none" {
		// Generate distance matrix.
		readoutDist := hclust.Distance(data.Abundance, dataset.Parameters.Distance, false)

		// Cluster.
		readoutClust, err := hclust.Cluster(readoutDist, dataset.Parameters.ClusteringMethod)
		logmessage.CheckError(err, true)

		// Optimize clustering.
		if dataset.Parameters.ClusteringOptimize {
			readoutClust = hclust.Optimize(readoutClust, readoutDist, ignore)
		}

		// Create tree and get clustering order.
		readoutTree, err := hclust.Tree(readoutClust, data.Readouts)
		logmessage.CheckError(err, true)

		// Set readout order.
		readoutOrder = readoutTree.Order

		// Normalize distance matrix to 1.
		maxDist := float64(0)
		normalizedReadoutDist := readoutDist
		for _, row := range normalizedReadoutDist {
			for _, dist := range row {
				if dist > maxDist {
					maxDist = dist
				}
			}
		}
		for i, row := range normalizedReadoutDist {
			for j, dist := range row {
				normalizedReadoutDist[i][j] = dist / maxDist
			}
		}

		// Sort distance matrix.
		sortedReadoutDist, _ := hclust.Sort(normalizedReadoutDist, data.Readouts, readoutOrder, "column")
		sortedReadoutDist, _ = hclust.Sort(sortedReadoutDist, data.Readouts, readoutOrder, "row")

		// Write readout-readout svg.
		if dataset.Parameters.WriteDistance {
			SvgRR(sortedReadoutDist, readoutOrder, dataset.Parameters)

			// Generate pdfs and pngs.
			if dataset.Parameters.Pdf {
				svg.ConvertPdf([]string{"readout-readout.svg"})
			}
			if dataset.Parameters.Png {
				svg.ConvertPng([]string{"readout-readout.svg"})
			}

			// Write minimap.
			MinimapDistance(
				sortedReadoutDist,
				readoutOrder,
				dataset.Parameters.FillColor,
				"readout-readout",
			)

			// Create  minimaps from svg.
			svg.ConvertMap([]string{"readout-readout.svg"})

			// Create interactive files.
			distanceParams := dataset.Parameters
			distanceParams.AbundanceCap = 1
			distanceParams.MinAbundance = 0
			json := interactive.ParseHeatmap(
				"heatmap",
				sortedReadoutDist,
				[][]float64{},
				[][]float64{},
				readoutOrder,
				readoutOrder,
				true,
				distanceParams,
				"minimap/readout-readout.png",
				dataset.Parameters.Readout,
				dataset.Parameters.Readout,
			)
			afero.WriteFile(fs.Instance, "interactive/readout-readout.json", []byte(json), 0644)
		}

		// Write readout-readout cytoscape file.
		WritePPCytoscape(readoutDist, data.Readouts)

		// Write newick tree to file.
		afero.WriteFile(fs.Instance, "other/readout-dendrogram.txt", []byte(readoutTree.Newick), 0644)
	} else {
		// Only keep user specified readouts that are actually in the file
		readoutOrder = checkList(dataset.FileData, "readout", dataset.Parameters.ReadoutList)
	}

	// Sort matrices.
	sortedAbundance, _ := hclust.Sort(data.Abundance, data.Conditions, conditionOrder, "column")
	sortedAbundance, _ = hclust.Sort(sortedAbundance, data.Readouts, readoutOrder, "row")
	sortedRatios := helper.NormalizeMatrix(sortedAbundance)
	sortedScores, _ := hclust.Sort(data.Score, data.Conditions, conditionOrder, "column")
	sortedScores, _ = hclust.Sort(sortedScores, data.Readouts, readoutOrder, "row")

	// Write condition-readout dotplot.
	if dataset.Parameters.WriteDotplot {
		SvgDotplot(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			conditionOrder,
			readoutOrder,
			false,
			dataset.Parameters,
		)

		// Write minimap.
		Minimap(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			conditionOrder,
			readoutOrder,
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
		json := interactive.ParseHeatmap(
			"dotplot",
			sortedAbundance,
			sortedRatios,
			sortedScores,
			conditionOrder,
			readoutOrder,
			false,
			dataset.Parameters,
			"minimap/dotplot.png",
			dataset.Parameters.Condition,
			dataset.Parameters.Readout,
		)
		afero.WriteFile(fs.Instance, "interactive/dotplot.json", []byte(json), 0644)
	}

	// Write condition-readout heatmap.
	if dataset.Parameters.WriteHeatmap {
		SvgHeatmap(
			sortedAbundance,
			conditionOrder,
			readoutOrder,
			dataset.Parameters,
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
		distanceLegend := svg.Gradient(dataset.Parameters.FillColor, legendTitle, 101, 1, 0, false)
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

	// Write condition-readout cytoscape file.
	WriteBPCytoscape(dataset)

	// Write transformed data matrix to file.
	WriteMatrix(sortedAbundance, conditionOrder, readoutOrder, "other/data-transformed.txt")

	// Write transformed data matrix to file but as ratios instead of absolutes.
	WriteMatrix(sortedRatios, conditionOrder, readoutOrder, "other/data-transformed-ratios.txt")
	return
}
