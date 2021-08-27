package math_test

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/matrix/math"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Define matrix values", func() {
	It("should return 'positive' for non-negative values", func() {
		matrix := [][]float64{
			{6.5, 5, 34.7},
			{7, 10, 3},
			{24, 8.9, 0},
		}

		expected := "positive"
		Expect(math.DefineValues(matrix)).To(Equal(expected))
	})

	It("should return 'negative' for non-positive values", func() {
		matrix := [][]float64{
			{-6.5, -5, -34.7},
			{-7, -10, -3},
			{-24, -8.9, 0},
		}

		expected := "negative"
		Expect(math.DefineValues(matrix)).To(Equal(expected))
	})

	It("should return 'bidirectional' for positive and negative values", func() {
		matrix := [][]float64{
			{-6.5, 5, 34.7},
			{-7, -10, 3},
			{24, 8.9, 0},
		}

		expected := "bidirectional"
		Expect(math.DefineValues(matrix)).To(Equal(expected))
	})
})
