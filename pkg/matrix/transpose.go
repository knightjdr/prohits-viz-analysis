// Package matrix contains methods for operating on matrices.
package matrix

// Transpose transposes a 2D matrix.
func Transpose(matrix [][]float64) [][]float64 {
	noCols := len(matrix[0])
	noRows := len(matrix)

	transposed := make([][]float64, noCols)
	for i := range transposed {
		transposed[i] = make([]float64, noRows)
	}

	for i, row := range matrix {
		for j, value := range row {
			transposed[j][i] = value
		}
	}

	return transposed
}
