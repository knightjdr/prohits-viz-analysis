package dotplot

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/interactive"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// Biclustering using the biclustering method of Choi et. al (2010) to cluster
// data. The files produced from it will get put into a folder called
// "biclustering" and the order of conditions and readouts from that will be used
// for the images.
func Biclustering(dataset typedef.Dataset) {
	// Write log.
	LogParams(dataset.Parameters)

	// Create folder for biclustering files.
	biclustPath := filepath.Join(".", "biclustering")
	err := fs.Instance.MkdirAll(biclustPath, os.ModePerm)
	logmessage.CheckError(err, true)

	// Generate parameter file. This depends on whether approximate biclustering
	// should be performed.
	var parameters string
	if dataset.Parameters.BiclusteringApprox {
		// Get number of conditions.
		uniqueConditions := make(map[string]bool)
		for _, row := range dataset.FileData {
			condition := row["condition"]
			if _, ok := uniqueConditions[condition]; !ok {
				uniqueConditions[condition] = true
			}
		}
		nb := len(uniqueConditions)

		// Create optimized param file content.
		parameters = fmt.Sprintf("np 10\n"+
			"nb %d\n"+
			"a 1.0\n"+
			"b 1.0\n"+
			"lambda 0.0\n"+
			"nu 25.0\n"+
			"alpha 1.0\n"+
			"rho 1.0\n"+
			"gamma 1.0\n"+
			"nburn 50\n"+
			"niter 500\n", nb,
		)
	} else {
		// Create default param file content.
		parameters = fmt.Sprintln("np 10\n" +
			"nb 100\n" +
			"a 1.0\n" +
			"b 1.0\n" +
			"lambda 0.0\n" +
			"nu 25.0\n" +
			"alpha 1.0\n" +
			"rho 1.0\n" +
			"gamma 1.0\n" +
			"nburn 5000\n" +
			"niter 10000\n",
		)
	}
	afero.WriteFile(fs.Instance, "biclustering/parameters.txt", []byte(parameters), 0644)

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

	// Output svgs.

	// Write condition-readout dotplot.
	if dataset.Parameters.WriteDotplot {
		SvgDotplot(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			order.Conditions,
			order.Readouts,
			false,
			dataset.Parameters,
		)

		// Write minimap.
		Minimap(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			order.Conditions,
			order.Readouts,
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

	// Write condition-readout heatmap.
	if dataset.Parameters.WriteHeatmap {
		SvgHeatmap(
			sortedAbundance,
			order.Conditions,
			order.Readouts,
			dataset.Parameters.FillColor,
			dataset.Parameters.AbundanceCap,
			false,
		)
	}

	// Create pdfs from svg.
	svgList := make([]string, 0)
	svgMiniList := make([]string, 0)
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

	// Write condition-readout cytoscape file.
	WriteBPCytoscape(dataset)

	// Output other files.

	// Write transformed data matrix to file.
	WriteMatrix(sortedAbundance, order.Conditions, order.Readouts, "other/data-transformed.txt")

	// Write transformed data matrix to file but as ratios instead of absolutes.
	WriteMatrix(sortedRatios, order.Conditions, order.Readouts, "other/data-transformed-ratios.txt")

	// Create interactive files.
	if dataset.Parameters.WriteDotplot {
		json := interactive.ParseHeatmap(
			"dotplot",
			sortedAbundance,
			sortedRatios,
			sortedScores,
			order.Conditions,
			order.Readouts,
			false,
			dataset.Parameters,
			"minimap/dotplot.png",
		)
		afero.WriteFile(fs.Instance, "interactive/dotplot.json", []byte(json), 0644)
	}

	return
}
