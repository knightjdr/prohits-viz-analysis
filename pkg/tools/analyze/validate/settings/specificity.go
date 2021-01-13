package settings

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func validateSpecificitySettings(analysis *types.Analysis) {
	analysis.Columns = validateFileSettings(analysis.Settings)
}
