package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/spf13/afero"
)

// SvgBB draws a prey prey heatmap.
func SvgPP(dist [][]float64, sorted []string, colorSpace string) {
	// Heatmap params.
	params := map[string]interface{}{
		"colLabel":         "Preys",
		"colorSpace":       colorSpace,
		"invert":           true,
		"maximumAbundance": float64(1),
		"rowLabel":         "Preys",
	}
	heatmap := svg.Heatmap(dist, sorted, sorted, params)
	afero.WriteFile(fs.Instance, "svg/prey-prey.svg", []byte(heatmap), 0644)
	return
}
