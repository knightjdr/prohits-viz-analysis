// package math contains functions for performing matrix math
package math

import "math"

// AbsoluteValueEntries converts every value in a matrix to its absolute value.
func AbsoluteValueEntries(matrix [][]float64) [][]float64 {
	absoluteValue := make([][]float64, len(matrix))

	for i, row := range matrix {
		absoluteValue[i] = make([]float64, len(row))
		for j, value := range row {
			absoluteValue[i][j] = math.Abs(value)
		}
	}
	return absoluteValue
}
