package cc

import (
	"fmt"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createInteractive(data []types.ScatterPoint, settings types.Settings) {
	plotPrefix := fmt.Sprintf("%[1]s-%[1]s", settings.Condition)
	scoreSymbols := createLegendScoreSymbol(settings.ScoreType, false)

	interactiveData := &interactive.ScatterData{
		AnalysisType: "condition-condition",
		Filename:     fmt.Sprintf("interactive/%s.json", plotPrefix),
		Legend: []map[string]string{
			{
				"color": "#0066cc",
				"text": fmt.Sprintf(
					"%s %s %s",
					settings.Score,
					scoreSymbols[0],
					strconv.FormatFloat(settings.PrimaryFilter, 'f', -1, 64),
				),
			},
			{
				"color": "#99ccff",
				"text": fmt.Sprintf(
					"%s %s %s %s %s",
					strconv.FormatFloat(settings.PrimaryFilter, 'f', -1, 64),
					scoreSymbols[1],
					settings.Score,
					scoreSymbols[2],
					strconv.FormatFloat(settings.SecondaryFilter, 'f', -1, 64),
				),
			},
		},
		Parameters: settings,
		Plots: []types.ScatterPlot{
			{
				Labels: types.ScatterAxesLabels{X: settings.ConditionX, Y: settings.ConditionY},
				Name:   plotPrefix,
				Points: data,
			},
		},
		Settings: map[string]interface{}{
			"logBase": settings.LogBase,
			"xFilter": settings.MinAbundance,
			"yFilter": settings.MinAbundance,
		},
	}

	interactive.CreateScatter(interactiveData)
}
