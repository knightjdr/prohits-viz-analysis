package convert

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/flags"
)

func parseArguments() string {
	args := flags.Parse()
	file := flags.SetString("file", args, "")

	if file == "" {
		log.WriteAndExit("no file specified")
	}

	return file
}
