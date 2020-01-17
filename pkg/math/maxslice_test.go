package math_test

import (
	. "github.com/knightjdr/prohits-viz-analysis/pkg/math"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MaxSlice", func() {
	It("should return the index of the maximum value in a slice of floats", func() {
		slice := []float64{0.5, 0.3, 4.1, 0.31}
		Expect(MaxIndexSliceFloat(slice)).To(Equal(2))
	})

	It("should return the maximum value in a slice of floats", func() {
		slice := []float64{0.5, 0.3, 0.31, 4.1}
		Expect(MaxSliceFloat(slice)).To(Equal(4.1))
	})

	It("should return the index of the maximum value in a slice of int", func() {
		slice := []int{5, 3, 7, 3}
		Expect(MaxSliceInt(slice)).To(Equal(7))
	})
})
