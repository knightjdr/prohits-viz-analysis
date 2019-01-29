package circheatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

func addKnown(
	data map[string]map[string]map[string]float64,
	conditionMapping map[string]string,
	parameters typedef.Parameters,
) map[string]map[string]map[string]float64 {
	if parameters.Known {
		// Get known readout data
		knownReadouts := parseKnownReadouts(conditionMapping, parameters.KnownFile, parameters.Species)

		conditionData := make(map[string]map[string]map[string]float64, len(data))
		for name := range data {
			conditionData[name] = make(map[string]map[string]float64)
		}

		// Add known status to readouts
		for condition, readouts := range data {
			for readout := range readouts {
				conditionData[condition][readout] = data[condition][readout]
				if knownReadouts[condition][readout] {
					conditionData[condition][readout]["known"] = 1
				} else {
					conditionData[condition][readout]["known"] = 0
				}
			}
		}
		return conditionData
	}
	return data
}
