package math

// DefineValues returns "positive" for a matrix of non-negative values.
// "negative" for a matrix of non-positive values, and "bidirectional"
// for a matrix of both positive and negative values
func DefineValues(matrix [][]float64) string {
	min, max := MinMax(matrix)

	if max > 0 && min < 0 {
		return "bidirectional"
	}
	if max <= 0 && min < 0 {
		return "negative"
	}
	return "positive"
}
