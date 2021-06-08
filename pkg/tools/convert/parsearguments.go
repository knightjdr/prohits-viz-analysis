package convert

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/flags"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
)

type conversionSettings struct {
	file      string
	imageType string
}

func parseArguments() conversionSettings {
	settings := conversionSettings{}

	args := flags.Parse()
	settings.file = flags.SetString("file", args, "")
	settings.imageType = flags.SetString("imageType", args, "")

	if settings.file == "" {
		log.WriteAndExit("no file specified")
	}
	if settings.imageType == "" {
		log.WriteAndExit("image type not specified")
	}

	return settings
}
