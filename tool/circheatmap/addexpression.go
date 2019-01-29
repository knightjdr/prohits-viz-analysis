package circheatmap

import "github.com/knightjdr/prohits-viz-analysis/typedef"

func addExpression(
	data map[string]map[string]map[string]float64,
	readoutNames map[string]bool,
	metrics map[string]string,
	parameters typedef.Parameters,
) (map[string]map[string]map[string]float64, map[string]string) {
	if len(parameters.Tissues) > 0 {
		// Add tissue names to readout metrics
		metricsWithTissues := metrics
		for _, tissue := range parameters.Tissues {
			metricsWithTissues[tissue] = "RNA expression " + tissue
		}

		// Get expression data for readouts
		expression := parseTissues(readoutNames, parameters.TissueFile, parameters.Tissues)

		// Add expression data to condition data.
		conditionData := make(map[string]map[string]map[string]float64, len(data))
		for name := range data {
			conditionData[name] = make(map[string]map[string]float64)
		}
		for condition, readouts := range data {
			for readout := range readouts {
				conditionData[condition][readout] = data[condition][readout]
				for _, tissue := range parameters.Tissues {
					if expression[readout][tissue] > 0 {
						conditionData[condition][readout][tissue] = expression[readout][tissue]
					} else {
						conditionData[condition][readout][tissue] = 0
					}
				}
			}
		}
		return conditionData, metricsWithTissues
	}
	return data, metrics
}
