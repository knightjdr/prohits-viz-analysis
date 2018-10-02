package transform

import (
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// LogTransform log transforms readout abundance values.
func LogTransform(data []map[string]interface{}, base string) (transformed []map[string]interface{}) {
	transformed = data
	// Skip if log transformation not required.
	validLogs := map[string]bool{
		"2":  true,
		"e":  true,
		"10": true,
	}
	if _, ok := validLogs[base]; !ok {
		return
	}

	// Get log function.
	logFunc := LogFunc(base)

	// Iterate over data slice and log transform readout abundance.
	for _, row := range transformed {
		abundance := strings.Split(row["abundance"].(string), "|")
		transformedAbdStr := make([]string, 0) // Store as strings for joining.
		for _, abdValue := range abundance {
			transformedAbd, _ := strconv.ParseFloat(abdValue, 64)
			transformedAbd = logFunc(transformedAbd)
			if transformedAbd < 0 {
				transformedAbd = float64(0)
			} else {
				transformedAbd = helper.Round(transformedAbd, 0.01) // Round to nearest two decimals.
			}
			// Convert float to string and append.
			transformedAbdStr = append(transformedAbdStr, FloatToString(transformedAbd))
		}
		row["abundance"] = strings.Join(transformedAbdStr[:], "|")
	}
	return
}
