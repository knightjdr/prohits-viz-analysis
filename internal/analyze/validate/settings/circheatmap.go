package settings

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"

func validateCircHeatmapSettings(settings interface{}) (map[string]string, interface{}) {
	castSettings := settings.(*types.CircHeatmap)
	columnMap := validateFileSettings(castSettings.File)
	columnMap = appendOtherAbundanceColumns(columnMap, castSettings.OtherAbundance)
	return columnMap, settings
}

func appendOtherAbundanceColumns(columnMap map[string]string, otherAbundanceColumns []string) map[string]string {
	appendedColumnMap := columnMap

	if len(otherAbundanceColumns) > 0 {
		for _, column := range otherAbundanceColumns {
			appendedColumnMap[column] = column
		}
	}

	return appendedColumnMap
}
