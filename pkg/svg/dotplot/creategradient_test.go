package dotplot

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/color"
)

var _ = Describe("Create gradients", func() {
	It("should create color gradient", func() {
		dotplot := &Dotplot{
			EdgeColor: "blue",
			FillColor: "blue",
			Invert:    true,
			NumColors: 11,
		}

		expectedEdge := []color.Space{
			{Hex: "#ffffff", RGB: []int{255, 255, 255}},
			{Hex: "#ccd9ff", RGB: []int{204, 217, 255}},
			{Hex: "#99b3ff", RGB: []int{153, 179, 255}},
			{Hex: "#668cff", RGB: []int{102, 140, 255}},
			{Hex: "#3366ff", RGB: []int{51, 102, 255}},
			{Hex: "#0040ff", RGB: []int{0, 64, 255}},
			{Hex: "#0033cc", RGB: []int{0, 51, 204}},
			{Hex: "#002699", RGB: []int{0, 38, 153}},
			{Hex: "#001966", RGB: []int{0, 25, 102}},
			{Hex: "#000d33", RGB: []int{0, 13, 51}},
			{Hex: "#000000", RGB: []int{0, 0, 0}},
		}
		expectedFill := []color.Space{
			{Hex: "#000000", RGB: []int{0, 0, 0}},
			{Hex: "#000d33", RGB: []int{0, 13, 51}},
			{Hex: "#001966", RGB: []int{0, 25, 102}},
			{Hex: "#002699", RGB: []int{0, 38, 153}},
			{Hex: "#0033cc", RGB: []int{0, 51, 204}},
			{Hex: "#0040ff", RGB: []int{0, 64, 255}},
			{Hex: "#3366ff", RGB: []int{51, 102, 255}},
			{Hex: "#668cff", RGB: []int{102, 140, 255}},
			{Hex: "#99b3ff", RGB: []int{153, 179, 255}},
			{Hex: "#ccd9ff", RGB: []int{204, 217, 255}},
			{Hex: "#ffffff", RGB: []int{255, 255, 255}},
		}

		actualFill, actualEdge := createGradients(dotplot)
		Expect(actualFill).To(Equal(expectedFill))
		Expect(actualEdge).To(Equal(expectedEdge))
	})
})
