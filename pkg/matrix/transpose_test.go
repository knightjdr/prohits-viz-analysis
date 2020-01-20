package matrix_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/matrix"
)

var _ = Describe("Transpose", func() {
	It("should transpose matrix", func() {
		matrix := [][]float64{
			{5, 2, 14.3, 2.1},
			{23, 17.8, 0, 0.4},
			{10, 0, 7, 15.9},
		}

		expected := [][]float64{
			{5, 23, 10},
			{2, 17.8, 0},
			{14.3, 0, 7},
			{2.1, 0.4, 15.9},
		}
		Expect(Transpose(matrix)).To(Equal(expected))
	})
})
