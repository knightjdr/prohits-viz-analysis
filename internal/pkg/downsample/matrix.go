// Package downsample will downsample a matrix using area averaging.
package downsample

// Matrix downsamples a matrix.
func Matrix(matrix [][]float64, maxDimension int) [][]float64 {
	scale := calculateScale(matrix, maxDimension)
	return downsample(matrix, scale)
}
