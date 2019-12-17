// Package math defines common math operations.
package math

import (
	"math"
)

// Round will round to a specificed decimal based on the unit argument.
func Round(num, unit float64) (rounded float64) {
	rounded = math.Round(num/unit) * unit
	return
}
