package hierarchical

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Adjust analysis settings", func() {
	It("should adjust settings", func() {
		abundance := [][]float64{
			{-6.5, 5, 34.7},
			{-7, -10, 3},
			{24, 8.9, 0},
		}
		settings := types.Settings{
			AbundanceCap:         50,
			AutomaticallySetFill: true,
			MinAbundance:         0,
		}

		expected := types.Settings{
			AbundanceCap:         50,
			AbundanceType:        "bidirectional",
			AutomaticallySetFill: true,
			FillColor:            "blueRed",
			FillMax:              50,
			FillMin:              -50,
			InvertColor:          false,
			MinAbundance:         0,
		}
		Expect(AdjustSettings(settings, abundance)).To(Equal(expected))
	})
})
