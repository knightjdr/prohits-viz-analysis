package sort_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/sort"
)

var _ = Describe("Sort slice float64 by indices", func() {
	It("should sort slice", func() {
		s := []float64{1, 2, 3, 4}
		indices := []int{0, 3, 2, 1}

		expected := []float64{1, 4, 3, 2}
		Expect(ByIndicesFloat(s, indices)).To(Equal(expected))
	})
})

var _ = Describe("Sort slice int by indices", func() {
	It("should sort slice", func() {
		s := []int{1, 2, 3, 4}
		indices := []int{0, 3, 2, 1}

		expected := []int{1, 4, 3, 2}
		Expect(ByIndicesInt(s, indices)).To(Equal(expected))
	})
})
