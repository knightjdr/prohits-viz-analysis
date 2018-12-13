package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// SvgDotplot draws a condition readout heatmap.
func SvgDotplot(
	abundance, ratios, scores [][]float64,
	sortedColumns, sortedRows []string,
	invertColor bool,
	userParams typedef.Parameters,
) {
	// Define dotplot parameters.
	parameters := typedef.Parameters{
		AbundanceCap:    userParams.AbundanceCap,
		Condition:       "Conditions",
		EdgeColor:       userParams.EdgeColor,
		FillColor:       userParams.FillColor,
		InvertColor:     invertColor,
		PrimaryFilter:   userParams.PrimaryFilter,
		Readout:         "Readouts",
		SecondaryFilter: userParams.SecondaryFilter,
		ScoreType:       userParams.ScoreType,
	}
	data := typedef.Matrices{
		Abundance: abundance,
		Ratio:     ratios,
		Score:     scores,
	}
	dotplot := svg.Heatmap(
		"dotplot",
		data,
		typedef.Annotations{},
		typedef.Markers{},
		sortedColumns,
		sortedRows,
		false,
		parameters,
	)
	afero.WriteFile(fs.Instance, "svg/dotplot.svg", []byte(dotplot), 0644)
	return
}
