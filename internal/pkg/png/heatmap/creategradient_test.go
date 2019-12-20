package heatmap

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/color"
)

var _ = Describe("Create gradient", func() {
	It("should create a color gradient", func() {
		heatmap := &Heatmap{
			ColorSpace: "blue",
			Invert:     true,
			NumColors:  11,
		}

		expected := []color.Space{
			color.Space{Hex: "#000000", RGB: []int{0, 0, 0}},
			color.Space{Hex: "#000d33", RGB: []int{0, 13, 51}},
			color.Space{Hex: "#001966", RGB: []int{0, 25, 102}},
			color.Space{Hex: "#002699", RGB: []int{0, 38, 153}},
			color.Space{Hex: "#0033cc", RGB: []int{0, 51, 204}},
			color.Space{Hex: "#0040ff", RGB: []int{0, 64, 255}},
			color.Space{Hex: "#3366ff", RGB: []int{51, 102, 255}},
			color.Space{Hex: "#668cff", RGB: []int{102, 140, 255}},
			color.Space{Hex: "#99b3ff", RGB: []int{153, 179, 255}},
			color.Space{Hex: "#ccd9ff", RGB: []int{204, 217, 255}},
			color.Space{Hex: "#ffffff", RGB: []int{255, 255, 255}},
		}
		Expect(createGradient(heatmap)).To(Equal(expected))
	})
})
