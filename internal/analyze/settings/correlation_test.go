package settings

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log correlation settings", func() {
	It("should log settings", func() {
		settings := types.Settings{
			AlwaysIncludePreysPassingFilter: true,
			BaitAbundanceFilter:             5,
			BaitScoreFilter:                 0.01,
			Clustering:                      "hierarchical",
			ClusteringMethod:                "complete",
			ClusteringOptimize:              true,
			Correlation:                     "pearson",
			CytoscapeCutoff:                 0.7,
			Distance:                        "canberra",
			IgnoreSourceGenes:               true,
			MinBait:                         2,
			MockCountsForBait:               true,
			PreyAbundanceFilter:             10,
			PreyScoreFilter:                 0.05,
			ScoreType:                       "lte",
			UseReplicates:                   true,
		}

		expected := "Abundance\n" +
			"- minimum abundance for bait correlation: 5\n" +
			"- minimum abundance for prey correlation: 10\n\n" +
			"Scoring\n" +
			"- smaller scores are better\n" +
			"- score filter for bait correlation: 0.01\n" +
			"- score filter for prey correlation: 0.05\n\n" +
			"Correlation\n" +
			"- correlation method: pearson\n" +
			"- treat replicates as separate data points: true\n" +
			"- minimum bait requirement: 2\n" +
			"- always include preys passing filter criteria: true\n" +
			"- mock spectral counts for bait genes: true\n" +
			"- ignore source genes in pairwise correlations: true\n" +
			"- cytoscape cutoff: 0.7\n\n" +
			"Clustering\n" +
			"- hierarchical clustering was performed\n" +
			"- distance metric: canberra\n" +
			"- linkage method: complete\n" +
			"- leaf clusters were optimized\n\n"

		var messages strings.Builder
		logCorrelationSettings(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log correlation abundance settings", func() {
	It("should log settings", func() {
		settings := types.Settings{
			BaitAbundanceFilter: 5,
			PreyAbundanceFilter: 10,
		}

		expected := "Abundance\n" +
			"- minimum abundance for bait correlation: 5\n" +
			"- minimum abundance for prey correlation: 10\n\n"

		var messages strings.Builder
		logCorrelationAbundance(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log correlation score settings", func() {
	It("should log score settings for lte", func() {
		settings := types.Settings{
			BaitScoreFilter: 0.01,
			PreyScoreFilter: 0.05,
			ScoreType:       "lte",
		}

		expected := "Scoring\n" +
			"- smaller scores are better\n" +
			"- score filter for bait correlation: 0.01\n" +
			"- score filter for prey correlation: 0.05\n\n"

		var messages strings.Builder
		logCorrelationScoring(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log score settings for gte", func() {
		settings := types.Settings{
			BaitScoreFilter: 0.01,
			PreyScoreFilter: 0.05,
			ScoreType:       "gte",
		}

		expected := "Scoring\n" +
			"- larger scores are better\n" +
			"- score filter for bait correlation: 0.01\n" +
			"- score filter for prey correlation: 0.05\n\n"

		var messages strings.Builder
		logCorrelationScoring(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log correlation-specifc settings", func() {
	It("should log settings", func() {
		settings := types.Settings{
			Correlation:                     "pearson",
			UseReplicates:                   true,
			MinBait:                         2,
			AlwaysIncludePreysPassingFilter: true,
			MockCountsForBait:               true,
			IgnoreSourceGenes:               true,
			CytoscapeCutoff:                 0.7,
		}

		expected := "Correlation\n" +
			"- correlation method: pearson\n" +
			"- treat replicates as separate data points: true\n" +
			"- minimum bait requirement: 2\n" +
			"- always include preys passing filter criteria: true\n" +
			"- mock spectral counts for bait genes: true\n" +
			"- ignore source genes in pairwise correlations: true\n" +
			"- cytoscape cutoff: 0.7\n\n"

		var messages strings.Builder
		logCorrelationDetails(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log correlation clustering settings", func() {
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
		logCorrelationClustering(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})
