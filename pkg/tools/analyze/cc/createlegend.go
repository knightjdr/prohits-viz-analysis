package cc

import (
	"fmt"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/scatter"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createLegend(settings types.Settings) {
	scoreSymbols := createLegendScoreSymbol(settings.ScoreType, true)

	legendPrefix := fmt.Sprintf("%[1]s-%[1]s", settings.Condition)
	legendPoints := []map[string]string{
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
	}

	legendData := scatter.Legend{
		Filename: fmt.Sprintf("svg/%s-legend.svg", legendPrefix),
		Points:   legendPoints,
		Title:    legendPrefix,
	}
	scatter.CreateLegend(legendData)

	if settings.Png {
		svg.ConvertToPNG(fmt.Sprintf("svg/%s-legend.svg", legendPrefix), fmt.Sprintf("png/%s-legend.png", legendPrefix), "white")
	}
}

func createLegendScoreSymbol(scoreType string, withEntities bool) []string {
	if scoreType == "gte" && withEntities {
		return []string{"&#8805;", "&gt;", "&#8805;"}
	}
	if scoreType == "gte" && !withEntities {
		return []string{"≥", ">", "≥"}
	}
	if scoreType == "lte" && withEntities {
		return []string{"&#8804;", "&lt;", "&#8804;"}
	}
	return []string{"≤", "<", "≤"}
}
