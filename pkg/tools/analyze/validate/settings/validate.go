// Package settings validates user analysis settings.
package settings

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Validate user analysis settings.
func Validate(analysis *types.Analysis) {
	switch analysis.Settings.Type {
	case "circheatmap":
		validateCircHeatmapSettings(analysis)
	case "condition-condition":
		validateCCSettings(analysis)
	case "correlation":
		validateCorrelationSettings(analysis)
	case "dotplot":
		validateDotplotSettings(analysis)
	case "specificity":
		validateSpecificitySettings(analysis)
	default:
		log.WriteAndExit(fmt.Sprintf("Unknown analysis type: %s", analysis.Settings.Type))
	}
}
