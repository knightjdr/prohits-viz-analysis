package settings

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func addInferredSettings(csv []map[string]string, settings *types.Settings) {
	switch settings.Type {
	case "dotplot":
		inferDotplotSettings(csv, settings)
	case "heatmap":
		inferHeatmapSettings(csv, settings)
	}
}
