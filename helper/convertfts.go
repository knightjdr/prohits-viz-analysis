package helper

import "strconv"

// ConvertFts converts a slice of float64 to strings with the requested precision.
// Use -1 for precision if none required.
func ConvertFts(floatSlice []float64, prec int) (stringSlice []string) {
	stringSlice = make([]string, len(floatSlice))
	for i, value := range floatSlice {
		stringSlice[i] = strconv.FormatFloat(value, 'f', prec, 64)
	}
	return stringSlice
}
