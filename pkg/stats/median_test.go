package stats_test

import (
	"github.com/knightjdr/cmgo/pkg/stats"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Median", func() {
	Context("[]float64", func() {
		It("Should return median of an odd length slice", func() {
			slice := []float64{7, 6, 3, 9, 12}
			Expect(stats.MedianFloat(slice)).To(Equal(float64(7)))
		})

		It("Should return median of an even length slice", func() {
			slice := []float64{7, 6, 9, 12}
			Expect(stats.MedianFloat(slice)).To(Equal(float64(8)))
		})
	})

	Context("[]int", func() {
		It("Should return median of an odd length slice", func() {
			slice := []int{7, 6, 3, 9, 12}
			Expect(stats.MedianInt(slice)).To(Equal(float64(7)))
		})

		It("Should return median of an even length slice", func() {
			slice := []int{7, 6, 9, 12}
			Expect(stats.MedianInt(slice)).To(Equal(float64(8)))
		})
	})
})
