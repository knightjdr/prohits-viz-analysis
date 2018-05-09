package dotplot

import (
	"io/ioutil"

	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// SvgDotplot draws a bait prey heatmap.
func SvgDotplot(
	abundance, ratios, scores [][]float64,
	sortedColumns, sortedRows []string,
	userParams typedef.Parameters,
) {
	// Define dotplot parameters.
	params := map[string]interface{}{
		"colLabel":         "Baits",
		"colorSpace":       userParams.ColorSpace,
		"invert":           false,
		"maximumAbundance": userParams.MaximumAbundance,
		"primary":          userParams.PrimaryFilter,
		"rowLabel":         "Preys",
		"secondary":        userParams.SecondaryFilter,
		"scoreType":        userParams.ScoreType,
	}
	dotplot := svg.Dotplot(abundance, ratios, scores, sortedColumns, sortedRows, params)
	ioutil.WriteFile("svg/dotplot.svg", []byte(dotplot), 0644)
	return
}
