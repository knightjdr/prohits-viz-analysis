// Package normalize has functions for normalization.
package normalize

import "github.com/knightjdr/prohits-viz-analysis/pkg/math"

// Slice normalizes the values in a slice to one. The greatest value
// in the slice will be one and the other values will be divided by this.
func Slice(unnormalized []float64) (normalized []float64) {
	max := float64(0)
	for _, value := range unnormalized {
		if value > max {
			max = value
		}
	}

	normalized = make([]float64, len(unnormalized))
	for i, value := range unnormalized {
		normalized[i] = math.Round(value/max, 0.01)
	}
	return
}

// Matrix normalizes every row in a matrix to one.
func Matrix(matrix [][]float64) (normalized [][]float64) {
	for _, row := range matrix {
		normRow := Slice(row)
		normalized = append(normalized, normRow)
	}
	return
}
