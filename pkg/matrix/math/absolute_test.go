package math_test

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/matrix/math"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matrix math", func() {
	It("should convert the values of a matrix to their absolute form", func() {
		matrix := [][]float64{
			{5, 2, -14.3},
			{-23, 17.8, 0},
			{10, 0, -7},
		}

		expected := [][]float64{
			{5, 2, 14.3},
			{23, 17.8, 0},
			{10, 0, 7},
		}
		Expect(math.AbsoluteValueEntries(matrix)).To(Equal(expected))
	})
})
