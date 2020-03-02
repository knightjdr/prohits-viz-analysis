package convert

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/flags"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
)

func parseArguments() string {
	args := flags.Parse()
	file := flags.SetString("file", args, "")

	if file == "" {
		log.WriteAndExit("no file specified")
	}

	return file
}
