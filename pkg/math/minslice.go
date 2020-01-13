package math

import (
	"errors"
	"math"
)

// MinSliceFloat calculates the minimum value in a []float64
func MinSliceFloat(slice []float64) float64 {
	min := math.MaxFloat64
	for _, value := range slice {
		if value < min {
			min = value
		}
	}
	return min
}

// MinSliceInt calculates the minimum value in a []float64
func MinSliceInt(slice []int) (min int, err error) {
	if len(slice) == 0 {
		return 0, errors.New("slice has length 0")
	}

	min = math.MaxInt64
	for _, value := range slice {
		if value < min {
			min = value
		}
	}
	return
}
