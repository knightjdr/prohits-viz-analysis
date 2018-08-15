package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// SvgDotplot draws a bait prey heatmap.
func SvgDotplot(
	abundance, ratios, scores [][]float64,
	sortedColumns, sortedRows []string,
	invert bool,
	userParams typedef.Parameters,
) {
	// Define dotplot parameters.
	params := map[string]interface{}{
		"colLabel":         "Baits",
		"edgeColor":        userParams.EdgeColor,
		"fillColor":        userParams.FillColor,
		"invert":           invert,
		"maximumAbundance": userParams.MaximumAbundance,
		"primary":          userParams.PrimaryFilter,
		"rowLabel":         "Preys",
		"secondary":        userParams.SecondaryFilter,
		"scoreType":        userParams.ScoreType,
	}
	dotplot := svg.Dotplot(abundance, ratios, scores, sortedColumns, sortedRows, params)
	afero.WriteFile(fs.Instance, "svg/dotplot.svg", []byte(dotplot), 0644)
	return
}
