package settings

import "github.com/knightjdr/prohits-viz-analysis/pkg/types"

func validateSCVSettings(analysis *types.Analysis) {
	columnMap := validateFileSettings(analysis.Settings)

	otherColumns := analysis.Settings.OtherAbundance

	if analysis.Settings.ConditionMapColumn != "" {
		otherColumns = append(otherColumns, analysis.Settings.ConditionMapColumn)
	}

	if analysis.Settings.ReadoutMapColumn != "" {
		otherColumns = append(otherColumns, analysis.Settings.ReadoutMapColumn)
	}

	analysis.Columns = appendOtherColumns(columnMap, otherColumns)
	analysis.Settings.ParsimoniousReadoutFiltering = true
}

func appendOtherColumns(columnMap map[string]string, otherColumns []string) map[string]string {
	appendedColumnMap := columnMap

	if len(otherColumns) > 0 {
		for _, column := range otherColumns {
			appendedColumnMap[column] = column
		}
	}

	return appendedColumnMap
}
