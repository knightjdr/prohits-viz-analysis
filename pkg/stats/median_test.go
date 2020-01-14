package stats_test

import (
	. "github.com/knightjdr/prohits-viz-analysis/pkg/stats"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Median", func() {
	Context("[]float64", func() {
		It("Should return median of an odd length slice", func() {
			slice := []float64{7, 6, 3, 9, 12}
			Expect(MedianFloat(slice)).To(Equal(float64(7)))
		})

		It("Should return median of an even length slice", func() {
			slice := []float64{7, 6, 9, 12}
			Expect(MedianFloat(slice)).To(Equal(float64(8)))
		})
	})

	Context("[]int", func() {
		It("Should return median of an odd length slice", func() {
			slice := []int{7, 6, 3, 9, 12}
			Expect(MedianInt(slice)).To(Equal(float64(7)))
		})

		It("Should return median of an even length slice", func() {
			slice := []int{7, 6, 9, 12}
			Expect(MedianInt(slice)).To(Equal(float64(8)))
		})
	})
})
