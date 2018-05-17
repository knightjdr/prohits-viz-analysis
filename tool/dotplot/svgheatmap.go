package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/spf13/afero"
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
	afero.WriteFile(fs.Instance, "svg/heatmap.svg", []byte(heatmap), 0644)
	return
}
