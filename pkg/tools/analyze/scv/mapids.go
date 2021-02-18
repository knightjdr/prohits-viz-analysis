package scv

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/geneid"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func mapIDs(analysis *types.Analysis) map[string]map[string]string {
	mapped := map[string]map[string]string{
		"condition": {},
		"readout":   {},
	}

	if analysis.Settings.ConditionMapColumn != "" {
		mapped["condition"] = geneid.MapByColumn(analysis.Data, "condition", analysis.Settings.ConditionMapColumn)
	}

	if analysis.Settings.ReadoutMapColumn != "" {
		mapped["readout"] = geneid.MapByColumn(analysis.Data, "readout", analysis.Settings.ReadoutMapColumn)
	}

	return mapped
}
