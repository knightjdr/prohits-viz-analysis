package dotplot

import (
	"io/ioutil"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/svg"
)

// SvgHeatmap draws a bait prey heatmap.
func SvgHeatmap(
	matrix [][]float64,
	columns, rows, sortedColumns, sortedRows []string,
	colorSpace string,
	maximumAbundance float64,
) {
	// Sort matrix.
	sortedMatrix, _ := hclust.Sort(matrix, columns, sortedColumns, "column")
	sortedMatrix, _ = hclust.Sort(sortedMatrix, rows, sortedRows, "row")
	params := map[string]interface{}{
		"colLabel":         "Baits",
		"colorSpace":       colorSpace,
		"invert":           false,
		"maximumAbundance": maximumAbundance,
		"rowLabel":         "Preys",
	}
	heatmap := svg.Heatmap(sortedMatrix, sortedColumns, sortedRows, params)
	ioutil.WriteFile("svg/heatmap.svg", []byte(heatmap), 0644)
	return
}
