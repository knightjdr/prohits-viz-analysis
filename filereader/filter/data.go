// Package filter filters a slice map based on baits, preys and score
package filter

import (
	"errors"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// Data filters first by bait and prey, then score
func Data(
	data []map[string]string,
	primaryFilter float64,
	baits []string,
	preys []string,
	scoreType string,
	logFile string,
) ([]map[string]interface{}, error) {
	filtered := make([]map[string]interface{}, 0)
	filteredBaitPrey := make([]map[string]string, 0)
	// filter by both baits and preys if there are lists for both
	if (len(baits) > 0) && (len(preys) > 0) {
		filteredBaitPrey = BaitPrey(data, baits, preys)
	} else if len(baits) > 0 { // filter by baits only
		filteredBaitPrey = Baits(data, baits)
	} else if len(preys) > 0 { // filter by preys only
		filteredBaitPrey = Preys(data, preys)
	} else {
		filteredBaitPrey = data
	}

	// if filteredBaitPrey slice is empty, return error
	if len(filteredBaitPrey) == 0 {
		filteredErr := errors.New("No parsed results matching bait and prey criteria")
		// log message and return error
		logmessage.Write(logFile, filteredErr.Error())
		return filtered, filteredErr
	}

	// filter by score
	filtered, err := Score(filteredBaitPrey, primaryFilter, scoreType)
	if err != nil {
		// log message and return error
		logmessage.Write(logFile, err.Error())
		return filtered, err
	}

	// if filtered slice is empty, return error
	if len(filtered) == 0 {
		filteredErr := errors.New("No parsed results matching filter criteria")
		// log message and return error
		logmessage.Write(logFile, filteredErr.Error())
		return filtered, filteredErr
	}

	return filtered, err
}
