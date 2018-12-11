package transform

import (
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// ControlSubtraction will subtract the average control value from the readout abundance.
func ControlSubtraction(data []map[string]string, controlColumn string) (transformed []map[string]string) {
	transformed = data
	// Skip if no control column is specified.
	if controlColumn == "" {
		return
	}

	// Iterate over data slice and subtract control average from readout abundance.
	for _, row := range transformed {
		// Calculate control average.
		controls := strings.Split(row["control"], "|")
		var controlSum float64
		controlSum = 0
		for _, controlValue := range controls {
			valueAsFloat, _ := strconv.ParseFloat(controlValue, 64)
			controlSum += valueAsFloat
		}
		controlAvg := controlSum / float64(len(controls))
		// Subtract control average from each abundance value.
		abundance := strings.Split(row["abundance"], "|")
		transformedAbdStr := make([]string, 0) // Store as strings for joining.
		for _, abdValue := range abundance {
			transformedAbd, _ := strconv.ParseFloat(abdValue, 64)
			transformedAbd -= controlAvg
			if transformedAbd < 0 {
				transformedAbd = float64(0)
			} else {
				transformedAbd = helper.Round(transformedAbd, 0.01) // Round to nearest two decimals.
			}
			// Convert float to string and append.
			transformedAbdStr = append(transformedAbdStr, helper.FloatToString(transformedAbd))
		}
		row["abundance"] = strings.Join(transformedAbdStr[:], "|")
	}
	return
}
