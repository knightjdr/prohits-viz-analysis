package transform

import "math"

// Round will round to a specificed number of decimals
func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}
