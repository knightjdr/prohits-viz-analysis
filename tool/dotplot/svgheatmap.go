package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// SvgHeatmap draws a condition readout heatmap.
func SvgHeatmap(
	matrix [][]float64,
	sortedColumns, sortedRows []string,
	userParams typedef.Parameters,
	invertColor bool,
) {
	parameters := userParams
	parameters.XLabel = userParams.Condition
	parameters.YLabel = userParams.Readout
	data := typedef.Matrices{
		Abundance: matrix,
	}
	heatmap := svg.Heatmap(
		"heatmap",
		data,
		typedef.Annotations{},
		typedef.Markers{},
		sortedColumns,
		sortedRows,
		false,
		parameters,
	)
	afero.WriteFile(fs.Instance, "svg/heatmap.svg", []byte(heatmap), 0644)
	return
}
