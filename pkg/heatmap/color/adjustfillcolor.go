// Package color has utilities for setting and adjusting heatmap colours.
package color

import "github.com/knightjdr/prohits-viz-analysis/pkg/types"

func AdjustFillColor(settings *types.Settings) {
	if settings.AutomaticallySetFill {
		settings.FillColor = "blue"
		settings.InvertColor = false
		if settings.AbundanceType == "bidirectional" {
			settings.FillColor = "blueRed"
		}
		if settings.AbundanceType == "negative" {
			settings.InvertColor = true
		}
	}
}
