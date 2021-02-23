package color

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create mono-directional gradient", func() {
	It("should create blue gradient", func() {
		settings := &Gradient{
			ColorSpace: "blue",
			NumColors:  11,
		}

		expected := []Space{
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
		Expect(createMonoDirectionGradiant(settings)).To(Equal(expected))
	})

	It("should create green gradient", func() {
		settings := &Gradient{
			ColorSpace: "green",
			NumColors:  11,
		}

		expected := []Space{
			{Hex: "#ffffff", RGB: []int{255, 255, 255}},
			{Hex: "#ccffcc", RGB: []int{204, 255, 204}},
			{Hex: "#99ff99", RGB: []int{153, 255, 153}},
			{Hex: "#66ff66", RGB: []int{102, 255, 102}},
			{Hex: "#33ff33", RGB: []int{51, 255, 51}},
			{Hex: "#00ff00", RGB: []int{0, 255, 0}},
			{Hex: "#00cc00", RGB: []int{0, 204, 0}},
			{Hex: "#009900", RGB: []int{0, 153, 0}},
			{Hex: "#006600", RGB: []int{0, 102, 0}},
			{Hex: "#003300", RGB: []int{0, 51, 0}},
			{Hex: "#000000", RGB: []int{0, 0, 0}},
		}
		Expect(createMonoDirectionGradiant(settings)).To(Equal(expected))
	})

	It("should create grey gradient", func() {
		settings := &Gradient{
			ColorSpace: "grey",
			NumColors:  11,
		}

		expected := []Space{
			{Hex: "#ffffff", RGB: []int{255, 255, 255}},
			{Hex: "#e6e6e6", RGB: []int{230, 230, 230}},
			{Hex: "#cccccc", RGB: []int{204, 204, 204}},
			{Hex: "#b3b3b3", RGB: []int{179, 179, 179}},
			{Hex: "#999999", RGB: []int{153, 153, 153}},
			{Hex: "#808080", RGB: []int{128, 128, 128}},
			{Hex: "#666666", RGB: []int{102, 102, 102}},
			{Hex: "#4d4d4d", RGB: []int{77, 77, 77}},
			{Hex: "#333333", RGB: []int{51, 51, 51}},
			{Hex: "#1a1a1a", RGB: []int{26, 26, 26}},
			{Hex: "#000000", RGB: []int{0, 0, 0}},
		}
		Expect(createMonoDirectionGradiant(settings)).To(Equal(expected))
	})

	It("should create red gradient", func() {
		settings := &Gradient{
			ColorSpace: "red",
			NumColors:  11,
		}

		expected := []Space{
			{Hex: "#ffffff", RGB: []int{255, 255, 255}},
			{Hex: "#ffcccc", RGB: []int{255, 204, 204}},
			{Hex: "#ff9999", RGB: []int{255, 153, 153}},
			{Hex: "#ff6666", RGB: []int{255, 102, 102}},
			{Hex: "#ff3333", RGB: []int{255, 51, 51}},
			{Hex: "#ff0000", RGB: []int{255, 0, 0}},
			{Hex: "#cc0000", RGB: []int{204, 0, 0}},
			{Hex: "#990000", RGB: []int{153, 0, 0}},
			{Hex: "#660000", RGB: []int{102, 0, 0}},
			{Hex: "#330000", RGB: []int{51, 0, 0}},
			{Hex: "#000000", RGB: []int{0, 0, 0}},
		}
		Expect(createMonoDirectionGradiant(settings)).To(Equal(expected))
	})

	It("should create yellow gradient", func() {
		settings := &Gradient{
			ColorSpace: "yellow",
			NumColors:  11,
		}

		expected := []Space{
			{Hex: "#ffffff", RGB: []int{255, 255, 255}},
			{Hex: "#ffffcc", RGB: []int{255, 255, 204}},
			{Hex: "#ffff99", RGB: []int{255, 255, 153}},
			{Hex: "#ffff66", RGB: []int{255, 255, 102}},
			{Hex: "#ffff33", RGB: []int{255, 255, 51}},
			{Hex: "#ffff00", RGB: []int{255, 255, 0}},
			{Hex: "#cccc00", RGB: []int{204, 204, 0}},
			{Hex: "#999900", RGB: []int{153, 153, 0}},
			{Hex: "#666600", RGB: []int{102, 102, 0}},
			{Hex: "#333300", RGB: []int{51, 51, 0}},
			{Hex: "#000000", RGB: []int{0, 0, 0}},
		}
		Expect(createMonoDirectionGradiant(settings)).To(Equal(expected))
	})
})
