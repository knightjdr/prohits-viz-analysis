package scatter

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Legend", func() {
	It("should define legend for Binary plot", func() {
		format := 1
		plotSettings := &jsonSettings{
			Plot: "Binary",
		}

		expected := []map[string]string{
			{
				"color": "#ff0000",
				"text":  "Infinite fold change",
			},
		}
		Expect(defineLegend(plotSettings, format)).To(Equal(expected))
	})

	It("should define legend for Binary-biDist plot", func() {
		format := 2
		plotSettings := &jsonSettings{
			Plot:            "Binary-biDist",
			PrimaryFilter:   0.01,
			Score:           "FDR",
			ScoreType:       0,
			SecondaryFilter: 0.05,
		}

		expected := []map[string]string{
			{
				"color": "#0066cc",
				"text":  "FDR ≤ 0.01",
			},
			{
				"color": "#99ccff",
				"text":  "0.01 < FDR ≤ 0.05",
			},
		}
		Expect(defineLegend(plotSettings, format)).To(Equal(expected))
	})

	It("should define legend for Specificity plot in format 1", func() {
		format := 1
		plotSettings := &jsonSettings{
			Plot: "Specificity",
		}

		expected := []map[string]string{
			{
				"color": "#ff0000",
				"text":  "Infinite specificity",
			},
		}
		Expect(defineLegend(plotSettings, format)).To(Equal(expected))
	})

	It("should define legend for Specificity plot in format 1", func() {
		format := 1
		plotSettings := &jsonSettings{
			Plot: "Specificity",
		}

		expected := []map[string]string{
			{
				"color": "#ff0000",
				"text":  "Infinite specificity",
			},
		}
		Expect(defineLegend(plotSettings, format)).To(Equal(expected))
	})

	It("should define legend for Specificity plot in format 2", func() {
		format := 2
		plotSettings := &jsonSettings{
			Plot: "Specificity",
		}

		expected := []map[string]string{
			{
				"color": "#0066cc",
				"text":  "Infinite specificity",
			},
		}
		Expect(defineLegend(plotSettings, format)).To(Equal(expected))
	})
})

var _ = Describe("Legend score symbols", func() {
	It("should create score symbols for lte", func() {
		scoreType := 0
		expected := []string{"≤", "<", "≤"}
		Expect(createLegendScoreSymbols(scoreType)).To(Equal(expected))
	})

	It("should create score symbols for gte", func() {
		scoreType := 1
		expected := []string{"≥", ">", "≥"}
		Expect(createLegendScoreSymbols(scoreType)).To(Equal(expected))
	})
})
