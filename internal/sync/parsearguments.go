package sync

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/flags"
)

func parseArguments() string {
	args := flags.Parse()
	jsonFile := flags.SetString("file", args, "")

	if jsonFile == "" {
		log.WriteAndExit("no JSON file specified")
	}

	return jsonFile
}
