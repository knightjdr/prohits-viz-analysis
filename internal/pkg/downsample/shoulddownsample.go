package downsample

// Should a matrix be downsampled. Any matrix with a dimension
// > maxDimension should be.
func Should(matrix [][]float64, maxDimension int) bool {
	if len(matrix) > maxDimension || len(matrix[0]) > maxDimension {
		return true
	}
	return false
}
