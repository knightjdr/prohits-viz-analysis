package settings

import (
	"os"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func validateFileSettings(settings types.Settings) map[string]string {
	validateColumns(settings)
	return createColumnMap(settings)
}

func validateColumns(settings types.Settings) {
	hasFileError := false
	if settings.Abundance == "" {
		log.Write("No abundance column specified")
		hasFileError = true
	}
	if settings.Condition == "" {
		log.Write("No condition column specified")
		hasFileError = true
	}
	if len(settings.Files) == 0 {
		log.Write("No input file specified")
		hasFileError = true
	}
	if settings.Readout == "" {
		log.Write("No readout column specified")
		hasFileError = true
	}
	if settings.Score == "" {
		log.Write("No score column specified")
		hasFileError = true
	}

	if hasFileError {
		os.Exit(1)
	}
}

func createColumnMap(fileSettings types.Settings) map[string]string {
	return map[string]string{
		"abundance":     fileSettings.Abundance,
		"condition":     fileSettings.Condition,
		"control":       fileSettings.Control,
		"readout":       fileSettings.Readout,
		"readoutLength": fileSettings.ReadoutLength,
		"score":         fileSettings.Score,
	}
}
