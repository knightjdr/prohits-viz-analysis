package specificity

import (
	"sort"

	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createInteractive(data map[string][]types.ScatterPoint, settings types.Settings) {
	plots := make([]types.ScatterPlot, len(data))

	conditions := make([]string, len(data))
	i := 0
	for condition := range data {
		conditions[i] = condition
		i++
	}
	sort.Strings(conditions)

	i = 0
	for _, condition := range conditions {
		plots[i] = types.ScatterPlot{
			Labels: types.ScatterAxesLabels{X: settings.Abundance, Y: "Specificity"},
			Name:   condition,
			Points: data[condition],
		}
		i++
	}

	interactiveData := &interactive.ScatterData{
		AnalysisType: "specificity",
		Filename:     "interactive/specificity.json",
		Legend: []map[string]string{
			{
				"color": "#6e97ff",
				"text":  "Infinite specificity",
			},
			{
				"color": "#dfcd06",
				"text":  "Finite specificity",
			},
		},
		Parameters: settings,
		Plots:      plots,
		Settings: map[string]interface{}{
			"logBase": settings.LogBase,
			"xFilter": settings.MinAbundance,
			"yFilter": 0,
		},
	}

	interactive.CreateScatter(interactiveData)
}
