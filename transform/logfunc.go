package transform

import "math"

// LogFunc returns a logarithm function.
func LogFunc(base string) func(float64) float64 {
	if base == "2" {
		return func(num float64) float64 {
			if num <= 0 {
				return 0
			}
			return math.Log2(num)
		}
	} else if base == "10" {
		return func(num float64) float64 {
			if num <= 0 {
				return 0
			}
			return math.Log10(num)
		}
	} else {
		return func(num float64) float64 {
			if num <= 0 {
				return 0
			}
			return math.Log(num)
		}
	}
}
