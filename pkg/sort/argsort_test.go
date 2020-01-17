package sort_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/sort"
)

var _ = Describe("Argsort float", func() {
	It("should return indices of sorted slice", func() {
		s := []float64{1.3, 7, 2, 2, 5}

		expected := []int{0, 2, 3, 4, 1}
		Expect(ArgsortFloat(s)).To(Equal(expected))
	})
})

var _ = Describe("Argsort int", func() {
	It("should return indices of sorted slice", func() {
		s := []int{1, 7, 2, 2, 5}

		expected := []int{0, 2, 3, 4, 1}
		Expect(ArgsortInt(s)).To(Equal(expected))
	})
})
