package downsample

const defaultDownsampleThreshold = 1000

// Should a matrix be downsampled. Any matrix with a dimension
// > maxDimension should be.
func Should(matrix [][]float64, maxDimension int) bool {
	threshold := defineThreshold(maxDimension)

	if len(matrix) > threshold || len(matrix[0]) > threshold {
		return true
	}
	return false
}

func defineThreshold(userTreshold int) int {
	if userTreshold == 0 {
		return defaultDownsampleThreshold
	}
	return userTreshold
}
