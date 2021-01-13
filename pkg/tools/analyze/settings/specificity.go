package settings

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func logSpecificitySettings(messages *strings.Builder, settings types.Settings) {
	logSpecificityMetric(messages, settings)
	logSpecificityAbundance(messages, settings)
	logSpecificityScoring(messages, settings)
}

func logSpecificityMetric(messages *strings.Builder, settings types.Settings) {
	messages.WriteString(fmt.Sprintf("Metric\n- specificity metric: %s\n", settings.SpecificityMetric))
	messages.WriteString("\n")
}

func logSpecificityAbundance(messages *strings.Builder, settings types.Settings) {
	messages.WriteString(fmt.Sprintf("Abundance\n- minimum abundance required: %s\n", float.RemoveTrailingZeros(settings.MinAbundance)))
	messages.WriteString("\n")
}

func logSpecificityScoring(messages *strings.Builder, settings types.Settings) {
	messages.WriteString("Scoring\n")

	if settings.ScoreType == "gte" {
		messages.WriteString("- larger scores are better\n")
	} else {
		messages.WriteString("- smaller scores are better\n")
	}

	messages.WriteString(fmt.Sprintf("- primary filter: %s\n", float.RemoveTrailingZeros(settings.PrimaryFilter)))
	messages.WriteString("\n")
}
