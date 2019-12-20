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
		Expect(ReverseStrings(s)).To(Equal(expected))
	})
})
