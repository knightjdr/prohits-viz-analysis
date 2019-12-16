package stats_test

import (
	"github.com/knightjdr/cmgo/pkg/stats"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mean float", func() {
	It("Should return mean of a slice of floats", func() {
		slice := []float64{7, 6, 3, 9, 11}
		Expect(stats.MeanFloat(slice)).To(Equal(7.2))
	})

	It("Should return zero for empty slice", func() {
		slice := []float64{}
		Expect(stats.MeanFloat(slice)).To(Equal(float64(0)))
	})
})

var _ = Describe("Mean int", func() {
	It("Should return mean of a slice of ints", func() {
		slice := []int{7, 6, 3, 9, 11}
		Expect(stats.MeanInt(slice)).To(Equal(7.2))
	})

	It("Should return zero for empty slice", func() {
		slice := []int{}
		Expect(stats.MeanInt(slice)).To(Equal(float64(0)))
	})
})
