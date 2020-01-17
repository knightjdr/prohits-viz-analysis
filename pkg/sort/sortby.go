// Package sort contains functions for sorting slices.
package sort

// ByIndicesFloat sorts a []float64 by indices.
func ByIndicesFloat(s []float64, indices []int) []float64 {
	sorted := make([]float64, len(s))

	for i, index := range indices {
		sorted[i] = s[index]
	}

	return sorted
}

// ByIndicesInt sorts a []int by indices.
func ByIndicesInt(s []int, indices []int) []int {
	sorted := make([]int, len(s))

	for i, index := range indices {
		sorted[i] = s[index]
	}

	return sorted
}
