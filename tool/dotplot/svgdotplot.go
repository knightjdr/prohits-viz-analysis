package dotplot

import (
	"io/ioutil"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// SvgDotplot draws a bait prey heatmap.
func SvgDotplot(
	abundance, ratios, scores [][]float64,
	columns, rows, sortedColumns, sortedRows []string,
	userParams typedef.Parameters,
) {
	// Sort matrices.
	sortedAbundance, _ := hclust.Sort(abundance, columns, sortedColumns, "column")
	sortedAbundance, _ = hclust.Sort(sortedAbundance, rows, sortedRows, "row")
	sortedRatios, _ := hclust.Sort(ratios, columns, sortedColumns, "column")
	sortedRatios, _ = hclust.Sort(sortedRatios, rows, sortedRows, "row")
	sortedScores, _ := hclust.Sort(scores, columns, sortedColumns, "column")
	sortedScores, _ = hclust.Sort(sortedScores, rows, sortedRows, "row")

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
	dotplot := svg.Dotplot(sortedAbundance, sortedRatios, sortedScores, sortedColumns, sortedRows, params)
	ioutil.WriteFile("svg/dotplot.svg", []byte(dotplot), 0644)
	return
}
