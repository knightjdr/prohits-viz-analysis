package filter

import (
	"strconv"
	"strings"
)

// Score filters a slice map based on the score and minimum abundance.
func Score(
	data []map[string]string,
	primaryFilter float64,
	minimum float64,
	scoreType string,
) (filtered []map[string]interface{}, err error) {
	// Ensure score column is numeric and convert.
	filtered, err = NumericScore(data)
	if err != nil {
		return filtered, err
	}

	// Get filter function.
	filterFunc := FilterFunc(scoreType)

	// Find unique preys passing score.
	preys := make(map[string]bool, 0)
	for i, row := range filtered {
		// Check if score passes.
		passes := filterFunc(filtered[i]["score"].(float64), primaryFilter)

		// Get abundance value. Could be list of pipe-separated values, so
		// get average for comparing against the minimum requried.
		abundance := strings.Split(filtered[i]["abundance"].(string), "|")
		avgAbundance := float64(0)
		for _, abdValue := range abundance {
			abdFloat, _ := strconv.ParseFloat(abdValue, 64)
			avgAbundance += abdFloat
		}
		avgAbundance /= float64(len(abundance))
		if passes && avgAbundance > minimum {
			preyString := row["prey"].(string)
			if _, ok := preys[preyString]; !ok { // only add preys not already present
				preys[preyString] = true
			}
		}
	}

	// Remove preys not passing score.
	filteredlen := len(filtered)
	for i := filteredlen - 1; i >= 0; i-- {
		preyString := filtered[i]["prey"].(string)
		if _, ok := preys[preyString]; !ok {
			filtered = append(filtered[:i], filtered[i+1:]...)
		}
	}

	return
}
