package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// Minimap draws a minimap for the dotplot.
func Minimap(
	abundance, ratios, scores [][]float64,
	sortedColumns, sortedRows []string,
	invertColor bool,
	userParams typedef.Parameters,
) {
	// Define dotplot parameters.
	parameters := typedef.Parameters{
		AbundanceCap:    userParams.AbundanceCap,
		EdgeColor:       userParams.EdgeColor,
		FillColor:       userParams.FillColor,
		InvertColor:     invertColor,
		PrimaryFilter:   userParams.PrimaryFilter,
		SecondaryFilter: userParams.SecondaryFilter,
		ScoreType:       userParams.ScoreType,
	}
	heatmap := typedef.Matrices{
		Abundance: abundance,
		Ratio:     ratios,
		Score:     scores,
	}
	minimap := svg.Heatmap(
		"dotplot",
		heatmap,
		typedef.Annotations{},
		typedef.Markers{},
		sortedColumns,
		sortedRows,
		true,
		parameters,
	)
	afero.WriteFile(fs.Instance, "minimap/dotplot.svg", []byte(minimap), 0644)
	return
}
