package settings

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func logDotplotSettings(messages *strings.Builder, settings types.Settings) {
	logFiltering(messages, settings)
	logDotplotAbundance(messages, settings)
	logDotplotScoring(messages, settings)
	logDotplotClustering(messages, settings)
}

func logDotplotAbundance(messages *strings.Builder, settings types.Settings) {
	messages.WriteString(fmt.Sprintf("Abundance\n- minimum abundance required: %s\n", float.RemoveTrailingZeros(settings.MinAbundance)))
	if settings.AbundanceCap > 0 {
		messages.WriteString(fmt.Sprintf("- abundances were capped at %s for visualization\n", float.RemoveTrailingZeros(settings.AbundanceCap)))
	}
	messages.WriteString("\n")
}

func logDotplotScoring(messages *strings.Builder, settings types.Settings) {
	messages.WriteString("Scoring\n")

	if settings.ScoreType == "gte" {
		messages.WriteString("- larger scores are better\n")
	} else {
		messages.WriteString("- smaller scores are better\n")
	}

	messages.WriteString(fmt.Sprintf("- primary filter: %s\n", float.RemoveTrailingZeros(settings.PrimaryFilter)))

	if settings.SecondaryFilter != 0 {
		messages.WriteString(fmt.Sprintf("- secondary filter: %s\n", float.RemoveTrailingZeros(settings.SecondaryFilter)))
	}

	messages.WriteString("\n")
}

func logDotplotClustering(messages *strings.Builder, settings types.Settings) {
	messages.WriteString("Clustering\n")
	logBiclustering(messages, settings)
	logHierarchical(messages, settings)
	logNoClustering(messages, settings)
	logClusteringOptimization(messages, settings)
	messages.WriteString("\n")
}
