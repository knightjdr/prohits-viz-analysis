package color_test

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/color"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Adjust fill color", func() {
	It("should return input color settings when adjustment not requested", func() {
		settings := types.Settings{
			AutomaticallySetFill: false,
			FillColor:            "red",
			InvertColor:          false,
		}

		expected := types.Settings{
			AutomaticallySetFill: false,
			FillColor:            "red",
			InvertColor:          false,
		}
		color.AdjustFillColor(&settings)
		Expect(settings).To(Equal(expected))
	})

	It("should adjust color settings for non-negative abundances", func() {
		settings := types.Settings{
			AbundanceType:        "positive",
			AutomaticallySetFill: true,
			FillColor:            "red",
			InvertColor:          false,
		}

		expected := types.Settings{
			AbundanceType:        "positive",
			AutomaticallySetFill: true,
			FillColor:            "blue",
			InvertColor:          false,
		}
		color.AdjustFillColor(&settings)
		Expect(settings).To(Equal(expected))
	})

	It("should adjust color settings for non-positive abundances", func() {
		settings := types.Settings{
			AbundanceType:        "negative",
			AutomaticallySetFill: true,
			FillColor:            "red",
			InvertColor:          false,
		}

		expected := types.Settings{
			AbundanceType:        "negative",
			AutomaticallySetFill: true,
			FillColor:            "blue",
			InvertColor:          true,
		}
		color.AdjustFillColor(&settings)
		Expect(settings).To(Equal(expected))
	})

	It("should adjust color settings for positive and negative abundances", func() {
		settings := types.Settings{
			AbundanceType:        "bidirectional",
			AutomaticallySetFill: true,
			FillColor:            "red",
			InvertColor:          false,
		}

		expected := types.Settings{
			AbundanceType:        "bidirectional",
			AutomaticallySetFill: true,
			FillColor:            "blueRed",
			InvertColor:          false,
		}
		color.AdjustFillColor(&settings)
		Expect(settings).To(Equal(expected))
	})
})
