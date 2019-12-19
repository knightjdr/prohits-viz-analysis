package downsample

// Should a matrix be downsampled. Any matrix with a dimension
// > 1000 should be.
func Should(matrix [][]float64) bool {
	if len(matrix) > 1000 || len(matrix[0]) > 1000 {
		return true
	}
	return false
}
