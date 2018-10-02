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
	parameters := map[string]interface{}{
		"colLabel":     "Conditions",
		"edgeColor":    userParams.EdgeColor,
		"fillColor":    userParams.FillColor,
		"invertColor":  invertColor,
		"abundanceCap": userParams.AbundanceCap,
		"primary":      userParams.PrimaryFilter,
		"rowLabel":     "Readouts",
		"secondary":    userParams.SecondaryFilter,
		"scoreType":    userParams.ScoreType,
	}
	dotplot := svg.Dotplot(abundance, ratios, scores, typedef.Annotations{}, typedef.Markers{}, sortedColumns, sortedRows, parameters)
	afero.WriteFile(fs.Instance, "svg/dotplot.svg", []byte(dotplot), 0644)
	return
}
