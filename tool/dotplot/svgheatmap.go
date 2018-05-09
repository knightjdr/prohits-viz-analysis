package dotplot

import (
	"io/ioutil"

	"github.com/knightjdr/prohits-viz-analysis/svg"
)

// SvgHeatmap draws a bait prey heatmap.
func SvgHeatmap(
	matrix [][]float64,
	sortedColumns, sortedRows []string,
	colorSpace string,
	maximumAbundance float64,
) {
	params := map[string]interface{}{
		"colLabel":         "Baits",
		"colorSpace":       colorSpace,
		"invert":           false,
		"maximumAbundance": maximumAbundance,
		"rowLabel":         "Preys",
	}
	heatmap := svg.Heatmap(matrix, sortedColumns, sortedRows, params)
	ioutil.WriteFile("svg/heatmap.svg", []byte(heatmap), 0644)
	return
}
