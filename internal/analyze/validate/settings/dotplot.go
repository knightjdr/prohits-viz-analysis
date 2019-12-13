package settings

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func validateDotplotSettings(settings interface{}) (map[string]string, interface{}) {
	columnMap := validateFileSettings(settings.(*types.Dotplot).File)
	return columnMap, settings
}
