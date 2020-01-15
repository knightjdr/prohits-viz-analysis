package settings

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func logSharedSettings(messages *strings.Builder, settings types.Settings) {
	logAnalysisType(messages, settings.Type)
	logFiles(messages, settings.Files)
	logColumns(messages, settings)
	logTransformations(messages, settings)
}

func logAnalysisType(messages *strings.Builder, analysisType string) {
	messages.WriteString(fmt.Sprintf("Analysis type: %s\n\n", analysisType))
}

func logFiles(messages *strings.Builder, inputFiles []string) {
	messages.WriteString("Files used\n")

	fileNames := files.ParseBaseNames(inputFiles)
	for _, file := range fileNames {
		messages.WriteString(fmt.Sprintf("- %s\n", file))
	}
	messages.WriteString("\n")
}

func logColumns(messages *strings.Builder, settings types.Settings) {
	messages.WriteString(
		fmt.Sprintf(
			"Columns used\n- abundance: %s\n- condition: %s\n- readout: %s\n- score: %s\n",
			settings.Abundance,
			settings.Condition,
			settings.Readout,
			settings.Score,
		),
	)

	if settings.Control != "" {
		messages.WriteString(fmt.Sprintf("- control: %s\n", settings.Control))
	}
	if settings.ReadoutLength != "" {
		messages.WriteString(fmt.Sprintf("- readout length: %s\n", settings.ReadoutLength))
	}

	messages.WriteString("\n")
}

func logTransformations(messages *strings.Builder, settings types.Settings) {
	messages.WriteString("Readout abundance transformations\n")

	if settings.Control != "" {
		messages.WriteString("- control subtraction was performed\n")
	}
	if settings.ReadoutLength != "" {
		messages.WriteString("- readout length normalization was performed\n")
	}
	if settings.Normalization == "total" {
		messages.WriteString("- condition normalization was performed using total abundance\n")
	}
	if settings.Normalization == "readout" {
		messages.WriteString(fmt.Sprintf("- condition normalization was performed using the readout: %s\n", settings.NormalizationReadout))
	}
	if settings.LogBase != "" {
		messages.WriteString(fmt.Sprintf("- data was log-transformed with base %s\n", settings.LogBase))
	}

	messages.WriteString("\n")
}
