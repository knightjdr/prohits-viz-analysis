package parse_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/parse"
)

var _ = Describe("Parse pipe separated string to average float", func() {
	It("should parse a single float", func() {
		str := "5.5"
		expected := 5.5
		Expect(PipeSeparatedStringToMean(str)).To(Equal(expected))
	})

	It("should parse and average pipe separated floats", func() {
		str := "5.5|.|4.5||."
		expected := 2.0
		Expect(PipeSeparatedStringToMean(str)).To(Equal(expected))
	})

	It("should return 0 for unparseable value", func() {
		str := "|.|||."
		expected := float64(0)
		Expect(PipeSeparatedStringToMean(str)).To(Equal(expected))
	})
})

var _ = Describe("Parse pipe separated string to array of float", func() {
	It("should parse a single float", func() {
		str := "5.5"
		expected := []float64{5.5}
		Expect(PipeSeparatedStringToArray(str)).To(Equal(expected))
	})

	It("should parse and average pipe separated floats", func() {
		str := "5.5|.|4.5||."
		expected := []float64{5.5, 0, 4.5, 0, 0}
		Expect(PipeSeparatedStringToArray(str)).To(Equal(expected))
	})

	It("should return 0 for unparseable value", func() {
		str := "|.|||."
		expected := []float64{0, 0, 0, 0, 0}
		Expect(PipeSeparatedStringToArray(str)).To(Equal(expected))
	})
})
