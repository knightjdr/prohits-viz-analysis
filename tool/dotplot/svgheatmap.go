package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// SvgHeatmap draws a bait prey heatmap.
func SvgHeatmap(
	matrix [][]float64,
	sortedColumns, sortedRows []string,
	fillColor string,
	abundanceCap float64,
	invertColor bool,
) {
	parameters := map[string]interface{}{
		"colLabel":     "Baits",
		"fillColor":    fillColor,
		"invertColor":  invertColor,
		"abundanceCap": abundanceCap,
		"rowLabel":     "Preys",
	}
	heatmap := svg.Heatmap(matrix, typedef.Annotations{}, typedef.Markers{}, sortedColumns, sortedRows, parameters)
	afero.WriteFile(fs.Instance, "svg/heatmap.svg", []byte(heatmap), 0644)
	return
}
