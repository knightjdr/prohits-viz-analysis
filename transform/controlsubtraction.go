package transform

import (
	"strconv"
	"strings"
)

// ControlSubtraction will subtract the average control value from the prey abundance
func ControlSubtraction(data []map[string]interface{}, control string) ([]map[string]interface{}, error) {
	var err error
	// if no control column is specified, skip
	if control == "" {
		return data, err
	}

	// iterate over data slice and subtract control average from prey abundance
	for _, row := range data {
		// calculate control average
		controls := strings.Split(row["control"].(string), "|")
		var controlSum float64
		controlSum = 0
		for _, controlValue := range controls {
			valueAsFloat, _ := strconv.ParseFloat(controlValue, 64)
			controlSum += valueAsFloat
		}
		controlAvg := controlSum / float64(len(controls))
		// subtract control average from each abundance value
		abundance := strings.Split(row["abundance"].(string), "|")
		transformedAbdStr := make([]string, 0) // will store as strings for joining
		for i, abdValue := range abundance {
			transformedAbd, _ := strconv.ParseFloat(abdValue, 64)
			transformedAbd -= controlAvg
			if transformedAbd < 0 {
				transformedAbd = 0
			}
			transformedAbd = Round(transformedAbd, 0.01)         // round to nearest two decimals
			transformedAbdStr[i] = FloatToString(transformedAbd) //convert float to string
		}
		row["abundance"] = strings.Join(transformedAbdStr[:], "|")
	}
	return data, err
}
