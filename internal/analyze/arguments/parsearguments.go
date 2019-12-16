// Package arguments parses and validates command line arguments.
package arguments

import (
	"os"

	"github.com/knightjdr/prohits-viz-analysis/internal/analyze/validate/settings"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/flags"
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

	isMissingArgument := false
	if settingsFile == "" {
		log.Write("No settings file specified")
		isMissingArgument = true
	}

	if isMissingArgument {
		os.Exit(1)
	}

	return readSettings(settingsFile)
}
