package scatter

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
)

func defineLegend(plotSettings *jsonSettings, format int) []map[string]string {
	if plotSettings.Plot == "Binary" {
		return []map[string]string{
			{
				"color": "#ff0000",
				"text":  "Infinite fold change",
			},
		}
	}
	if plotSettings.Plot == "Binary-biDist" {
		scoreSymbols := createLegendScoreSymbols(plotSettings.ScoreType)
		return []map[string]string{
			{
				"color": "#0066cc",
				"text": fmt.Sprintf(
					"%s %s %s",
					plotSettings.Score,
					scoreSymbols[0],
					float.RemoveTrailingZeros(plotSettings.PrimaryFilter),
				),
			},
			{
				"color": "#99ccff",
				"text": fmt.Sprintf(
					"%s %s %s %s %s",
					float.RemoveTrailingZeros(plotSettings.PrimaryFilter),
					scoreSymbols[1],
					plotSettings.Score,
					scoreSymbols[2],
					float.RemoveTrailingZeros(plotSettings.SecondaryFilter),
				),
			},
		}
	}
	if plotSettings.Plot == "Specificity" && format == 1 {
		return []map[string]string{
			{
				"color": "#ff0000",
				"text":  "Infinite specificity",
			},
		}
	}
	if plotSettings.Plot == "Specificity" && format == 2 {
		return []map[string]string{
			{
				"color": "#0066cc",
				"text":  "Infinite specificity",
			},
		}
	}
	return []map[string]string{}
}

func createLegendScoreSymbols(scoreType int) []string {
	if scoreType == 1 {
		return []string{"≥", ">", "≥"}
	}
	return []string{"≤", "<", "≤"}
}
