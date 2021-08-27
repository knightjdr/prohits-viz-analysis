package correlation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Condition-readout images", func() {
	Describe("adjust settings for these images", func() {
		It("should change settings for matrix with non-negative values", func() {
			matrix := [][]float64{
				{6.5, 5, 34.7},
				{7, 10, 3},
				{24, 8.9, 0},
			}
			settings := types.Settings{
				ReadoutAbundanceFilter: 0,
				ReadoutScoreFilter:     0.01,
			}

			expected := types.Settings{
				AbundanceCap:           35,
				AbundanceType:          "positive",
				AutomaticallySetFill:   true,
				FillColor:              "blue",
				FillMax:                35,
				FillMin:                0,
				InvertColor:            false,
				MinAbundance:           0,
				ReadoutAbundanceFilter: 0,
				ReadoutScoreFilter:     0.01,
				PrimaryFilter:          0.01,
			}

			Expect(adjustConditionReadoutSettings(settings, matrix)).To(Equal(expected))
		})

		It("should change settings for matrix with non-positive values", func() {
			matrix := [][]float64{
				{-6.5, -5, -34.7},
				{-7, -10, -3},
				{-24, -8.9, 0},
			}
			settings := types.Settings{
				ReadoutAbundanceFilter: 0,
				ReadoutScoreFilter:     0.01,
			}

			expected := types.Settings{
				AbundanceCap:           35,
				AbundanceType:          "negative",
				AutomaticallySetFill:   true,
				FillColor:              "blue",
				FillMax:                0,
				FillMin:                -35,
				InvertColor:            true,
				MinAbundance:           0,
				ReadoutAbundanceFilter: 0,
				ReadoutScoreFilter:     0.01,
				PrimaryFilter:          0.01,
			}

			Expect(adjustConditionReadoutSettings(settings, matrix)).To(Equal(expected))
		})

		It("should change settings for matrix with positive and negative values", func() {
			matrix := [][]float64{
				{-6.5, 5, 34.7},
				{-7, -10, 3},
				{24, 8.9, 0},
			}
			settings := types.Settings{
				ReadoutAbundanceFilter: 0,
				ReadoutScoreFilter:     0.01,
			}

			expected := types.Settings{
				AbundanceCap:           35,
				AbundanceType:          "bidirectional",
				AutomaticallySetFill:   true,
				FillColor:              "blueRed",
				FillMax:                35,
				FillMin:                -35,
				InvertColor:            false,
				MinAbundance:           0,
				ReadoutAbundanceFilter: 0,
				ReadoutScoreFilter:     0.01,
				PrimaryFilter:          0.01,
			}

			Expect(adjustConditionReadoutSettings(settings, matrix)).To(Equal(expected))
		})
	})
})
