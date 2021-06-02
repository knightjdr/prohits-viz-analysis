package scv

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createInteractive(plots []types.CircHeatmap, legend types.CircHeatmapLegend, settings types.Settings) {

	interactiveData := &interactive.CircHeatmapData{
		Filename:   "interactive/scv.json",
		Legend:     legend,
		Parameters: settings,
		Plots:      plots,
		Settings: map[string]interface{}{
			"sortByKnown": shownKnownElement(settings.Known),
		},
	}
	interactive.CreateCircHeatmap(interactiveData)
}

func shownKnownElement(known string) bool {
	if known != "" {
		return true
	}
	return false
}
