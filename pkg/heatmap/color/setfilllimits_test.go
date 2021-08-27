package color_test

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/color"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set fill limits", func() {
	It("should set limits for a positive values", func() {
		settings := types.Settings{
			AbundanceCap:  50,
			AbundanceType: "positive",
			MinAbundance:  0,
		}

		expected := types.Settings{
			AbundanceCap:  50,
			AbundanceType: "positive",
			FillMax:       50,
			FillMin:       0,
			MinAbundance:  0,
		}
		color.SetFillLimits(&settings)
		Expect(settings).To(Equal(expected))
	})

	It("should set limits for a negative values", func() {
		settings := types.Settings{
			AbundanceCap:  50,
			AbundanceType: "negative",
			MinAbundance:  0,
		}

		expected := types.Settings{
			AbundanceCap:  50,
			AbundanceType: "negative",
			FillMax:       0,
			FillMin:       -50,
			MinAbundance:  0,
		}
		color.SetFillLimits(&settings)
		Expect(settings).To(Equal(expected))
	})

	It("should set limits for a bidirectional values", func() {
		settings := types.Settings{
			AbundanceCap:  50,
			AbundanceType: "bidirectional",
			MinAbundance:  0,
		}

		expected := types.Settings{
			AbundanceCap:  50,
			AbundanceType: "bidirectional",
			FillMax:       50,
			FillMin:       -50,
			MinAbundance:  0,
		}
		color.SetFillLimits(&settings)
		Expect(settings).To(Equal(expected))
	})
})
