package helper

// NormalizeSlice normalizes the values in a slice to one. The greatest value
// in the slice will be one and the other values will be divided by this.
func NormalizeSlice(unnormalized []float64) (normalized []float64) {
	// Find max.
	max := float64(0)
	for _, value := range unnormalized {
		if value > max {
			max = value
		}
	}

	// Normalize input slice.
	normalized = make([]float64, len(unnormalized))
	for i, value := range unnormalized {
		normalized[i] = Round(value/max, 0.01)
	}
	return
}

// NormalizeMatrix every row in a matrix to one.
func NormalizeMatrix(matrix [][]float64) (normalized [][]float64) {
	for _, row := range matrix {
		normRow := NormalizeSlice(row)
		normalized = append(normalized, normRow)
	}
	return
}
