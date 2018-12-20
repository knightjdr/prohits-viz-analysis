package nocluster

import (
	"fmt"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/image/svg"
	"github.com/knightjdr/prohits-viz-analysis/tool/dotplot/file"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

const ignore = 250000

// Run creates a dotplot using a list of conditions and readouts.
func Run(dataset *typedef.Dataset) {
	// Write log.
	file.LogParams(dataset.Parameters)

	// Generate condition-readout table.
	matrices := helper.ConditionReadoutMatrix(dataset.FileData, dataset.Parameters.ScoreType, true, false)

	// Cluster conditions.
	conditionOrder := order("condition", dataset, matrices)
	readoutOrder := order("readout", dataset, matrices)

	// Sort matrices.
	sortedAbundance, _ := hclust.Sort(matrices.Abundance, matrices.Conditions, conditionOrder, "column")
	sortedAbundance, _ = hclust.Sort(sortedAbundance, matrices.Readouts, readoutOrder, "row")
	sortedRatios := helper.NormalizeMatrix(sortedAbundance)
	sortedScores, _ := hclust.Sort(matrices.Score, matrices.Conditions, conditionOrder, "column")
	sortedScores, _ = hclust.Sort(sortedScores, matrices.Readouts, readoutOrder, "row")

	// Dotplot
	if dataset.Parameters.WriteDotplot {
		// svg
		file.Heatmap("dotplot", sortedAbundance, sortedRatios, sortedScores, conditionOrder, readoutOrder, dataset.Parameters)

		// minimap
		file.Minimap("dotplot", sortedAbundance, sortedRatios, sortedScores, conditionOrder, readoutOrder, dataset.Parameters)

		// Write dotplot legend.
		legendTitle := fmt.Sprintf("Dotplot - %s", dataset.Parameters.Abundance)
		dotplotLegend := svg.DotplotLegend(legendTitle, 101, dataset.Parameters)
		afero.WriteFile(fs.Instance, "svg/dotplot-legend.svg", []byte(dotplotLegend), 0644)

		// interactive file
		file.InteractiveHeatmap(
			"dotplot",
			sortedAbundance,
			sortedRatios,
			sortedScores,
			conditionOrder,
			readoutOrder,
			dataset.Parameters,
		)
	}

	// Write condition-readout heatmap.
	if dataset.Parameters.WriteHeatmap {
		file.Heatmap("heatmap", sortedAbundance, sortedRatios, sortedScores, conditionOrder, readoutOrder, dataset.Parameters)
	}

	// Write distance legend.
	if dataset.Parameters.WriteDistance {
		// Write distance legend.
		legendTitle := fmt.Sprintf("Distance - %s", dataset.Parameters.Abundance)
		distanceLegend := svg.Gradient(dataset.Parameters.FillColor, legendTitle, 101, 1, 0, false)
		afero.WriteFile(fs.Instance, "svg/distance-legend.svg", []byte(distanceLegend), 0644)
	}

	// Create pdfs from svg.
	svgList := make([]string, 0)
	if dataset.Parameters.WriteDistance {
		svgList = append(svgList, "distance-legend.svg")
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

	// Output other files.
	// Write condition-readout cytoscape file.
	file.WriteBPCytoscape(dataset)

	// Write transformed data matrix to file.
	file.WriteMatrix(sortedAbundance, conditionOrder, readoutOrder, "other/data-transformed.txt")

	// Write transformed data matrix to file but as ratios instead of absolutes.
	file.WriteMatrix(sortedRatios, conditionOrder, readoutOrder, "other/data-transformed-ratios.txt")
	return
}
