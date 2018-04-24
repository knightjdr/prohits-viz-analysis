// Package filter filters a slice map based on baits, preys and score.
package filter

import (
	"errors"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// Data filters first by bait and prey, then score.
func Data(
	data []map[string]string,
	primaryFilter float64,
	baits []string,
	preys []string,
	scoreType string,
	logFile string,
) (filtered []map[string]interface{}, err error) {
	filteredBaitPrey := make([]map[string]string, 0)
	// Filter by both baits and preys if there are lists for both.
	if (len(baits) > 0) && (len(preys) > 0) {
		filteredBaitPrey = BaitPrey(data, baits, preys)
	} else if len(baits) > 0 { // Filter by baits only.
		filteredBaitPrey = Baits(data, baits)
	} else if len(preys) > 0 { // Filter by preys only.
		filteredBaitPrey = Preys(data, preys)
	} else {
		filteredBaitPrey = data
	}

	// If filteredBaitPrey slice is empty, return error.
	if len(filteredBaitPrey) == 0 {
		err = errors.New("No parsed results matching bait and prey criteria")
		// Log message and return error.
		logmessage.Write(logFile, err.Error())
		return make([]map[string]interface{}, 0), err
	}

	// Filter by score.
	filtered, err = Score(filteredBaitPrey, primaryFilter, scoreType)
	if err != nil {
		// Log message and return error.
		logmessage.Write(logFile, err.Error())
		return
	}

	// If filtered slice is empty, return error.
	if len(filtered) == 0 {
		err = errors.New("No parsed results matching filter criteria")
		// Log message and return error.
		logmessage.Write(logFile, err.Error())
		return
	}

	return
}
