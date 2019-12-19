package math

import (
	goMath "math"
)

// Floor a float value to an int.
func Floor(value float64) int {
	return int(goMath.Floor(value))
}
