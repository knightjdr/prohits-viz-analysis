package color

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HSL to Hex", func() {
	It("should convert from HSL to Hex", func() {
		tests := []HSL{
			{h: float64(225) / float64(360), s: 1, l: 0.4},
			{h: float64(115) / float64(360), s: 0, l: 0.67},
			{h: float64(183) / float64(360), s: 0.23, l: 0.67},
			{h: float64(324) / float64(360), s: 0.52, l: 0.77},
			{h: float64(28) / float64(360), s: 0.52, l: 0.19},
		}

		expected := []Space{
			{Hex: "#0033cc", RGB: []int{0, 51, 204}},
			{Hex: "#ababab", RGB: []int{171, 171, 171}},
			{Hex: "#97bcbe", RGB: []int{151, 188, 190}},
			{Hex: "#e3a6ca", RGB: []int{227, 166, 202}},
			{Hex: "#4a2f17", RGB: []int{74, 47, 23}},
		}
		for i, test := range tests {
			Expect(convertHSLtoSpace(test)).To(Equal(expected[i]))
		}
	})
})
