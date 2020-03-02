// Package arguments parses and validates command line arguments.
package arguments

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/flags"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/validate/settings"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Parse and validate command line arguments.
func Parse() *types.Analysis {
	analysis := readArguments()
	settings.Validate(analysis)
	return analysis
}

func readArguments() *types.Analysis {
	args := flags.Parse()
	settingsFile := flags.SetString("settings", args, "")

	if settingsFile == "" {
		log.WriteAndExit("no settings file specified")
	}

	return readSettings(settingsFile)
}
