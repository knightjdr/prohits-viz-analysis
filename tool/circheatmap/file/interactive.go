package file

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/image/interactive"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Interactive creates an interactive file for a circular heatmap.
func Interactive(
	plots []typedef.CircHeatmapPlot,
	userParameters typedef.Parameters,
	segmentSettings []typedef.CircHeatmapSetttings,
) {
	parameters := map[string]interface{}{
		"files":     helper.Filename(userParameters.Files),
		"imageType": "circ-heatmap",
	}
	settings := map[string]interface{}{
		"known": userParameters.Known,
		"plot":  0,
	}
	interactive.CircHeatmap(plots, parameters, settings, segmentSettings, "interactive/circheatmap.json")
}
