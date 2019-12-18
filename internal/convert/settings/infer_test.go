package settings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

var _ = Describe("Infer", func() {
	It("should infer settings for dotplot", func() {
		csv := []map[string]string{
			{"abundance": "5"},
			{"abundance": "0.05"},
			{"abundance": "23.8"},
			{"abundance": "7.5"},
		}
		settings := types.Settings{
			Type: "dotplot",
		}

		expected := types.Settings{
			AbundanceCap: 50,
			EdgeColor:    "blue",
			FillColor:    "blue",
			MinAbundance: 0,
			Type:         "dotplot",
		}
		addInferredSettings(csv, &settings)
		Expect(settings).To(Equal(expected))
	})

	It("should infer settings for heatmap", func() {
		csv := []map[string]string{
			{"abundance": "5"},
			{"abundance": "-1.5"},
			{"abundance": "23.8"},
			{"abundance": "7.5"},
		}
		settings := types.Settings{
			Type: "heatmap",
		}

		expected := types.Settings{
			AbundanceCap: 24,
			FillColor:    "blueRed",
			MinAbundance: -2,
			Type:         "heatmap",
		}
		addInferredSettings(csv, &settings)
		Expect(settings).To(Equal(expected))
	})
})
