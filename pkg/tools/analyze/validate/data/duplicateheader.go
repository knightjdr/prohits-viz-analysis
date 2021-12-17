package data

import (
	"errors"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func checkForDuplicateHeader(data []map[string]string, settings types.Settings) (err error) {
	header := map[string]bool{
		settings.Abundance: true,
		settings.Condition: true,
		settings.Readout:   true,
		settings.Score:     true,
	}

	for _, row := range data {
		_, matchAbundance := header[row["abundance"]]
		_, matchCondition := header[row["condition"]]
		_, matchReadout := header[row["readout"]]
		_, matchScore := header[row["score"]]

		if matchAbundance && matchCondition && matchReadout && matchScore {
			err = errors.New(
				"the file should only contain a single header row - duplicates detected; " +
					"this can happen when you manually merge multiple files and do not remove the " +
					"header rows from the additional files",
			)
			return
		}
	}

	return
}
