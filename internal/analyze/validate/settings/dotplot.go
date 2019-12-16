package settings

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func validateDotplotSettings(analysis *types.Analysis) {
	analysis.Columns = validateFileSettings(analysis.Settings)
}
