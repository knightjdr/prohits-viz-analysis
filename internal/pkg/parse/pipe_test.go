package parse_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/parse"
)

var _ = Describe("Pipe", func() {
	It("should parse a single float", func() {
		str := "5.5"
		expected := 5.5
		Expect(PipeSeparatedFloat(str)).To(Equal(expected))
	})

	It("should parse and average pipe separated floats", func() {
		str := "5.5|.|4.5||."
		expected := 2.0
		Expect(PipeSeparatedFloat(str)).To(Equal(expected))
	})

	It("should return 0 for unparseable value", func() {
		str := "|.|||."
		expected := float64(0)
		Expect(PipeSeparatedFloat(str)).To(Equal(expected))
	})
})
