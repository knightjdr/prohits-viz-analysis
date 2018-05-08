package dotplot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormMatrix(t *testing.T) {
	matrix := [][]float64{
		{10, 5, 2},
		{5, 20, 25},
		{100, 67, 3},
	}

	// TEST2: normalize a slice.
	wantSlice := []float64{1, 0.67, 0.03}
	assert.Equal(t, wantSlice, NormalizeSlice(matrix[2]), "Slice not normalized correctly")

	// TEST2: normalize a matrix.
	wantMatrix := [][]float64{
		{1, 0.5, 0.2},
		{0.2, 0.8, 1},
		{1, 0.67, 0.03},
	}
	assert.Equal(t, wantMatrix, NormalizeMatrix(matrix), "Matrix not normalized by rows correctly")
}
