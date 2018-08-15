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
	fillColor string,
	maximumAbundance float64,
	invert bool,
) {
	params := map[string]interface{}{
		"colLabel":         "Baits",
		"fillColor":        fillColor,
		"invert":           invert,
		"maximumAbundance": maximumAbundance,
		"rowLabel":         "Preys",
	}
	heatmap := svg.Heatmap(matrix, sortedColumns, sortedRows, params)
	afero.WriteFile(fs.Instance, "svg/heatmap.svg", []byte(heatmap), 0644)
	return
}
