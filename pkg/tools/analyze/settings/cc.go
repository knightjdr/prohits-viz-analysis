package settings

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func logCCSettings(messages *strings.Builder, settings types.Settings) {
	logCCConditions(messages, settings)
	logCCAbundance(messages, settings)
	logCCScoring(messages, settings)
}

func logCCConditions(messages *strings.Builder, settings types.Settings) {
	messages.WriteString(fmt.Sprintf("Conditions\n- x-axis: %s\n", settings.ConditionX))
	messages.WriteString(fmt.Sprintf("- y-axis: %s\n", settings.ConditionY))
	messages.WriteString("\n")
}

func logCCAbundance(messages *strings.Builder, settings types.Settings) {
	messages.WriteString(fmt.Sprintf("Abundance\n- minimum abundance required: %s\n", float.RemoveTrailingZeros(settings.MinAbundance)))
	messages.WriteString("\n")
}

func logCCScoring(messages *strings.Builder, settings types.Settings) {
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
