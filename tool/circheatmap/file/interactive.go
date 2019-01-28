package file

import (
	"github.com/knightjdr/prohits-viz-analysis/image/interactive"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Interactive creates an interactive file for a circular heatmap.
func Interactive(
	conditionOrder []string,
	data map[string]map[string]map[string]float64,
	parameters typedef.Parameters,
	readoutMetrics map[string]string,
) {
	settings := map[string]interface{}{
		"known": parameters.Known,
		"plot":  0,
	}
	interactive.CircHeatmap(conditionOrder, data, settings, readoutMetrics, "interactive/circheatmap.json")
}
