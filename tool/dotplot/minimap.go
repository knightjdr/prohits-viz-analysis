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
	parameters := map[string]interface{}{
		"abundanceCap": userParams.AbundanceCap,
		"edgeColor":    userParams.EdgeColor,
		"fillColor":    userParams.FillColor,
		"invertColor":  invertColor,
		"primary":      userParams.PrimaryFilter,
		"secondary":    userParams.SecondaryFilter,
		"scoreType":    userParams.ScoreType,
	}
	minimap := svg.Dotplot(abundance, ratios, scores, typedef.Annotations{}, typedef.Markers{}, sortedColumns, sortedRows, true, parameters)
	afero.WriteFile(fs.Instance, "minimap/dotplot.svg", []byte(minimap), 0644)
	return
}
