package settings

import "github.com/knightjdr/prohits-viz-analysis/pkg/types"

func convert(fileSettings *jsonSettings) types.Settings {
	return types.Settings{
		Abundance:       fileSettings.AbundanceColumn,
		InvertColor:     invertColorToBool(fileSettings.InvertColor),
		PrimaryFilter:   fileSettings.PrimaryFilter,
		Score:           fileSettings.ScoreColumn,
		ScoreType:       scoreTypeToBool(fileSettings.ScoreType),
		SecondaryFilter: fileSettings.SecondaryFilter,
		Type:            fileSettings.Type,
		XLabel:          fileSettings.XLabel,
		YLabel:          fileSettings.YLabel,
	}
}

func invertColorToBool(invertColor int) bool {
	if invertColor == 1 {
		return true
	}
	return false
}

func scoreTypeToBool(scoreType int) string {
	if scoreType == 1 {
		return "gte"
	}
	return "lte"
}
