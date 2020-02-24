package settings

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
)

func logCorrelationSettings(messages *strings.Builder, settings types.Settings) {
	logFiltering(messages, settings)
	logCorrelationAbundance(messages, settings)
	logCorrelationScoring(messages, settings)
	logCorrelationDetails(messages, settings)
	logCorrelationClustering(messages, settings)
}

func logCorrelationAbundance(messages *strings.Builder, settings types.Settings) {
	messages.WriteString(
		fmt.Sprintf(
			"Abundance\n"+
				"- minimum abundance for %s correlation: %s\n"+
				"- minimum abundance for %s correlation: %s\n\n",
			settings.Condition,
			float.RemoveTrailingZeros(settings.ConditionAbundanceFilter),
			settings.Readout,
			float.RemoveTrailingZeros(settings.ReadoutAbundanceFilter),
		),
	)
}

func logCorrelationScoring(messages *strings.Builder, settings types.Settings) {
	messages.WriteString("Scoring\n")

	if settings.ScoreType == "gte" {
		messages.WriteString("- larger scores are better\n")
	} else {
		messages.WriteString("- smaller scores are better\n")
	}

	messages.WriteString(
		fmt.Sprintf(
			"- score filter for %s correlation: %s\n"+
				"- score filter for %s correlation: %s\n\n",
			settings.Condition,
			float.RemoveTrailingZeros(settings.ConditionScoreFilter),
			settings.Readout,
			float.RemoveTrailingZeros(settings.ReadoutScoreFilter),
		),
	)
}

func logCorrelationDetails(messages *strings.Builder, settings types.Settings) {
	messages.WriteString(
		fmt.Sprintf(
			"Correlation\n"+
				"- correlation method: %s\n"+
				"- treat replicates as separate data points: %t\n"+
				"- ignore source genes in pairwise correlations: %t\n"+
				"- cytoscape cutoff: %s\n\n",
			settings.Correlation,
			settings.UseReplicates,
			settings.IgnoreSourceTargetMatches,
			float.RemoveTrailingZeros(settings.CytoscapeCutoff),
		),
	)
}

func logCorrelationClustering(messages *strings.Builder, settings types.Settings) {
	messages.WriteString("Clustering\n")
	logHierarchical(messages, settings)
	logClusteringOptimization(messages, settings)
	messages.WriteString("\n")
}
