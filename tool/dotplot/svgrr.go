package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// SvgRR draws a readout readout heatmap.
func SvgRR(dist [][]float64, sorted []string, userParams typedef.Parameters) {
	// Heatmap parameters.
	parameters := userParams
	parameters.AbundanceCap = float64(1)
	parameters.XLabel = userParams.Readout
	parameters.InvertColor = true
	parameters.YLabel = userParams.Readout

	data := typedef.Matrices{
		Abundance: dist,
	}
	heatmap := svg.Heatmap(
		"heatmap",
		data,
		typedef.Annotations{},
		typedef.Markers{},
		sorted,
		sorted,
		false,
		parameters,
	)
	afero.WriteFile(fs.Instance, "svg/readout-readout.svg", []byte(heatmap), 0644)
	return
}
