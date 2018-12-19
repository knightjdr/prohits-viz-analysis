package nocluster

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

func order(orderType string, dataset *typedef.Dataset, matrices *typedef.Matrices) (orderedNames []string) {
	var cluster string
	var cytoscapeFunc func([][]float64, []string)
	var filename string
	var inputList []string
	var label string
	var names []string
	if orderType == "condition" {
		cluster = dataset.Parameters.ConditionClustering
		cytoscapeFunc = file.WriteBBCytoscape
		filename = "condition-condition"
		inputList = dataset.Parameters.ConditionList
		label = dataset.Parameters.Condition
		names = matrices.Conditions
	} else {
		cluster = dataset.Parameters.ReadoutClustering
		cytoscapeFunc = file.WritePPCytoscape
		filename = "readout-readout"
		inputList = dataset.Parameters.ReadoutList
		label = dataset.Parameters.Readout
		names = matrices.Readouts
	}

	if cluster != "none" {
		// Generate distance matrix.
		dist := hclust.Distance(matrices.Abundance, dataset.Parameters.Distance, true)

		// Cluster.
		clust, err := hclust.Cluster(dist, dataset.Parameters.ClusteringMethod)
		logmessage.CheckError(err, true)

		// Optimize clustering.
		if dataset.Parameters.ClusteringOptimize {
			clust = hclust.Optimize(clust, dist, ignore)
		}

		// Create tree and get clustering order.
		tree, err := hclust.Tree(clust, names)
		logmessage.CheckError(err, true)

		// Set condition order.
		orderedNames = tree.Order

		// Normalize distance matrix to 1.
		normalizedDist := helper.NormalizeMatrix(dist)

		// Sort distance matrix.
		sortedDist, _ := hclust.Sort(normalizedDist, names, orderedNames, "column")
		sortedDist, _ = hclust.Sort(sortedDist, names, orderedNames, "row")

		// Distance images.
		if dataset.Parameters.WriteDistance {
			// svg
			file.Distance(sortedDist, orderedNames, label, filename, dataset.Parameters)

			// Write minimap.
			file.MinimapDistance(sortedDist, orderedNames, dataset.Parameters.FillColor, filename)

			// Generate pdfs and pngs.
			svgName := fmt.Sprintf("%s.svg", filename)
			if dataset.Parameters.Pdf {
				svg.ConvertPdf([]string{svgName})
			}
			if dataset.Parameters.Png {
				svg.ConvertPng([]string{svgName})
			}

			file.InteractiveDistance(sortedDist, orderedNames, label, filename, dataset.Parameters)
		}

		// Write condition-condition cytoscape file.
		cytoscapeFunc(dist, names)

		// Write newick tree to file.
		treeFilename := fmt.Sprintf("other/%s-dendrogram.txt", orderType)
		afero.WriteFile(fs.Instance, treeFilename, []byte(tree.Newick), 0644)
	} else {
		// Only keep user specified conditions that are actually in the file
		orderedNames = checkList(dataset.FileData, orderType, inputList)
	}

	return
}
