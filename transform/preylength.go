package transform

import (
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// PreyLength adjusts prey abundance based on the length of a prey relative to
// all others in the data set. It takes the median length of all preys (this is
// after filtering by score in previous steps), calculates the median and then
// divides the median by each prey's length to get the multiplier for that prey.
func PreyLength(data []map[string]interface{}, preyColumn string) (transformed []map[string]interface{}) {
	transformed = data
	// Skip if no prey length column is specified.
	if preyColumn == "" {
		return
	}

	// Get all unique preys and their lengths.
	preys := map[string]int{}
	preyLengths := make([]int, 0)
	for _, row := range transformed {
		preyName := row["prey"].(string)
		if _, ok := preys[preyName]; !ok {
			length, _ := strconv.Atoi(row["preyLength"].(string))
			preys[preyName] = length
			preyLengths = append(preyLengths, length)
		}
	}

	// Calculate median prey length and prey multiplier for each prey.
	median := MedianInt(preyLengths)
	multiplier := map[string]float64{}
	for prey, length := range preys {
		multiplier[prey] = median / float64(length)
	}

	// Iterate over data slice and multiply prey abundance by multiplier.
	for _, row := range transformed {
		preyName := row["prey"].(string)
		abundance := strings.Split(row["abundance"].(string), "|")
		transformedAbdStr := make([]string, 0) // Store as strings for joining.
		for _, abdValue := range abundance {
			transformedAbd, _ := strconv.ParseFloat(abdValue, 64)
			transformedAbd *= multiplier[preyName]
			transformedAbd = helper.Round(transformedAbd, 0.01) // Round to nearest two decimals.
			// Convert float to string and append.
			transformedAbdStr = append(transformedAbdStr, FloatToString(transformedAbd))
		}
		row["abundance"] = strings.Join(transformedAbdStr[:], "|")
	}

	return
}
