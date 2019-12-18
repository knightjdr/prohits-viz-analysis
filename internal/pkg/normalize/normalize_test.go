package normalize_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/normalize"
)

var _ = Describe("Normalize slice", func() {
	It("should normalize slice", func() {
		slice := []float64{100, 67, 3}

		expected := []float64{1, 0.67, 0.03}
		Expect(Slice(slice)).To(Equal(expected))
	})
})

var _ = Describe("Normalize matrix", func() {
	It("should normalize matrix", func() {
		matrix := [][]float64{
			{10, 5, 2},
			{5, 20, 25},
			{100, 67, 3},
		}

		expected := [][]float64{
			{1, 0.5, 0.2},
			{0.2, 0.8, 1},
			{1, 0.67, 0.03},
		}
		Expect(Matrix(matrix)).To(Equal(expected))
	})
})
