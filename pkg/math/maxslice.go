package math

import "math"

// MaxIndexSliceFloat returns the index of the maximum value in a []float64
func MaxIndexSliceFloat(slice []float64) int {
	max := -math.MaxFloat64
	maxIndex := -1
	for i, value := range slice {
		if value > max {
			max = value
			maxIndex = i
		}
	}
	return maxIndex
}

// MaxSliceFloat calculates the maximum value in a []float64
func MaxSliceFloat(slice []float64) float64 {
	max := -math.MaxFloat64
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}

// MaxSliceInt calculates the maximum value in a []int
func MaxSliceInt(slice []int) int {
	max := math.MinInt64
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}
