package dotplot

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// Biclustering using the biclustering method of Choi et. al (2010) to cluster
// data. The files produced from it will get put into a folder called
// "biclustering" and the order of baits and preys from that will be used
// for the images.
func Biclustering(dataset typedef.Dataset) {
	// Write log.
	LogParams(dataset.Params)

	// Create folder for biclustering files.
	biclustPath := filepath.Join(".", "biclustering")
	err := fs.Instance.MkdirAll(biclustPath, os.ModePerm)
	logmessage.CheckError(err, true)

	// Generate parameter file. This depends on whether approximate biclustering
	// should be performed.
	var params string
	if dataset.Params.BiclusteringApprox {
		// Get number of baits.
		uniqueBaits := make(map[string]bool)
		for _, row := range dataset.Data {
			bait := row["bait"].(string)
			if _, ok := uniqueBaits[bait]; !ok {
				uniqueBaits[bait] = true
			}
		}
		nb := len(uniqueBaits)

		// Create optimized param file content.
		params = fmt.Sprintf("np 10\n"+
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
		params = fmt.Sprintln("np 10\n" +
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
	afero.WriteFile(fs.Instance, "biclustering/params.txt", []byte(params), 0644)

	// Generate bait-prey table.
	data := BaitPreyMatrix(dataset.Data, dataset.Params.ScoreType)

	// Subset data matrix to only include preys that pass the minimum abundance
	// cutoff for at least two baits and return normalized matrix. Will also return
	// a list of the "singletons" (preys that don't meet criteria) for
	// appending to the clustering order.
	filteredData := BiclustFormat(data, dataset.Params.MinimumAbundance)

	// Run nested cluster.
	order := NestedClustering()

	// Add singletons to nested cluster order.
	order.Preys = append(order.Preys, filteredData.Singles...)

	// Sort matrices.
	sortedAbundance, _ := hclust.Sort(data.Abundance, data.Baits, order.Baits, "column")
	sortedAbundance, _ = hclust.Sort(sortedAbundance, data.Preys, order.Preys, "row")
	sortedRatios := NormalizeMatrix(sortedAbundance)
	sortedScores, _ := hclust.Sort(data.Score, data.Baits, order.Baits, "column")
	sortedScores, _ = hclust.Sort(sortedScores, data.Preys, order.Preys, "row")

	// Output svgs.

	// Write bait-prey dotplot.
	if dataset.Params.WriteDotplot {
		SvgDotplot(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			order.Baits,
			order.Preys,
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
			order.Baits,
			order.Preys,
			dataset.Params.ColorSpace,
			dataset.Params.MaximumAbundance,
		)
	}

	// Create pdfs from svg.
	svgList := make([]string, 0)
	svgMiniList := make([]string, 0)
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

	// Write bait-prey cytoscape file.
	WriteBPCytoscape(dataset)

	// Output other files.

	// Write transformed data matrix to file.
	WriteMatrix(sortedAbundance, order.Baits, order.Preys, "other/data-transformed.txt")

	// Write transformed data matrix to file but as ratios instead of absolutes.
	WriteMatrix(sortedRatios, order.Baits, order.Preys, "other/data-transformed-ratios.txt")

	// Create interactive files.
	if dataset.Params.WriteDotplot {
		json := InteractiveDotplot(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			order.Baits,
			order.Preys,
			dataset.Params,
			"minimap/dotplot.png",
		)
		afero.WriteFile(fs.Instance, "interactive/dotplot.json", []byte(json), 0644)
	}

	return
}
