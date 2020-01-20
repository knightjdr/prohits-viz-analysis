package slice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/slice"
)

var _ = Describe("Reverse", func() {
	It("should reverse a slice of strings", func() {
		s := []string{"a", "b", "c", "d"}

		expected := []string{"d", "c", "b", "a"}
		Expect(ReverseString(s)).To(Equal(expected))
	})

	It("should reverse a slice of ints", func() {
		s := []int{1, 2, 3, 4}

		expected := []int{4, 3, 2, 1}
		Expect(ReverseInt(s)).To(Equal(expected))
	})
})
