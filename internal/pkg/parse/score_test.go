package parse_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/parse"
)

var _ = Describe("Parse score", func() {
	It("should parse a float", func() {
		score := "0.1"
		expected := 0.1
		Expect(Score(score)).To(Equal(expected))
	})

	It("should return 0 when score cannot be parsed", func() {
		score := "a"
		expected := float64(0)
		Expect(Score(score)).To(Equal(expected))
	})
})
