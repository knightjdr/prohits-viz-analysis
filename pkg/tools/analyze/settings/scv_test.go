package settings

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log scv settings", func() {
	It("should log settings", func() {
		settings := types.Settings{
			AbundanceCap:          10.00,
			AbundanceFilterColumn: "AvgSpec",
			Condition:             "Bait",
			ConditionIDType:       "symbol",
			ConditionMapColumn:    "baitID",
			Known:                 "interaction",
			MinAbundance:          5.00,
			PrimaryFilter:         0.01,
			Readout:               "Prey",
			ReadoutIDType:         "refseqp",
			ReadoutMapColumn:      "preyID",
			Score:                 "scoreColumn",
			ScoreType:             "lte",
		}

		expected := "Abundance\n" +
			"- column for abundance filtering: AvgSpec\n" +
			"- minimum abundance required: 5\n" +
			"- abundances were capped at 10 for visualization\n\n" +
			"Scoring\n" +
			"- smaller scores are better\n" +
			"- primary filter: 0.01\n\n" +
			"ID mapping\n" +
			"- Bait ID type: symbol\n" +
			"- Bait IDs were mapped by column: baitID\n" +
			"- Prey ID type: refseqp\n" +
			"- Prey IDs were mapped by column: preyID\n\n" +
			"Known metric\n" +
			"- Prey knownness was evaluated by: interaction\n\n"

		var messages strings.Builder
		logSCVSettings(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log scv abundance settings", func() {
	It("should log minimum abundance", func() {
		settings := types.Settings{
			AbundanceFilterColumn: "AvgSpec",
			MinAbundance:          5.00,
		}

		expected := "Abundance\n" +
			"- column for abundance filtering: AvgSpec\n" +
			"- minimum abundance required: 5\n\n"

		var messages strings.Builder
		logSCVAbundance(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log abundance cap", func() {
		settings := types.Settings{
			AbundanceCap:          10.00,
			AbundanceFilterColumn: "AvgSpec",
			MinAbundance:          5.00,
		}

		expected := "Abundance\n" +
			"- column for abundance filtering: AvgSpec\n" +
			"- minimum abundance required: 5\n" +
			"- abundances were capped at 10 for visualization\n\n"

		var messages strings.Builder
		logSCVAbundance(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log scv known settings", func() {
	It("should log known settings", func() {
		settings := types.Settings{
			Known:   "interaction",
			Readout: "Prey",
		}

		expected := "Known metric\n" +
			"- Prey knownness was evaluated by: interaction\n\n"

		var messages strings.Builder
		logSCVKnown(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should not log known settings", func() {
		settings := types.Settings{
			Known:   "",
			Readout: "Prey",
		}

		expected := ""

		var messages strings.Builder
		logSCVKnown(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log scv mapping settings", func() {
	It("should log mapping settings when mapping by IDs by themselves", func() {
		settings := types.Settings{
			Condition:       "Bait",
			ConditionIDType: "symbol",
			Readout:         "Prey",
			ReadoutIDType:   "symbol",
		}

		expected := "ID mapping\n" +
			"- Bait ID type: symbol\n" +
			"- Prey ID type: symbol\n\n"

		var messages strings.Builder
		logSCVMapping(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log mapping settings when mapping by column", func() {
		settings := types.Settings{
			Condition:          "Bait",
			ConditionIDType:    "symbol",
			ConditionMapColumn: "baitID",
			Readout:            "Prey",
			ReadoutIDType:      "refseqp",
			ReadoutMapColumn:   "preyID",
		}

		expected := "ID mapping\n" +
			"- Bait ID type: symbol\n" +
			"- Bait IDs were mapped by column: baitID\n" +
			"- Prey ID type: refseqp\n" +
			"- Prey IDs were mapped by column: preyID\n\n"

		var messages strings.Builder
		logSCVMapping(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log mapping settings when mapping by file", func() {
		settings := types.Settings{
			Condition:        "Bait",
			ConditionIDType:  "symbol",
			ConditionMapFile: "bait-map.txt",
			Readout:          "Prey",
			ReadoutIDType:    "refseqp",
			ReadoutMapFile:   "prey-map.txt",
		}

		expected := "ID mapping\n" +
			"- Bait ID type: symbol\n" +
			"- Bait IDs were mapped by file\n" +
			"- Prey ID type: refseqp\n" +
			"- Prey IDs were mapped by file\n\n"

		var messages strings.Builder
		logSCVMapping(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log scv score settings", func() {
	It("should log score settings for lte", func() {
		settings := types.Settings{
			PrimaryFilter: 0.01,
			ScoreType:     "lte",
		}

		expected := "Scoring\n" +
			"- smaller scores are better\n" +
			"- primary filter: 0.01\n\n"

		var messages strings.Builder
		logSCVScoring(&messages, settings)
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
		logSCVScoring(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})
