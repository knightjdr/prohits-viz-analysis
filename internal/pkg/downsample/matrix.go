// Package downsample will downsample a matrix using area averaging.
package downsample

// Matrix downsamples a matrix.
func Matrix(matrix [][]float64, maxDimension int) [][]float64 {
	treshold := defineThreshold(maxDimension)
	scale := calculateScale(matrix, treshold)
	return downsample(matrix, scale)
}
