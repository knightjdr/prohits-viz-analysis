package scv

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/specificity"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func addSpecificity(data map[string]map[string]map[string]float64, analysis *types.Analysis) {
	conditions := make([]string, len(data))
	i := 0
	for condition := range data {
		conditions[i] = condition
		i++
	}

	if analysis.Settings.Specificity && len(conditions) > 1 {
		specificityData := specificity.Calculate(analysis)

		for condition, conditionData := range data {
			for readout := range conditionData {
				data[condition][readout]["Specificity"] = specificityData[condition][readout]["specificity"]
			}
		}
	}
}
