package transform

import (
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// SummedNormalization normalizes readout abundances between conditions by summing the
// readout abundances for each condition, taking the median of that and using the
// median / total readout abundance for each condition as the readout multplier.
func SummedNormalization(data []map[string]string) (transformed []map[string]string) {
	transformed = data
	// Sum thhe readout abundances for each condition.
	conditions := map[string]float64{}
	for _, row := range transformed {
		// Add condition to map if it is not present.
		conditionName := row["condition"]
		if _, ok := conditions[conditionName]; !ok {
			conditions[conditionName] = 0
		}
		// Add readout abundance.
		abundance := strings.Split(row["abundance"], "|")
		var abundanceSum float64
		abundanceSum = 0
		for _, abdValue := range abundance {
			abundanceFloat, _ := strconv.ParseFloat(abdValue, 64)
			abundanceSum += abundanceFloat
		}
		conditions[conditionName] += abundanceSum
	}

	// Get total abundance median.
	conditionValues := make([]float64, 0)
	for _, value := range conditions {
		conditionValues = append(conditionValues, value)
	}
	median := helper.MedianFloat(conditionValues)

	// Calculate readout multipliers.
	multiplier := map[string]float64{}
	for condition, value := range conditions {
		abundance := value
		if abundance == 0 {
			multiplier[condition] = 1
		} else {
			multiplier[condition] = median / abundance
		}
	}

	// Transform readouts.
	for _, row := range transformed {
		abundance := strings.Split(row["abundance"], "|")
		conditionName := row["condition"]
		transformedAbdStr := make([]string, 0) // Store as strings for joining.
		for _, abdValue := range abundance {
			transformedAbd, _ := strconv.ParseFloat(abdValue, 64)
			transformedAbd *= multiplier[conditionName]
			transformedAbd = helper.Round(transformedAbd, 0.01) // Round to nearest two decimals.
			// Convert float to string and append.
			transformedAbdStr = append(transformedAbdStr, helper.FloatToString(transformedAbd))
		}
		row["abundance"] = strings.Join(transformedAbdStr[:], "|")
	}

	return transformed
}
