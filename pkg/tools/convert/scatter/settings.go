package scatter

import "github.com/knightjdr/prohits-viz-analysis/pkg/types"

func inferSettings(plotSettings *jsonSettings) types.Settings {
	settings := types.Settings{
		Score: plotSettings.Score,
		Type:  determineAnalysisType(plotSettings.Plot),
	}

	if plotSettings.Plot == "Binary-biDist" {
		settings.ScoreType = convertScoreTypeFromInt(plotSettings.ScoreType)
	}
	if plotSettings.Plot == "Specificity" {
		settings.Abundance = plotSettings.XLabel
	}

	return settings
}

func convertScoreTypeFromInt(scoreType int) string {
	if scoreType == 1 {
		return "gte"
	}
	return "lte"
}

func determineAnalysisType(plotType string) string {
	if plotType == "Binary" {
		return "condition-condition"
	}
	if plotType == "Binary-biDist" {
		return "condition-condition"
	}
	if plotType == "Specificity" {
		return "specificity"
	}
	return ""
}
