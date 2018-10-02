package transform

import (
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// ReadoutNormalization normalizes readout abundances between conditions using a specific readout.
// For each condition it grabs the readout abundance and then the median is used to determine
// the readout multiplier for each condition. The multiplier is calculated as
// median / readout abundance for each condition.
func ReadoutNormalization(
	data []map[string]interface{},
	normalizationReadout string,
) (transformed []map[string]interface{}) {
	transformed = data
	// Grab the readout abundance for each condition.
	conditions := map[string]float64{}
	readoutValues := make([]float64, 0)
	for _, row := range transformed {
		// Add condition to map if it is not present.
		conditionName := row["condition"].(string)
		if _, ok := conditions[conditionName]; !ok {
			conditions[conditionName] = 0 // Use 0 to indicate missing value.
		}
		// If current readout matches readout for normalization set its condition value.
		readoutName := row["readout"].(string)
		if readoutName == normalizationReadout {
			abundance := strings.Split(row["abundance"].(string), "|")
			var abundanceSum float64
			abundanceSum = 0
			for _, abdValue := range abundance {
				abundanceFloat, _ := strconv.ParseFloat(abdValue, 64)
				abundanceSum += abundanceFloat
			}
			conditions[conditionName] = abundanceSum
			readoutValues = append(readoutValues, abundanceSum)
		}
	}

	// Get readout abundance median.
	median := MedianFloat(readoutValues)

	// Calculate readout multipliers.
	multiplier := map[string]float64{}
	for condition, value := range conditions {
		abundance := value
		if abundance == 0 {
			abundance = median
		}
		multiplier[condition] = median / abundance
	}

	// Transform readouts.
	for _, row := range transformed {
		abundance := strings.Split(row["abundance"].(string), "|")
		conditionName := row["condition"].(string)
		transformedAbdStr := make([]string, 0) // Store as strings for joining.
		for _, abdValue := range abundance {
			transformedAbd, _ := strconv.ParseFloat(abdValue, 64)
			transformedAbd *= multiplier[conditionName]
			transformedAbd = helper.Round(transformedAbd, 0.01) // Round to nearest two decimals.
			// Convert float to string and append.
			transformedAbdStr = append(transformedAbdStr, FloatToString(transformedAbd))
		}
		row["abundance"] = strings.Join(transformedAbdStr[:], "|")
	}

	return transformed
}
