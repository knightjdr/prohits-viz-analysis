// Package arguments parses and validates command line arguments.
package arguments

import (
	"os"

	settingValidation "github.com/knightjdr/prohits-viz-analysis/internal/analyze/validate/settings"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/flags"
)

// Parse and validate command line arguments.
func Parse() types.Analysis {
	analysis := types.Analysis{}
	analysis.Type, analysis.Settings = readArguments()
	analysis.Columns, analysis.Settings = settingValidation.Validate(analysis)
	return analysis
}

func readArguments() (string, interface{}) {
	args := flags.Parse()
	analysisType := flags.SetString("analysisType", args, "")
	settingsFile := flags.SetString("settings", args, "")

	isMissingArgument := false
	if analysisType == "" {
		log.Write("No analysis type specified")
		isMissingArgument = true
	}
	if settingsFile == "" {
		log.Write("No settings file specified")
		isMissingArgument = true
	}

	if isMissingArgument {
		os.Exit(1)
	}

	settings := readSettings(analysisType, settingsFile)

	return analysisType, settings
}
