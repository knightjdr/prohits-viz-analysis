package main

import "strconv"

// Converts a csv of strings to interface type.
func toInterface(csv []map[string]string) []map[string]interface{} {
	converted := make([]map[string]interface{}, len(csv))
	for i, datum := range csv {
		datumInterface := make(map[string]interface{})
		datumInterface["abundance"], _ = strconv.ParseFloat(datum["abundance"], 64)
		datumInterface["condition"] = datum["condition"]
		datumInterface["readout"] = datum["readout"]
		if value, ok := datum["score"]; ok {
			datumInterface["score"], _ = strconv.ParseFloat(value, 64)
		} else {
			datumInterface["score"] = float64(0)
		}
		converted[i] = datumInterface
	}
	return converted
}
