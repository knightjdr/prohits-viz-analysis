package slice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/slice"
)

var _ = Describe("Index of string", func() {
	It("should return index of string", func() {
		s := []string{"a", "b", "c", "d", "e"}
		Expect(IndexOfString("c", s)).To(Equal(2))
	})

	It("should return -1 when string not found", func() {
		s := []string{"a", "b", "c", "d", "e"}
		Expect(IndexOfString("f", s)).To(Equal(-1))
	})
})
