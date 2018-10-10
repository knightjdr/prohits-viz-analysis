package transform

import (
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// ReadoutLength adjusts readout abundance based on the length of a readout relative to
// all others in the data set. It takes the median length of all readouts (this is
// after filtering by score in previous steps), calculates the median and then
// divides the median by each readout's length to get the multiplier for that readout.
func ReadoutLength(data []map[string]interface{}, readoutColumn string) (transformed []map[string]interface{}) {
	transformed = data
	// Skip if no readout length column is specified.
	if readoutColumn == "" {
		return
	}

	// Get all unique readouts and their lengths.
	readouts := map[string]int{}
	readoutLengths := make([]int, 0)
	for _, row := range transformed {
		readoutName := row["readout"].(string)
		if _, ok := readouts[readoutName]; !ok {
			length, _ := strconv.Atoi(row["readoutLength"].(string))
			readouts[readoutName] = length
			readoutLengths = append(readoutLengths, length)
		}
	}

	// Calculate median readout length and readout multiplier for each readout.
	median := MedianInt(readoutLengths)
	multiplier := map[string]float64{}
	for readout, length := range readouts {
		multiplier[readout] = median / float64(length)
	}

	// Iterate over data slice and multiply readout abundance by multiplier.
	for _, row := range transformed {
		readoutName := row["readout"].(string)
		abundance := strings.Split(row["abundance"].(string), "|")
		transformedAbdStr := make([]string, 0) // Store as strings for joining.
		for _, abdValue := range abundance {
			transformedAbd, _ := strconv.ParseFloat(abdValue, 64)
			transformedAbd *= multiplier[readoutName]
			transformedAbd = helper.Round(transformedAbd, 0.01) // Round to nearest two decimals.
			// Convert float to string and append.
			transformedAbdStr = append(transformedAbdStr, FloatToString(transformedAbd))
		}
		row["abundance"] = strings.Join(transformedAbdStr[:], "|")
	}

	return
}
