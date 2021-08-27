package hierarchical

import (
	heatmapColor "github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/color"
	matrixMath "github.com/knightjdr/prohits-viz-analysis/pkg/matrix/math"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// AdjustSettings changes some input settings based on the values present in the data
// and defines some others.
func AdjustSettings(settings types.Settings, abundance [][]float64) types.Settings {
	adjusted := settings
	adjusted.AbundanceType = matrixMath.DefineValues(abundance)
	heatmapColor.SetFillLimits(&adjusted)
	heatmapColor.AdjustFillColor(&adjusted)
	return adjusted
}
