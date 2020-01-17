package math_test

import (
	. "github.com/knightjdr/prohits-viz-analysis/pkg/math"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MinSlice", func() {
	It("should return the minimum in a slice of floats", func() {
		slice := []float64{0.5, 0.3, 0.31, 4.1}
		Expect(MinSliceFloat(slice)).To(Equal(0.3))
	})

	Context("Slice of ints", func() {
		It("should return the minimum in a slice of ints", func() {
			slice := []int{50, 30, 31, 41}
			actual, _ := MinSliceInt(slice)
			Expect(actual).To(Equal(30))
		})

		It("should return an error when slice is of length 0", func() {
			slice := []int{}
			min, err := MinSliceInt(slice)
			Expect(err).Should(HaveOccurred())
			Expect(min).To(Equal(0), "should return 0 for min value")
		})
	})
})
