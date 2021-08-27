package settings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
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
			AbundanceCap:  50,
			AbundanceType: "positive",
			EdgeColor:     "blue",
			FillColor:     "blue",
			FillMax:       50,
			FillMin:       0,
			MinAbundance:  0,
			Type:          "dotplot",
		}
		addInferredSettings(csv, &settings)
		Expect(settings).To(Equal(expected))
	})

	It("should infer settings for heatmap", func() {
		csv := []map[string]string{
			{"abundance": "5"},
			{"abundance": "0.05"},
			{"abundance": "23.8"},
			{"abundance": "7.5"},
		}
		settings := types.Settings{
			Type: "heatmap",
		}

		expected := types.Settings{
			AbundanceCap:  50,
			AbundanceType: "positive",
			FillColor:     "blue",
			FillMax:       50,
			FillMin:       0,
			MinAbundance:  0,
			Type:          "heatmap",
		}
		addInferredSettings(csv, &settings)
		Expect(settings).To(Equal(expected))
	})

	It("should infer settings for correlation heatmap", func() {
		csv := []map[string]string{
			{"abundance": "1"},
			{"abundance": "-0.5"},
			{"abundance": "0.2"},
			{"abundance": "0.3"},
		}
		settings := types.Settings{
			Type: "heatmap",
		}

		expected := types.Settings{
			AbundanceCap:  1,
			AbundanceType: "bidirectional",
			FillColor:     "blueRed",
			FillMax:       1,
			FillMin:       -1,
			MinAbundance:  0,
			Type:          "heatmap",
		}
		addInferredSettings(csv, &settings)
		Expect(settings).To(Equal(expected))
	})

	It("should infer settings for distance heatmap", func() {
		csv := []map[string]string{
			{"abundance": "1"},
			{"abundance": "0.5"},
			{"abundance": "0"},
			{"abundance": "0.3"},
		}
		settings := types.Settings{
			Type: "heatmap",
		}

		expected := types.Settings{
			AbundanceCap:  1,
			AbundanceType: "positive",
			FillColor:     "blue",
			FillMax:       1,
			FillMin:       0,
			MinAbundance:  0,
			Type:          "heatmap",
		}
		addInferredSettings(csv, &settings)
		Expect(settings).To(Equal(expected))
	})
})
