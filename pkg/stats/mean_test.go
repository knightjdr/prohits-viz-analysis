package stats_test

import (
	. "github.com/knightjdr/prohits-viz-analysis/pkg/stats"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mean string", func() {
	It("Should return mean of a slice of strings", func() {
		slice := []string{"7", "6", "3", "9", "11"}
		Expect(MeanString(slice)).To(Equal(7.2))
	})

	It("Should return zero for empty slice", func() {
		slice := []string{}
		Expect(MeanString(slice)).To(Equal(float64(0)))
	})

	It("Should return zero for slice containing a value that cannot be parsed", func() {
		slice := []string{"7", "6", "a", "9", "11"}
		Expect(MeanString(slice)).To(Equal(float64(0)))
	})
})
