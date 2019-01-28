package hierarchical

import (
	"fmt"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/image/svg"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/tool/dotplot/file"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

const ignore = 250000

// Run clusters dataset hierarchically and outputs files.
func Run(dataset *typedef.Dataset) {
	// Generate condition-readout table.
	matrices := helper.ConditionReadoutMatrix(dataset.FileData, dataset.Parameters.ScoreType, true, false)

	// Generate condition and readout distance matrices.
	conditionDist := hclust.Distance(matrices.Abundance, dataset.Parameters.Distance, true)
	readoutDist := hclust.Distance(matrices.Abundance, dataset.Parameters.Distance, false)

	// Condition and readout clustering.
	conditionClust, err := hclust.Cluster(conditionDist, dataset.Parameters.ClusteringMethod)
	logmessage.CheckError(err, true)
	readoutClust, err := hclust.Cluster(readoutDist, dataset.Parameters.ClusteringMethod)
	logmessage.CheckError(err, true)

	// Optimize clustering.
	if dataset.Parameters.ClusteringOptimize {
		conditionClust = hclust.Optimize(conditionClust, conditionDist, ignore)
		readoutClust = hclust.Optimize(readoutClust, readoutDist, ignore)
	}

	// Create tree and get clustering order.
	conditionTree, err := hclust.Tree(conditionClust, matrices.Conditions)
	logmessage.CheckError(err, true)
	readoutTree, err := hclust.Tree(readoutClust, matrices.Readouts)
	logmessage.CheckError(err, true)

	// Normalize distance matrices to 1.
	normalizedConditionDist := helper.NormalizeMatrix(conditionDist)
	normalizedReadoutDist := helper.NormalizeMatrix(readoutDist)

	// Sort matrices.
	sortedAbundance, _ := hclust.Sort(matrices.Abundance, matrices.Conditions, conditionTree.Order, "column")
	sortedAbundance, _ = hclust.Sort(sortedAbundance, matrices.Readouts, readoutTree.Order, "row")
	sortedConditionDist, _ := hclust.Sort(normalizedConditionDist, matrices.Conditions, conditionTree.Order, "column")
	sortedConditionDist, _ = hclust.Sort(sortedConditionDist, matrices.Conditions, conditionTree.Order, "row")
	sortedReadoutDist, _ := hclust.Sort(normalizedReadoutDist, matrices.Readouts, readoutTree.Order, "column")
	sortedReadoutDist, _ = hclust.Sort(sortedReadoutDist, matrices.Readouts, readoutTree.Order, "row")
	sortedRatios := helper.NormalizeMatrix(sortedAbundance)
	sortedScores, _ := hclust.Sort(matrices.Score, matrices.Conditions, conditionTree.Order, "column")
	sortedScores, _ = hclust.Sort(sortedScores, matrices.Readouts, readoutTree.Order, "row")

	// Dotplot
	if dataset.Parameters.WriteDotplot {
		// svg
		file.Heatmap("dotplot", sortedAbundance, sortedRatios, sortedScores, conditionTree.Order, readoutTree.Order, dataset.Parameters)

		// minimap
		file.Minimap("dotplot", sortedAbundance, sortedRatios, sortedScores, conditionTree.Order, readoutTree.Order, dataset.Parameters)

		// interactive file
		file.InteractiveHeatmap(
			"dotplot",
			sortedAbundance,
			sortedRatios,
			sortedScores,
			conditionTree.Order,
			readoutTree.Order,
			dataset.Parameters,
		)

		// Write dotplot legend.
		legendTitle := fmt.Sprintf("Dotplot - %s", dataset.Parameters.Abundance)
		dotplotLegend := svg.DotplotLegend(legendTitle, 101, dataset.Parameters)
		afero.WriteFile(fs.Instance, "svg/dotplot-legend.svg", []byte(dotplotLegend), 0644)
	}

	// Distance images.
	if dataset.Parameters.WriteDistance {
		// svg
		file.Distance(sortedConditionDist, conditionTree.Order, dataset.Parameters.Condition, "condition-condition", dataset.Parameters)
		file.Distance(sortedReadoutDist, readoutTree.Order, dataset.Parameters.Readout, "readout-readout", dataset.Parameters)

		// minimaps
		file.MinimapDistance(sortedConditionDist, conditionTree.Order, dataset.Parameters.FillColor, "condition-condition")
		file.MinimapDistance(sortedReadoutDist, readoutTree.Order, dataset.Parameters.FillColor, "readout-readout")

		// interactive files
		file.InteractiveDistance(sortedConditionDist, conditionTree.Order, dataset.Parameters.Condition, "condition-condition", dataset.Parameters)
		file.InteractiveDistance(sortedReadoutDist, readoutTree.Order, dataset.Parameters.Readout, "readout-readout", dataset.Parameters)

		// legend
		legendTitle := fmt.Sprintf("Distance - %s", dataset.Parameters.Abundance)
		distanceLegend := svg.Gradient(dataset.Parameters.FillColor, legendTitle, 101, 1, 0, false)
		afero.WriteFile(fs.Instance, "svg/distance-legend.svg", []byte(distanceLegend), 0644)
	}

	// Write condition-readout heatmap.
	if dataset.Parameters.WriteHeatmap {
		file.Heatmap("heatmap", sortedAbundance, sortedRatios, sortedScores, conditionTree.Order, readoutTree.Order, dataset.Parameters)
	}

	// Create pdfs from svg.
	svgList := make([]string, 0)
	if dataset.Parameters.WriteDistance {
		svgList = append(svgList, "condition-condition.svg")
		svgList = append(svgList, "distance-legend.svg")
		svgList = append(svgList, "readout-readout.svg")
	}
	if dataset.Parameters.WriteDotplot {
		svgList = append(svgList, "dotplot.svg")
		svgList = append(svgList, "dotplot-legend.svg")
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

	// Cytoscape files.
	file.WriteBBCytoscape(conditionDist, matrices.Conditions)
	file.WriteBPCytoscape(dataset)
	file.WritePPCytoscape(readoutDist, matrices.Readouts)

	// Other files.
	// Write transformed data matrix to file.
	file.WriteMatrix(sortedAbundance, conditionTree.Order, readoutTree.Order, "other/data-transformed.txt")

	// Write transformed data matrix to file but as ratios instead of absolutes.
	file.WriteMatrix(sortedRatios, conditionTree.Order, readoutTree.Order, "other/data-transformed-ratios.txt")

	// Write newick trees to files.
	afero.WriteFile(fs.Instance, "other/condition-dendrogram.txt", []byte(conditionTree.Newick), 0644)
	afero.WriteFile(fs.Instance, "other/readout-dendrogram.txt", []byte(readoutTree.Newick), 0644)

	return
}
