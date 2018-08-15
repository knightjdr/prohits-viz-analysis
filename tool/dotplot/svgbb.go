package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/spf13/afero"
)

// SvgBB draws a bait bait heatmap.
func SvgBB(dist [][]float64, baits []string, fillColor string) {
	// Heatmap params.
	params := map[string]interface{}{
		"colLabel":         "Baits",
		"fillColor":        fillColor,
		"invert":           true,
		"maximumAbundance": float64(1),
		"rowLabel":         "Baits",
	}
	heatmap := svg.Heatmap(dist, baits, baits, params)
	afero.WriteFile(fs.Instance, "svg/bait-bait.svg", []byte(heatmap), 0644)
	return
}
