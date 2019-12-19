package downsample_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"
)

var _ = Describe("Matrix", func() {
	It("should downsample matrix", func() {
		matrix := [][]float64{
			{1, 1, 2, 2, 2},
			{2, 2, 2, 3, 3},
			{1, 1, 4, 1, 1},
			{4, 2, 2, 3, 1},
			{5, 3, 4, 1, 3},
		}
		maxDimension := 2

		expected := [][]float64{
			{1.6, 2.24},
			{3.04, 2.08},
		}
		Expect(Matrix(matrix, maxDimension)).To(Equal(expected))
	})
})
