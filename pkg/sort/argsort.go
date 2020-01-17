package sort

import "gonum.org/v1/gonum/floats"

// ArgsortFloat returns the indices after sorting an []float64.
func ArgsortFloat(x []float64) []int {
	n := len(x)
	xCopy := make([]float64, n)
	copy(xCopy, x)

	xIndices := make([]int, n)
	floats.Argsort(xCopy, xIndices)

	return xIndices
}

// ArgsortInt returns the indices after sorting an []int.
func ArgsortInt(x []int) []int {
	n := len(x)
	xCopy := make([]float64, n)
	for i, value := range x {
		xCopy[i] = float64(value)
	}

	xIndices := make([]int, n)
	floats.Argsort(xCopy, xIndices)

	return xIndices
}
