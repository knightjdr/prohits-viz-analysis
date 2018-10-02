package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// SvgPP draws a prey prey heatmap.
func SvgPP(dist [][]float64, sorted []string, fillColor string) {
	// Heatmap parameters.
	parameters := map[string]interface{}{
		"colLabel":     "Preys",
		"fillColor":    fillColor,
		"invertColor":  true,
		"abundanceCap": float64(1),
		"rowLabel":     "Preys",
	}
	heatmap := svg.Heatmap(dist, typedef.Annotations{}, typedef.Markers{}, sorted, sorted, parameters)
	afero.WriteFile(fs.Instance, "svg/prey-prey.svg", []byte(heatmap), 0644)
	return
}
