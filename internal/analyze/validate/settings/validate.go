// Package settings validates user analysis settings.
package settings

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// Validate user analysis settings.
func Validate(analysis *types.Analysis) {
	switch analysis.Settings.Type {
	case "circheatmap":
		validateCircHeatmapSettings(analysis)
	case "correlation":
		validateCorrelationSettings(analysis)
	case "dotplot":
		validateDotplotSettings(analysis)
	default:
		log.WriteAndExit(fmt.Sprintf("Unknown analysis type: %s", analysis.Settings.Type))
	}
}
