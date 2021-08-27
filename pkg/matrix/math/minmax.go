package math

import "math"

// MinMax returns the min and max values in a matrix
func MinMax(matrix [][]float64) (float64, float64) {
	max := math.Inf(-1)
	min := math.Inf(1)
	for _, row := range matrix {
		for _, value := range row {
			if value > max {
				max = value
			}
			if value < min {
				min = value
			}
		}
	}
	return min, max
}
