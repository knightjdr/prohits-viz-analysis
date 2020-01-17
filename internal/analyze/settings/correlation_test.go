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
			Condition:                    "Bait",
			ConditionAbundanceFilter:     5,
			ConditionScoreFilter:         0.01,
			Clustering:                   "hierarchical",
			ClusteringMethod:             "complete",
			ClusteringOptimize:           true,
			Correlation:                  "pearson",
			CytoscapeCutoff:              0.7,
			Distance:                     "canberra",
			IgnoreSourceTargetPairs:      true,
			MinConditions:                2,
			MockConditionAbundance:       true,
			ParsimoniousReadoutFiltering: true,
			Readout:                      "Prey",
			ReadoutAbundanceFilter:       10,
			ReadoutScoreFilter:           0.05,
			ScoreType:                    "lte",
			UseReplicates:                true,
		}

		expected := "Filtering\n" +
			"- minimum Bait requirement: 2\n" +
			"- parsimonius Prey inclusion: true\n\n" +
			"Abundance\n" +
			"- minimum abundance for Bait correlation: 5\n" +
			"- minimum abundance for Prey correlation: 10\n\n" +
			"Scoring\n" +
			"- smaller scores are better\n" +
			"- score filter for Bait correlation: 0.01\n" +
			"- score filter for Prey correlation: 0.05\n\n" +
			"Correlation\n" +
			"- correlation method: pearson\n" +
			"- treat replicates as separate data points: true\n" +
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
			Condition:                "Bait",
			ConditionAbundanceFilter: 5,
			Readout:                  "Prey",
			ReadoutAbundanceFilter:   10,
		}

		expected := "Abundance\n" +
			"- minimum abundance for Bait correlation: 5\n" +
			"- minimum abundance for Prey correlation: 10\n\n"

		var messages strings.Builder
		logCorrelationAbundance(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log correlation score settings", func() {
	It("should log score settings for lte", func() {
		settings := types.Settings{
			Condition:            "Bait",
			ConditionScoreFilter: 0.01,
			Readout:              "Prey",
			ReadoutScoreFilter:   0.05,
			ScoreType:            "lte",
		}

		expected := "Scoring\n" +
			"- smaller scores are better\n" +
			"- score filter for Bait correlation: 0.01\n" +
			"- score filter for Prey correlation: 0.05\n\n"

		var messages strings.Builder
		logCorrelationScoring(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log score settings for gte", func() {
		settings := types.Settings{
			Condition:            "Bait",
			ConditionScoreFilter: 0.01,
			Readout:              "Prey",
			ReadoutScoreFilter:   0.05,
			ScoreType:            "gte",
		}

		expected := "Scoring\n" +
			"- larger scores are better\n" +
			"- score filter for Bait correlation: 0.01\n" +
			"- score filter for Prey correlation: 0.05\n\n"

		var messages strings.Builder
		logCorrelationScoring(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log correlation-specifc settings", func() {
	It("should log settings", func() {
		settings := types.Settings{
			Correlation:             "pearson",
			UseReplicates:           true,
			IgnoreSourceTargetPairs: true,
			CytoscapeCutoff:         0.7,
		}

		expected := "Correlation\n" +
			"- correlation method: pearson\n" +
			"- treat replicates as separate data points: true\n" +
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
