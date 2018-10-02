package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// SvgBB draws a bait bait heatmap.
func SvgBB(dist [][]float64, baits []string, fillColor string) {
	// Heatmap parameters.
	parameters := map[string]interface{}{
		"abundanceCap": float64(1),
		"colLabel":     "Baits",
		"fillColor":    fillColor,
		"invertColor":  true,
		"rowLabel":     "Baits",
	}
	heatmap := svg.Heatmap(dist, typedef.Annotations{}, typedef.Markers{}, baits, baits, parameters)
	afero.WriteFile(fs.Instance, "svg/bait-bait.svg", []byte(heatmap), 0644)
	return
}
