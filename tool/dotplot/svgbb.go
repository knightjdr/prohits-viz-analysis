package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// SvgCC draws a condition condition heatmap.
func SvgCC(dist [][]float64, conditions []string, fillColor string) {
	// Heatmap parameters.
	parameters := map[string]interface{}{
		"abundanceCap": float64(1),
		"colLabel":     "Conditions",
		"fillColor":    fillColor,
		"invertColor":  true,
		"rowLabel":     "Conditions",
	}
	heatmap := svg.Heatmap(dist, typedef.Annotations{}, typedef.Markers{}, conditions, conditions, parameters)
	afero.WriteFile(fs.Instance, "svg/condition-condition.svg", []byte(heatmap), 0644)
	return
}
