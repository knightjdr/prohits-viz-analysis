package transform

import (
	"strconv"
	"strings"
)

// PreyNormalization normalizes prey abundances between baits using a specific prey.
// For each bait it grabs the prey abundance and then the median is used to determine
// the prey multiplier for each bait. The multiplier is calculated as
// median / prey abundance for each bait.
func PreyNormalization(
	data []map[string]interface{},
	normalizationPrey string,
) (transformed []map[string]interface{}) {
	transformed = data
	// Grab the prey abundance for each bait.
	baits := map[string]float64{}
	preyValues := make([]float64, 0)
	for _, row := range transformed {
		// Add bait to map if it is not present.
		baitName := row["bait"].(string)
		if _, ok := baits[baitName]; !ok {
			baits[baitName] = 0 // Use 0 to indicate missing value.
		}
		// If current prey matches prey for normalization set its bait value.
		preyName := row["prey"].(string)
		if preyName == normalizationPrey {
			abundance := strings.Split(row["abundance"].(string), "|")
			var abundanceSum float64
			abundanceSum = 0
			for _, abdValue := range abundance {
				abundanceFloat, _ := strconv.ParseFloat(abdValue, 64)
				abundanceSum += abundanceFloat
			}
			baits[baitName] = abundanceSum
			preyValues = append(preyValues, abundanceSum)
		}
	}

	// Get prey abundance median.
	median := MedianFloat(preyValues)

	// Calculate prey multipliers.
	multiplier := map[string]float64{}
	for bait, value := range baits {
		abundance := value
		if abundance == 0 {
			abundance = median
		}
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
