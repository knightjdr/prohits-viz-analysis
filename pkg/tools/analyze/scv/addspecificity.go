package scv

import (
	"math"

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
			hasInfiniteValue := false
			max := 0.0

			for readout := range conditionData {
				data[condition][readout]["Specificity"] = specificityData[condition][readout]["specificity"]
				isInf := math.IsInf(data[condition][readout]["Specificity"], 1)
				if isInf {
					hasInfiniteValue = true
				}
				if !isInf && data[condition][readout]["Specificity"] > max {
					max = data[condition][readout]["Specificity"]
				}
			}
			if hasInfiniteValue {
				for readout := range conditionData {
					if math.IsInf(data[condition][readout]["Specificity"], 1) {
						data[condition][readout]["Specificity"] = max
					}
				}
			}
		}
	}
}
