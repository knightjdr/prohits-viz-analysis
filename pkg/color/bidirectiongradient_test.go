package color

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Bi-directional gradient", func() {
	It("should create blueYellow gradient", func() {
		settings := &Gradient{
			ColorSpace: "blueYellow",
			NumColors:  11,
		}

		expected := []Space{
			{Hex: "#0040ff", RGB: []int{0, 64, 255}},
			{Hex: "#3366ff", RGB: []int{51, 102, 255}},
			{Hex: "#668cff", RGB: []int{102, 140, 255}},
			{Hex: "#99b3ff", RGB: []int{153, 179, 255}},
			{Hex: "#ccd9ff", RGB: []int{204, 217, 255}},
			{Hex: "#ffffff", RGB: []int{255, 255, 255}},
			{Hex: "#ffffcc", RGB: []int{255, 255, 204}},
			{Hex: "#ffff99", RGB: []int{255, 255, 153}},
			{Hex: "#ffff66", RGB: []int{255, 255, 102}},
			{Hex: "#ffff33", RGB: []int{255, 255, 51}},
			{Hex: "#ffff00", RGB: []int{255, 255, 0}},
		}
		Expect(createBiDirectionGradient(settings)).To(Equal(expected))
	})

	It("should create blueRed gradient", func() {
		settings := &Gradient{
			ColorSpace: "blueRed",
			NumColors:  11,
		}

		expected := []Space{
			{Hex: "#0040ff", RGB: []int{0, 64, 255}},
			{Hex: "#3366ff", RGB: []int{51, 102, 255}},
			{Hex: "#668cff", RGB: []int{102, 140, 255}},
			{Hex: "#99b3ff", RGB: []int{153, 179, 255}},
			{Hex: "#ccd9ff", RGB: []int{204, 217, 255}},
			{Hex: "#ffffff", RGB: []int{255, 255, 255}},
			{Hex: "#ffcccc", RGB: []int{255, 204, 204}},
			{Hex: "#ff9999", RGB: []int{255, 153, 153}},
			{Hex: "#ff6666", RGB: []int{255, 102, 102}},
			{Hex: "#ff3333", RGB: []int{255, 51, 51}},
			{Hex: "#ff0000", RGB: []int{255, 0, 0}},
		}
		Expect(createBiDirectionGradient(settings)).To(Equal(expected))
	})
})
