package color_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/color"
)

var _ = Describe("Create gradient", func() {
	It("should create a bi-directional gradient", func() {
		gradient := InitializeGradient()
		gradient.ColorSpace = "blueRed"
		gradient.Invert = false
		gradient.NumColors = 11

		expected := []Space{
			Space{Hex: "#0040ff", RGB: []int{0, 64, 255}},
			Space{Hex: "#3366ff", RGB: []int{51, 102, 255}},
			Space{Hex: "#668cff", RGB: []int{102, 140, 255}},
			Space{Hex: "#99b3ff", RGB: []int{153, 179, 255}},
			Space{Hex: "#ccd9ff", RGB: []int{204, 217, 255}},
			Space{Hex: "#ffffff", RGB: []int{255, 255, 255}},
			Space{Hex: "#ffcccc", RGB: []int{255, 204, 204}},
			Space{Hex: "#ff9999", RGB: []int{255, 153, 153}},
			Space{Hex: "#ff6666", RGB: []int{255, 102, 102}},
			Space{Hex: "#ff3333", RGB: []int{255, 51, 51}},
			Space{Hex: "#ff0000", RGB: []int{255, 0, 0}},
		}
		Expect(gradient.CreateColorGradient()).To(Equal(expected))
	})

	It("should create a mono-directional gradient", func() {
		gradient := InitializeGradient()
		gradient.ColorSpace = "blue"
		gradient.Invert = false
		gradient.NumColors = 11

		expected := []Space{
			Space{Hex: "#ffffff", RGB: []int{255, 255, 255}},
			Space{Hex: "#ccd9ff", RGB: []int{204, 217, 255}},
			Space{Hex: "#99b3ff", RGB: []int{153, 179, 255}},
			Space{Hex: "#668cff", RGB: []int{102, 140, 255}},
			Space{Hex: "#3366ff", RGB: []int{51, 102, 255}},
			Space{Hex: "#0040ff", RGB: []int{0, 64, 255}},
			Space{Hex: "#0033cc", RGB: []int{0, 51, 204}},
			Space{Hex: "#002699", RGB: []int{0, 38, 153}},
			Space{Hex: "#001966", RGB: []int{0, 25, 102}},
			Space{Hex: "#000d33", RGB: []int{0, 13, 51}},
			Space{Hex: "#000000", RGB: []int{0, 0, 0}},
		}
		Expect(gradient.CreateColorGradient()).To(Equal(expected))
	})

	It("should create an inverted gradient", func() {
		gradient := InitializeGradient()
		gradient.ColorSpace = "blue"
		gradient.Invert = true
		gradient.NumColors = 11

		expected := []Space{
			Space{Hex: "#000000", RGB: []int{0, 0, 0}},
			Space{Hex: "#000d33", RGB: []int{0, 13, 51}},
			Space{Hex: "#001966", RGB: []int{0, 25, 102}},
			Space{Hex: "#002699", RGB: []int{0, 38, 153}},
			Space{Hex: "#0033cc", RGB: []int{0, 51, 204}},
			Space{Hex: "#0040ff", RGB: []int{0, 64, 255}},
			Space{Hex: "#3366ff", RGB: []int{51, 102, 255}},
			Space{Hex: "#668cff", RGB: []int{102, 140, 255}},
			Space{Hex: "#99b3ff", RGB: []int{153, 179, 255}},
			Space{Hex: "#ccd9ff", RGB: []int{204, 217, 255}},
			Space{Hex: "#ffffff", RGB: []int{255, 255, 255}},
		}
		Expect(gradient.CreateColorGradient()).To(Equal(expected))
	})
})
