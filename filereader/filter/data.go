// Package filter filters a slice map based on conditions, readouts and score.
package filter

import (
	"errors"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Data filters first by condition and readout, then score.
func Data(
	data []map[string]string,
	parameters typedef.Parameters,
) (filtered []map[string]string) {
	filteredConditionReadout := make([]map[string]string, 0)

	// Filter by both conditions and readouts if desired, or either alone.
	if parameters.ConditionClustering == "none" && parameters.ReadoutClustering == "none" {
		filteredConditionReadout = ConditionReadout(data, parameters.ConditionList, parameters.ReadoutList)
	} else if parameters.ConditionClustering == "none" {
		filteredConditionReadout = Conditions(data, parameters.ConditionList)
	} else if parameters.ReadoutClustering == "none" {
		filteredConditionReadout = Readouts(data, parameters.ReadoutList)
	} else {
		filteredConditionReadout = data
	}

	// If filteredConditionReadout slice is empty, return error.
	if len(filteredConditionReadout) == 0 {
		logmessage.CheckError(errors.New("No parsed results matching condition and readout criteria"), true)
	}

	// Filter by score.
	filtered, err := Reduce(
		filteredConditionReadout,
		parameters.PrimaryFilter,
		parameters.MinAbundance,
		parameters.ScoreType,
	)
	logmessage.CheckError(err, true)

	// If filtered slice is empty, return error.
	if len(filtered) == 0 {
		// Log message and panics.
		logmessage.CheckError(errors.New("No parsed results matching filter criteria"), true)
	}

	return
}
