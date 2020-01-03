package export

import (
	"os"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/flags"
	"github.com/knightjdr/prohits-viz-analysis/pkg/slice"
)

type parameters struct {
	format    string
	imageType string
	jsonFile  string
}

func parseArguments() parameters {
	params := parameters{}

	args := flags.Parse()
	params.format = flags.SetString("format", args, "svg")
	params.imageType = flags.SetString("imageType", args, "")
	params.jsonFile = flags.SetString("file", args, "")

	errors := false
	if params.imageType == "" {
		log.Write("image type not specified")
		errors = true
	}
	if params.jsonFile == "" {
		log.Write("no JSON file specified")
		errors = true
	}

	if errors {
		os.Exit(1)
	}

	acceptedOutputFormats := []string{"png", "svg"}
	if !slice.ContainsString(params.format, acceptedOutputFormats) {
		log.WriteAndExit("invalid output format")
	}

	return params
}
