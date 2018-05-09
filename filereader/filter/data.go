// Package filter filters a slice map based on baits, preys and score.
package filter

import (
	"errors"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Data filters first by bait and prey, then score.
func Data(
	data []map[string]string,
	params typedef.Parameters,
) (filtered []map[string]interface{}) {
	filteredBaitPrey := make([]map[string]string, 0)

	// Filter by both baits and preys if there are lists for both.
	if (len(params.BaitList) > 0) && (len(params.PreyList) > 0) {
		filteredBaitPrey = BaitPrey(data, params.BaitList, params.PreyList)
	} else if len(params.BaitList) > 0 { // Filter by baits only.
		filteredBaitPrey = Baits(data, params.BaitList)
	} else if len(params.PreyList) > 0 { // Filter by preys only.
		filteredBaitPrey = Preys(data, params.PreyList)
	} else {
		filteredBaitPrey = data
	}

	// If filteredBaitPrey slice is empty, return error.
	if len(filteredBaitPrey) == 0 {
		// Log message and panic.
		logmessage.CheckError(errors.New("No parsed results matching bait and prey criteria"), true)
	}

	// Filter by score.
	filtered, err := Score(
		filteredBaitPrey,
		params.PrimaryFilter,
		params.MinimumAbundance,
		params.ScoreType,
	)
	logmessage.CheckError(err, true)

	// If filtered slice is empty, return error.
	if len(filtered) == 0 {
		// Log message and panics.
		logmessage.CheckError(errors.New("No parsed results matching filter criteria"), true)
	}

	return
}
