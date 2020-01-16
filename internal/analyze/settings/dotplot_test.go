package settings

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log dotplot settings", func() {
	It("should log settings", func() {
		settings := types.Settings{
			AbundanceCap:       10.00,
			BiclusteringApprox: true,
			Condition:          "Bait",
			Clustering:         "biclustering",
			MinAbundance:       5.00,
			MinConditions:      0,
			PrimaryFilter:      0.01,
			Readout:            "Prey",
			Score:              "scoreColumn",
			ScoreType:          "lte",
			SecondaryFilter:    0.05,
		}

		expected := "Filtering\n" +
			"- minimum Bait requirement: 0\n" +
			"- parsimonius Prey inclusion: false\n\n" +
			"Abundance\n" +
			"- minimum abundance required: 5\n" +
			"- abundances were capped at 10 for visualization\n\n" +
			"Scoring\n" +
			"- smaller scores are better\n" +
			"- primary filter: 0.01\n" +
			"- secondary filter: 0.05\n\n" +
			"Clustering\n" +
			"- approximate biclustering was performed\n\n"

		var messages strings.Builder
		logDotplotSettings(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log dotplot abundance settings", func() {
	It("should log minimum abundance", func() {
		settings := types.Settings{
			MinAbundance: 5.00,
		}

		expected := "Abundance\n" +
			"- minimum abundance required: 5\n\n"

		var messages strings.Builder
		logDotplotAbundance(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log abundance cap", func() {
		settings := types.Settings{
			AbundanceCap: 10.00,
			MinAbundance: 5.00,
		}

		expected := "Abundance\n" +
			"- minimum abundance required: 5\n" +
			"- abundances were capped at 10 for visualization\n\n"

		var messages strings.Builder
		logDotplotAbundance(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log dotplot score settings", func() {
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
		logDotplotScoring(&messages, settings)
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
		logDotplotScoring(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log dotplot clustering settings", func() {
	It("should log biclustering settings", func() {
		settings := types.Settings{
			BiclusteringApprox: true,
			Clustering:         "biclustering",
		}

		expected := "Clustering\n" +
			"- approximate biclustering was performed\n\n"

		var messages strings.Builder
		logDotplotClustering(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log hierarchical clustering settings", func() {
		settings := types.Settings{
			Clustering:         "hierarchical",
			ClusteringMethod:   "complete",
			ClusteringOptimize: true,
			Distance:           "canberra",
		}

		expected := "Clustering\n" +
			"- hierarchical clustering was performed\n" +
			"- distance metric: canberra\n" +
			"- linkage method: complete\n" +
			"- leaf clusters were optimized\n\n"

		var messages strings.Builder
		logDotplotClustering(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log no clustering settings", func() {
		settings := types.Settings{
			Clustering:          "none",
			ConditionClustering: "none",
			ReadoutClustering:   "none",
		}

		expected := "Clustering\n" +
			"- no clustering was performed\n\n"

		var messages strings.Builder
		logDotplotClustering(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})
