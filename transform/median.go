package transform

import "sort"

// MedianInt calulates the median of a slice of integers.
func MedianInt(x []int) (median float64) {
	// Sort the slice.
	sort.Ints(x)

	xLen := len(x)
	if xLen%2 == 0 {
		// Length is even so calculate the average between the two middle values.
		median = float64(x[xLen/2]+x[(xLen/2)-1]) / float64(2)
		return
	}
	// Length is odd so return the middle value.
	median = float64(x[(xLen-1)/2])
	return
}

// MedianFloat calulates the median of a slice of float64.
func MedianFloat(x []float64) (median float64) {
	// Sort the slice.
	sort.Float64s(x)

	xLen := len(x)
	if xLen%2 == 0 {
		// Length is even so calculate the average between the two middle values.
		median = (x[xLen/2] + x[(xLen/2)-1]) / float64(2)
		return
	}
	// Length is odd so return the middle value.
	median = (x[(xLen-1)/2])
	return
}
