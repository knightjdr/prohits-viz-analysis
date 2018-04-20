// Package filter will filter an array map based on baits and preys
package filter

import "errors"

// Data filters first by bait and prey, then score
func Data(
	data []map[string]string,
	primaryFilter float64,
	baits []string,
	preys []string,
	scoreType string,
) ([]map[string]string, error) {
	var filtered = make([]map[string]string, 0)
	// filter by both baits and preys if there are lists for both
	if (len(baits) > 0) && (len(preys) > 0) {
		filtered = Baitprey(data, baits, preys)
	} else if len(baits) > 0 { // filter by baits only
		filtered = Baits(data, baits)
	} else if len(preys) > 0 { // filter by preys only
		filtered = Preys(data, preys)
	} else {
		filtered = data
	}

	// if parsed array is empty, return error
	var err error
	if len(filtered) == 0 {
		err = errors.New("No parsed results")
	}
	return filtered, err
}
