package settings

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log Specificity settings", func() {
	It("should log settings", func() {
		settings := types.Settings{
			MinAbundance:      5.00,
			PrimaryFilter:     0.01,
			ScoreType:         "lte",
			SpecificityMetric: "fe",
		}

		expected := "Metric\n" +
			"- specificity metric: fe\n\n" +
			"Abundance\n" +
			"- minimum abundance required: 5\n\n" +
			"Scoring\n" +
			"- smaller scores are better\n" +
			"- primary filter: 0.01\n\n"

		var messages strings.Builder
		logSpecificitySettings(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log Specificity abundance settings", func() {
	It("should log minimum abundance", func() {
		settings := types.Settings{
			MinAbundance: 5.00,
		}

		expected := "Abundance\n" +
			"- minimum abundance required: 5\n\n"

		var messages strings.Builder
		logSpecificityAbundance(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log Specificity metric setting", func() {
	It("should log minimum abundance", func() {
		settings := types.Settings{
			SpecificityMetric: "fe",
		}

		expected := "Metric\n" +
			"- specificity metric: fe\n\n"

		var messages strings.Builder
		logSpecificityMetric(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log Specificity score settings", func() {
	It("should log score settings for lte", func() {
		settings := types.Settings{
			PrimaryFilter: 0.01,
			ScoreType:     "lte",
		}

		expected := "Scoring\n" +
			"- smaller scores are better\n" +
			"- primary filter: 0.01\n\n"

		var messages strings.Builder
		logSpecificityScoring(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log score settings for gte", func() {
		settings := types.Settings{
			PrimaryFilter: 0.01,
			ScoreType:     "gte",
		}

		expected := "Scoring\n" +
			"- larger scores are better\n" +
			"- primary filter: 0.01\n\n"

		var messages strings.Builder
		logSpecificityScoring(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})
