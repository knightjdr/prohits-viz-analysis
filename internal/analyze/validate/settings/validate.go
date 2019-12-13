// Package settings validates user analysis settings.
package settings

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// Validate user analysis settings.
func Validate(analysis types.Analysis) (map[string]string, interface{}) {
	var columnMap map[string]string
	var validatedSettings interface{}

	switch analysis.Type {
	case "circheatmap":
		columnMap, validatedSettings = validateCircHeatmapSettings(analysis.Settings)
	case "dotplot":
		columnMap, validatedSettings = validateDotplotSettings(analysis.Settings)
	default:
		log.WriteAndExit(fmt.Sprintf("Unknown anaylsis type: %s", analysis.Type))
	}

	return columnMap, validatedSettings
}
