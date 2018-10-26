package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// SvgRR draws a readout readout heatmap.
func SvgRR(dist [][]float64, sorted []string, fillColor string) {
	// Heatmap parameters.
	parameters := map[string]interface{}{
		"colLabel":     "Readouts",
		"fillColor":    fillColor,
		"invertColor":  true,
		"abundanceCap": float64(1),
		"rowLabel":     "Readouts",
	}
	heatmap := svg.Heatmap(dist, typedef.Annotations{}, typedef.Markers{}, sorted, sorted, false, parameters)
	afero.WriteFile(fs.Instance, "svg/readout-readout.svg", []byte(heatmap), 0644)
	return
}
