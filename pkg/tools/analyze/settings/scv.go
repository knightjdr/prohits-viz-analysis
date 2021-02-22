package settings

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func logSCVSettings(messages *strings.Builder, settings types.Settings) {
	logSCVAbundance(messages, settings)
	logSCVScoring(messages, settings)
	logSCVMapping(messages, settings)
	logSCVKnown(messages, settings)
}

func logSCVAbundance(messages *strings.Builder, settings types.Settings) {
	messages.WriteString(fmt.Sprintf("Abundance\n- column for abundance filtering: %s\n", settings.AbundanceFilterColumn))
	messages.WriteString(fmt.Sprintf("- minimum abundance required: %s\n", float.RemoveTrailingZeros(settings.MinAbundance)))
	if settings.AbundanceCap > 0 {
		messages.WriteString(fmt.Sprintf("- abundances were capped at %s for visualization\n", float.RemoveTrailingZeros(settings.AbundanceCap)))
	}
	messages.WriteString("\n")
}

func logSCVKnown(messages *strings.Builder, settings types.Settings) {
	if settings.Known != "" {
		messages.WriteString("Known metric\n")
		messages.WriteString(fmt.Sprintf("- %s knownness was evaluated by: %s\n", settings.Readout, settings.Known))
		messages.WriteString("\n")
	}
}

func logSCVMapping(messages *strings.Builder, settings types.Settings) {
	messages.WriteString("ID mapping\n")

	messages.WriteString(fmt.Sprintf("- %s ID type: %s\n", settings.Condition, settings.ConditionIDType))
	if settings.ConditionMapColumn != "" {
		messages.WriteString(fmt.Sprintf("- %s IDs were mapped by column: %s\n", settings.Condition, settings.ConditionMapColumn))
	}
	if settings.ConditionMapFile != "" {
		messages.WriteString(fmt.Sprintf("- %s IDs were mapped by file\n", settings.Condition))
	}

	messages.WriteString(fmt.Sprintf("- %s ID type: %s\n", settings.Readout, settings.ReadoutIDType))
	if settings.ReadoutMapColumn != "" {
		messages.WriteString(fmt.Sprintf("- %s IDs were mapped by column: %s\n", settings.Readout, settings.ReadoutMapColumn))
	}
	if settings.ReadoutMapFile != "" {
		messages.WriteString(fmt.Sprintf("- %s IDs were mapped by file\n", settings.Readout))
	}

	messages.WriteString("\n")
}

func logSCVScoring(messages *strings.Builder, settings types.Settings) {
	messages.WriteString("Scoring\n")

	if settings.ScoreType == "gte" {
		messages.WriteString("- larger scores are better\n")
	} else {
		messages.WriteString("- smaller scores are better\n")
	}

	messages.WriteString(fmt.Sprintf("- primary filter: %s\n", float.RemoveTrailingZeros(settings.PrimaryFilter)))
	messages.WriteString("\n")
}
