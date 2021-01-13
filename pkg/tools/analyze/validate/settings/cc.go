package settings

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func validateCCSettings(analysis *types.Analysis) {
	analysis.Columns = validateFileSettings(analysis.Settings)
}
