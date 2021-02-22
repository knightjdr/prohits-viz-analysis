package scv

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/geneid"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func mapIDs(analysis *types.Analysis) map[string]map[string]string {
	mapped := map[string]map[string]string{
		"condition": {},
		"readout":   {},
	}

	if analysis.Settings.ConditionMapColumn != "" {
		mapped["condition"] = geneid.MapByColumn(analysis.Data, "condition", analysis.Settings.ConditionMapColumn, false)
	} else if analysis.Settings.ConditionMapFile != "" {
		mapped["condition"] = geneid.MapByFile(analysis.Data, "condition", fmt.Sprintf("helper-files/%s", analysis.Settings.ConditionMapFile))
	} else {
		mapped["condition"] = geneid.MapByColumn(analysis.Data, "condition", "condition", true)
	}
	settings := &geneid.HGNCsettings{
		File: analysis.Settings.GeneFile,
	}
	mapped["condition"], settings = geneid.MapToHGNC(mapped["condition"], analysis.Settings.ConditionIDType, settings)

	if analysis.Settings.ReadoutMapColumn != "" {
		mapped["readout"] = geneid.MapByColumn(analysis.Data, "readout", analysis.Settings.ReadoutMapColumn, false)
	} else if analysis.Settings.ReadoutMapFile != "" {
		mapped["readout"] = geneid.MapByFile(analysis.Data, "readout", fmt.Sprintf("helper-files/%s", analysis.Settings.ReadoutMapFile))
	} else {
		mapped["readout"] = geneid.MapByColumn(analysis.Data, "readout", "readout", true)
	}
	mapped["readout"], _ = geneid.MapToHGNC(mapped["readout"], analysis.Settings.ReadoutIDType, settings)

	return mapped
}
