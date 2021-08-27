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
			AbundanceCap:  50,
			AbundanceType: "positive",
			EdgeColor:     "blue",
			FillColor:     "blue",
			FillMax:       50,
			FillMin:       0,
			MinAbundance:  0,
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
			AbundanceCap:  24,
			AbundanceType: "bidirectional",
			FillColor:     "blueRed",
			FillMax:       24,
			FillMin:       -2,
			MinAbundance:  0,
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

var _ = Describe("Set abundance type", func() {
	It("should abundance type and fill for image with only non-negative values", func() {
		min := float64(1)
		settings := types.Settings{}

		setAbundanceType(min, &settings)
		Expect(settings.AbundanceType).To(Equal("positive"), "should set abundance type")
		Expect(settings.FillColor).To(Equal("blue"), "should set fill color")
	})

	It("should abundance type and fill for image with negative values", func() {
		min := float64(-1.5)
		settings := types.Settings{}

		setAbundanceType(min, &settings)
		Expect(settings.AbundanceType).To(Equal("bidirectional"), "should set abundance type")
		Expect(settings.FillColor).To(Equal("blueRed"), "should set fill color")
	})
})

var _ = Describe("Set fill parameters", func() {
	It("should set parameters for images with non-negative numbers and max > 1", func() {
		max := float64(10)
		min := float64(0)
		settings := types.Settings{}

		setFillParameters(min, max, &settings)
		Expect(settings.AbundanceCap).To(Equal(float64(50)), "should set abundance cap")
		Expect(settings.FillMax).To(Equal(float64(50)), "should set fill max")
		Expect(settings.FillMin).To(Equal(float64(0)), "should set fill min")
		Expect(settings.MinAbundance).To(Equal(float64(0)), "should set min abundance")
	})

	It("should set parameters for images with non-negative numbers and max <= 1", func() {
		max := float64(1)
		min := float64(0)
		settings := types.Settings{}

		setFillParameters(min, max, &settings)
		Expect(settings.AbundanceCap).To(Equal(float64(1)), "should set abundance cap")
		Expect(settings.FillMax).To(Equal(float64(1)), "should set fill max")
		Expect(settings.FillMin).To(Equal(float64(0)), "should set fill min")
		Expect(settings.MinAbundance).To(Equal(float64(0)), "should set min abundance")
	})

	It("should set parameters for images containing negative numbers", func() {
		max := float64(0.98)
		min := float64(-0.99)
		settings := types.Settings{}

		setFillParameters(min, max, &settings)
		Expect(settings.AbundanceCap).To(Equal(float64(1)), "should set abundance cap")
		Expect(settings.FillMax).To(Equal(float64(1)), "should set fill max")
		Expect(settings.FillMin).To(Equal(float64(-1)), "should set fill min")
		Expect(settings.MinAbundance).To(Equal(float64(0)), "should set min abundance")
	})
})
