package transform

import (
	"strconv"
	"strings"
)

// SummedNormalization normalizes prey abundances between baits by summing the
// prey abundances for each bait, taking the median of that and using the
// median / total prey abundance for each bait as the prey multplier.
func SummedNormalization(data []map[string]interface{}) (transformed []map[string]interface{}) {
	transformed = data
	// Sum thhe prey abundances for each bait.
	baits := map[string]float64{}
	for _, row := range transformed {
		// Add bait to map if it is not present.
		baitName := row["bait"].(string)
		if _, ok := baits[baitName]; !ok {
			baits[baitName] = 0
		}
		// Add prey abundance.
		abundance := strings.Split(row["abundance"].(string), "|")
		var abundanceSum float64
		abundanceSum = 0
		for _, abdValue := range abundance {
			abundanceFloat, _ := strconv.ParseFloat(abdValue, 64)
			abundanceSum += abundanceFloat
		}
		baits[baitName] += abundanceSum
	}

	// Get total abundance median.
	baitValues := make([]float64, 0)
	for _, value := range baits {
		baitValues = append(baitValues, value)
	}
	median := MedianFloat(baitValues)

	// Calculate prey multipliers.
	multiplier := map[string]float64{}
	for bait, value := range baits {
		abundance := value
		multiplier[bait] = median / abundance
	}

	// Transform preys.
	for _, row := range transformed {
		abundance := strings.Split(row["abundance"].(string), "|")
		baitName := row["bait"].(string)
		transformedAbdStr := make([]string, 0) // Store as strings for joining.
		for _, abdValue := range abundance {
			transformedAbd, _ := strconv.ParseFloat(abdValue, 64)
			transformedAbd *= multiplier[baitName]
			transformedAbd = Round(transformedAbd, 0.01) // Round to nearest two decimals.
			// Convert float to string and append.
			transformedAbdStr = append(transformedAbdStr, FloatToString(transformedAbd))
		}
		row["abundance"] = strings.Join(transformedAbdStr[:], "|")
	}

	return transformed
}
