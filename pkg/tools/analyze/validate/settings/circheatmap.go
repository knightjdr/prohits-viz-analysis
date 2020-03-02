package settings

import "github.com/knightjdr/prohits-viz-analysis/pkg/types"

func validateCircHeatmapSettings(analysis *types.Analysis) {
	columnMap := validateFileSettings(analysis.Settings)
	analysis.Columns = appendOtherAbundanceColumns(columnMap, analysis.Settings.OtherAbundance)
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
