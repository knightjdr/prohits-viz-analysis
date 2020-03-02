package settings

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Infer dotplot settings", func() {
	It("should infer settings", func() {
		csv := []map[string]string{
			{"abundance": "5"},
			{"abundance": "0.05"},
			{"abundance": "23.8"},
			{"abundance": "7.5"},
		}
		settings := types.Settings{}

		expected := types.Settings{
			AbundanceCap: 50,
			EdgeColor:    "blue",
			FillColor:    "blue",
			MinAbundance: 0,
		}
		inferDotplotSettings(csv, &settings)
		Expect(settings).To(Equal(expected))
	})
})

var _ = Describe("Infer heatmap settings", func() {
	It("should infer settings", func() {
		csv := []map[string]string{
			{"abundance": "5"},
			{"abundance": "-1.5"},
			{"abundance": "23.8"},
			{"abundance": "7.5"},
		}
		settings := types.Settings{}

		expected := types.Settings{
			AbundanceCap: 24,
			FillColor:    "blueRed",
			MinAbundance: -2,
		}
		inferHeatmapSettings(csv, &settings)
		Expect(settings).To(Equal(expected))
	})
})

var _ = Describe("Find min and max", func() {
	It("should find min and max values", func() {
		csv := []map[string]string{
			{"abundance": "5"},
			{"abundance": "0.05"},
			{"abundance": "23.8"},
			{"abundance": "7.5"},
		}

		expectedMax := 23.8
		expectedMin := 0.05
		actualMin, actualMax := findMinMax(csv)
		Expect(actualMax).To(Equal(expectedMax), "should find maximum value")
		Expect(actualMin).To(Equal(expectedMin), "should find minimum value")
	})
})

var _ = Describe("Set file and min abundance", func() {
	It("should fill and min abundance for image with only non-negative values", func() {
		min := float64(1)
		settings := types.Settings{}

		setFillAndMinAbundance(min, &settings)
		Expect(settings.FillColor).To(Equal("blue"), "should set fill color")
		Expect(settings.MinAbundance).To(Equal(float64(0)), "should set min abundance")
	})

	It("should fill and min abundance for image with negative values", func() {
		min := float64(-1.5)
		settings := types.Settings{}

		setFillAndMinAbundance(min, &settings)
		Expect(settings.FillColor).To(Equal("blueRed"), "should set fill color")
		Expect(settings.MinAbundance).To(Equal(float64(-2)), "should set min abundance")
	})
})

var _ = Describe("Set abundance cap", func() {
	It("should set value to 50 for dotplot images with min of 0 or greater", func() {
		isDotplot := true
		max := float64(10)
		min := float64(0)
		settings := types.Settings{}

		expected := float64(50)
		setAbundanceCap(min, max, &settings, isDotplot)
		Expect(settings.AbundanceCap).To(Equal(expected))
	})

	It("should set value to max arg for dotplot images with min less than zero", func() {
		isDotplot := true
		max := float64(10)
		min := float64(-1)
		settings := types.Settings{}

		expected := float64(10)
		setAbundanceCap(min, max, &settings, isDotplot)
		Expect(settings.AbundanceCap).To(Equal(expected))
	})

	It("should set value to max arg for heatmap images", func() {
		isDotplot := false
		max := float64(10)
		min := float64(-1)
		settings := types.Settings{}

		expected := float64(10)
		setAbundanceCap(min, max, &settings, isDotplot)
		Expect(settings.AbundanceCap).To(Equal(expected))
	})
})
