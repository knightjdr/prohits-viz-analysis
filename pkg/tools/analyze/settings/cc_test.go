package settings

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log condition-condition settings", func() {
	It("should log settings", func() {
		settings := types.Settings{
			ConditionX:      "condition1",
			ConditionY:      "condition2",
			MinAbundance:    5.00,
			PrimaryFilter:   0.01,
			ScoreType:       "lte",
			SecondaryFilter: 0.05,
		}

		expected := "Conditions\n" +
			"- x-axis: condition1\n" +
			"- y-axis: condition2\n\n" +
			"Abundance\n" +
			"- minimum abundance required: 5\n\n" +
			"Scoring\n" +
			"- smaller scores are better\n" +
			"- primary filter: 0.01\n" +
			"- secondary filter: 0.05\n\n"

		var messages strings.Builder
		logCCSettings(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log condition-condition abundance settings", func() {
	It("should log minimum abundance", func() {
		settings := types.Settings{
			MinAbundance: 5.00,
		}

		expected := "Abundance\n" +
			"- minimum abundance required: 5\n\n"

		var messages strings.Builder
		logCCAbundance(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log condition-condition conditions", func() {
	It("should log minimum abundance", func() {
		settings := types.Settings{
			ConditionX: "condition1",
			ConditionY: "condition2",
		}

		expected := "Conditions\n" +
			"- x-axis: condition1\n" +
			"- y-axis: condition2\n\n"

		var messages strings.Builder
		logCCConditions(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log condition-condition score settings", func() {
	It("should log score settings for lte", func() {
		settings := types.Settings{
			PrimaryFilter:   0.01,
			ScoreType:       "lte",
			SecondaryFilter: 0.05,
		}

		expected := "Scoring\n" +
			"- smaller scores are better\n" +
			"- primary filter: 0.01\n" +
			"- secondary filter: 0.05\n\n"

		var messages strings.Builder
		logCCScoring(&messages, settings)
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
		logCCScoring(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})
