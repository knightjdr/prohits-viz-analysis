package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// SvgCC draws a condition condition heatmap.
func SvgCC(dist [][]float64, conditions []string, userParams typedef.Parameters) {
	// Heatmap parameters.
	parameters := userParams
	parameters.AbundanceCap = float64(1)
	parameters.XLabel = userParams.Condition
	parameters.InvertColor = true
	parameters.YLabel = userParams.Condition

	data := typedef.Matrices{
		Abundance: dist,
	}
	heatmap := svg.Heatmap(
		"heatmap",
		data,
		typedef.Annotations{},
		typedef.Markers{},
		conditions,
		conditions,
		false,
		parameters,
	)
	afero.WriteFile(fs.Instance, "svg/condition-condition.svg", []byte(heatmap), 0644)
	return
}
