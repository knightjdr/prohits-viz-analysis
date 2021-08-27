package color

import "github.com/knightjdr/prohits-viz-analysis/pkg/types"

// SetFillLimits for a fill gradient in types.Settings
func SetFillLimits(settings *types.Settings) {
	if settings.AbundanceType == "bidirectional" {
		settings.FillMax = settings.AbundanceCap
		settings.FillMin = -settings.AbundanceCap
	}
	if settings.AbundanceType == "negative" {
		settings.FillMax = -settings.MinAbundance
		settings.FillMin = -settings.AbundanceCap
	}
	if settings.AbundanceType == "positive" {
		settings.FillMax = settings.AbundanceCap
		settings.FillMin = settings.MinAbundance
	}
}
