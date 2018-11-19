package helper

import "math"

// TruncateFloat truncates a floating point to specified precision
func TruncateFloat(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Round(num*output)) / output
}
