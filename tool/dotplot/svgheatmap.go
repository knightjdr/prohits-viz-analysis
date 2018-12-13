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
	fillColor string,
	abundanceCap float64,
	invertColor bool,
) {
	parameters := typedef.Parameters{
		AbundanceCap: abundanceCap,
		Condition:    "Conditions",
		FillColor:    fillColor,
		InvertColor:  invertColor,
		Readout:      "Readouts",
	}
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
