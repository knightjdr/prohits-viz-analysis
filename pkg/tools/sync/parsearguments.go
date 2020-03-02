package sync

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/flags"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
)

func parseArguments() string {
	args := flags.Parse()
	jsonFile := flags.SetString("file", args, "")

	if jsonFile == "" {
		log.WriteAndExit("no JSON file specified")
	}

	return jsonFile
}
