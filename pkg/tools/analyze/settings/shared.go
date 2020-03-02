package settings

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
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
	var buffer strings.Builder

	if settings.Control != "" {
		buffer.WriteString("- control subtraction was performed\n")
	}
	if settings.ReadoutLength != "" {
		buffer.WriteString(fmt.Sprintf("- %s length normalization was performed\n", settings.Readout))
	}
	if settings.Normalization == "total" {
		buffer.WriteString(fmt.Sprintf("- %s normalization was performed using total abundance\n", settings.Condition))
	}
	if settings.Normalization == "readout" {
		buffer.WriteString(
			fmt.Sprintf(
				"- %s normalization was performed using the %s: %s\n",
				settings.Condition,
				settings.Readout,
				settings.NormalizationReadout,
			),
		)
	}
	if settings.LogBase != "" && settings.LogBase != "none" {
		buffer.WriteString(fmt.Sprintf("- data was log-transformed with base %s\n", settings.LogBase))
	}
	if settings.MockConditionAbundance == true {
		buffer.WriteString(fmt.Sprintf("- abundance values were mocked for %s(s) with missing values\n", settings.Condition))
	}

	if buffer.String() != "" {
		messages.WriteString(fmt.Sprintf("%s abundance transformations\n", settings.Readout))
		messages.WriteString(buffer.String())
		messages.WriteString("\n")
	}
}
