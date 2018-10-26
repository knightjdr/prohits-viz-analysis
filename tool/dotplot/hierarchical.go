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

	// Generate condition-readout table.
	data := ConditionReadoutMatrix(dataset.Data, dataset.Parameters.ScoreType)

	// Generate condition and readout distance matrices.
	conditionDist := hclust.Distance(data.Abundance, dataset.Parameters.Distance, true)
	readoutDist := hclust.Distance(data.Abundance, dataset.Parameters.Distance, false)

	// Condition and readout clustering.
	conditionClust, err := hclust.Cluster(conditionDist, dataset.Parameters.ClusteringMethod)
	logmessage.CheckError(err, true)
	readoutClust, err := hclust.Cluster(readoutDist, dataset.Parameters.ClusteringMethod)
	logmessage.CheckError(err, true)

	// Optimize clustering.
	if dataset.Parameters.ClusteringOptimize {
		conditionClust = hclust.Optimize(conditionClust, conditionDist)
		readoutClust = hclust.Optimize(readoutClust, readoutDist)
	}

	// Create tree and get clustering order.
	conditionTree, err := hclust.Tree(conditionClust, data.Conditions)
	logmessage.CheckError(err, true)
	readoutTree, err := hclust.Tree(readoutClust, data.Readouts)
	logmessage.CheckError(err, true)

	// Normalize distance matrices to 1.
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
	maxDist = float64(0)
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

	// Sort matrices.
	sortedAbundance, _ := hclust.Sort(data.Abundance, data.Conditions, conditionTree.Order, "column")
	sortedAbundance, _ = hclust.Sort(sortedAbundance, data.Readouts, readoutTree.Order, "row")
	sortedConditionDist, _ := hclust.Sort(normalizedConditionDist, data.Conditions, conditionTree.Order, "column")
	sortedConditionDist, _ = hclust.Sort(sortedConditionDist, data.Conditions, conditionTree.Order, "row")
	sortedReadoutDist, _ := hclust.Sort(normalizedReadoutDist, data.Readouts, readoutTree.Order, "column")
	sortedReadoutDist, _ = hclust.Sort(sortedReadoutDist, data.Readouts, readoutTree.Order, "row")
	sortedRatios := NormalizeMatrix(sortedAbundance)
	sortedScores, _ := hclust.Sort(data.Score, data.Conditions, conditionTree.Order, "column")
	sortedScores, _ = hclust.Sort(sortedScores, data.Readouts, readoutTree.Order, "row")

	// Output svgs.

	// Write condition-condition svg.
	if dataset.Parameters.WriteDistance {
		SvgCC(sortedConditionDist, conditionTree.Order, dataset.Parameters.FillColor)

		// Write minimap.
		MinimapDistance(
			sortedConditionDist,
			conditionTree.Order,
			dataset.Parameters.FillColor,
			"condition-condition",
		)

		// Write distance legend.
		legendTitle := fmt.Sprintf("Distance - %s", dataset.Parameters.Abundance)
		distanceLegend := svg.Gradient(dataset.Parameters.FillColor, legendTitle, 101, 0, dataset.Parameters.AbundanceCap)
		afero.WriteFile(fs.Instance, "svg/distance-legend.svg", []byte(distanceLegend), 0644)
	}

	// Write condition-readout dotplot.
	if dataset.Parameters.WriteDotplot {
		SvgDotplot(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			conditionTree.Order,
			readoutTree.Order,
			false,
			dataset.Parameters,
		)

		// Write minimap.
		Minimap(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			conditionTree.Order,
			readoutTree.Order,
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
			conditionTree.Order,
			readoutTree.Order,
			dataset.Parameters.FillColor,
			dataset.Parameters.AbundanceCap,
			false,
		)
	}

	// Write readout-readout svg.
	if dataset.Parameters.WriteDistance {
		SvgRR(sortedReadoutDist, readoutTree.Order, dataset.Parameters.FillColor)

		// Write minimap.
		MinimapDistance(
			sortedReadoutDist,
			readoutTree.Order,
			dataset.Parameters.FillColor,
			"readout-readout",
		)
	}

	// Create pdfs from svg.
	svgList := make([]string, 0)
	svgMiniList := make([]string, 0)
	if dataset.Parameters.WriteDistance {
		svgList = append(svgList, "condition-condition.svg")
		svgList = append(svgList, "distance-legend.svg")
		svgList = append(svgList, "readout-readout.svg")
		svgMiniList = append(svgMiniList, "condition-condition.svg")
		svgMiniList = append(svgMiniList, "readout-readout.svg")
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

	// Write condition-condition cytoscape file.
	WriteBBCytoscape(conditionDist, data.Conditions)

	// Write condition-readout cytoscape file.
	WriteBPCytoscape(dataset)

	// Write readout-readout cytoscape file.
	WritePPCytoscape(readoutDist, data.Readouts)

	// Output other files.

	// Write transformed data matrix to file.
	WriteMatrix(sortedAbundance, conditionTree.Order, readoutTree.Order, "other/data-transformed.txt")

	// Write transformed data matrix to file but as ratios instead of absolutes.
	WriteMatrix(sortedRatios, conditionTree.Order, readoutTree.Order, "other/data-transformed-ratios.txt")

	// Write newick trees to files.
	afero.WriteFile(fs.Instance, "other/condition-dendrogram.txt", []byte(conditionTree.Newick), 0644)
	afero.WriteFile(fs.Instance, "other/readout-dendrogram.txt", []byte(readoutTree.Newick), 0644)

	// Create interactive files.
	if dataset.Parameters.WriteDistance {
		distanceParams := dataset.Parameters
		distanceParams.AbundanceCap = 1
		distanceParams.MinAbundance = 0
		json := InteractiveHeatmap(
			normalizedConditionDist,
			conditionTree.Order,
			conditionTree.Order,
			true,
			distanceParams,
			"minimap/condition-condition.png",
		)
		afero.WriteFile(fs.Instance, "interactive/condition-condition.json", []byte(json), 0644)
		json = InteractiveHeatmap(
			normalizedReadoutDist,
			readoutTree.Order,
			readoutTree.Order,
			true,
			distanceParams,
			"minimap/readout-readout.png",
		)
		afero.WriteFile(fs.Instance, "interactive/readout-readout.json", []byte(json), 0644)
	}
	if dataset.Parameters.WriteDotplot {
		json := InteractiveDotplot(
			sortedAbundance,
			sortedRatios,
			sortedScores,
			conditionTree.Order,
			readoutTree.Order,
			dataset.Parameters,
			"minimap/dotplot.png",
		)
		afero.WriteFile(fs.Instance, "interactive/dotplot.json", []byte(json), 0644)
	}

	return
}
