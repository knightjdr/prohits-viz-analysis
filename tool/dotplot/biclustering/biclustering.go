package biclustering

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/image/svg"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/tool/dotplot/file"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// Run using the biclustering method of Choi et. al (2010) to cluster
// data. The files produced from it will be created in a folder called
// "biclustering" and the order of conditions and readouts from that will be used
// for the images.
func Run(dataset *typedef.Dataset) {
	// Create folder for biclustering files.
	biclustPath := filepath.Join(".", "biclustering")
	err := fs.Instance.MkdirAll(biclustPath, os.ModePerm)
	logmessage.CheckError(err, true)

	// Set parameters
	parameters(dataset.FileData, dataset.Parameters)

	// Generate condition-readout table.
	data := helper.ConditionReadoutMatrix(dataset.FileData, dataset.Parameters.ScoreType, true, false)

	// Subset data matrix to only include readouts that pass the minimum abundance
	// cutoff for at least two conditions and return normalized matrix. Will also return
	// a list of the "singletons" (readouts that don't meet criteria) for
	// appending to the clustering order.
	filteredData := BiclustFormat(data, dataset.Parameters.MinAbundance)

	// Run nested cluster.
	order := NestedClustering()

	// Add singletons to nested cluster order.
	order.Readouts = append(order.Readouts, filteredData.Singles...)

	// Sort matrices.
	sortedAbundance, _ := hclust.Sort(data.Abundance, data.Conditions, order.Conditions, "column")
	sortedAbundance, _ = hclust.Sort(sortedAbundance, data.Readouts, order.Readouts, "row")
	sortedRatios := helper.NormalizeMatrix(sortedAbundance)
	sortedScores, _ := hclust.Sort(data.Score, data.Conditions, order.Conditions, "column")
	sortedScores, _ = hclust.Sort(sortedScores, data.Readouts, order.Readouts, "row")

	// Dotplot
	if dataset.Parameters.WriteDotplot {
		// svg
		file.Heatmap("dotplot", sortedAbundance, sortedRatios, sortedScores, order.Conditions, order.Readouts, dataset.Parameters)

		// minimap
		file.Minimap("dotplot", sortedAbundance, sortedRatios, sortedScores, order.Conditions, order.Readouts, dataset.Parameters)

		// interactive
		file.InteractiveHeatmap(
			"dotplot",
			sortedAbundance,
			sortedRatios,
			sortedScores,
			order.Conditions,
			order.Readouts,
			dataset.Parameters,
		)

		// legend
		legendTitle := fmt.Sprintf("Dotplot - %s", dataset.Parameters.Abundance)
		dotplotLegend := svg.DotplotLegend(legendTitle, 101, dataset.Parameters)
		afero.WriteFile(fs.Instance, "svg/dotplot-legend.svg", []byte(dotplotLegend), 0644)
	}

	// Heatmap
	if dataset.Parameters.WriteHeatmap {
		file.Heatmap("heatmap", sortedAbundance, [][]float64{}, [][]float64{}, order.Conditions, order.Readouts, dataset.Parameters)
	}

	// Create pdfs from svg.
	svgList := make([]string, 0)
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

	// Write condition-readout cytoscape file.
	file.WriteBPCytoscape(dataset)

	// Output other files.
	// Write transformed data matrix to file.
	file.WriteMatrix(sortedAbundance, order.Conditions, order.Readouts, "other/data-transformed.txt")

	// Write transformed data matrix to file but as ratios instead of absolutes.
	file.WriteMatrix(sortedRatios, order.Conditions, order.Readouts, "other/data-transformed-ratios.txt")

	return
}
