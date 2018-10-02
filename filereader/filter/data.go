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
) (filtered []map[string]interface{}) {
	filteredConditionReadout := make([]map[string]string, 0)

	// Filter by both conditions and readouts if desired.
	if parameters.ConditionClustering == "conditions" && parameters.ReadoutClustering == "readouts" {
		filteredConditionReadout = ConditionReadout(data, parameters.ConditionList, parameters.ReadoutList)
	} else if parameters.ConditionClustering == "conditions" { // Filter by conditions only.
		filteredConditionReadout = Conditions(data, parameters.ConditionList)
	} else if parameters.ReadoutClustering == "readouts" { // Filter by readouts only.
		filteredConditionReadout = Readouts(data, parameters.ReadoutList)
	} else {
		filteredConditionReadout = data
	}

	// If filteredConditionReadout slice is empty, return error.
	if len(filteredConditionReadout) == 0 {
		// Log message and panic.
		logmessage.CheckError(errors.New("No parsed results matching condition and readout criteria"), true)
	}

	// Filter by score.
	filtered, err := Score(
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
