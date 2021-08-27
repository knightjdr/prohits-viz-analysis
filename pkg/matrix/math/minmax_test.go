package math_test

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/matrix/math"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matrix min and max", func() {
	It("should return min and max for matrix", func() {
		matrix := [][]float64{
			{6.5, 5, 34.7},
			{7, 10, 3},
			{24, 8.9, 0},
		}

		expectedMax := 34.7
		expectedMin := float64(0)
		actualMin, actualMax := math.MinMax(matrix)
		Expect(actualMax).To(Equal(expectedMax), "should return maximum value")
		Expect(actualMin).To(Equal(expectedMin), "should return minimum value")
	})
})
