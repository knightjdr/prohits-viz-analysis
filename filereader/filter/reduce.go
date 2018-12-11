package filter

import (
	"errors"
	"strconv"
	"strings"
)

// Reduce filters a slice map based on the score and minimum abundance.
func Reduce(
	data []map[string]string,
	primaryFilter float64,
	minimum float64,
	scoreType string,
) (filtered []map[string]string, err error) {
	// Check if first datum score is numeric, if not return err.
	_, err = strconv.ParseFloat(data[0]["score"], 64)
	if err != nil {
		err = errors.New("Score column is not numeric")
		return
	}

	// Get filter function.
	filterFunc := Scores(scoreType)

	// Find unique readouts passing score.
	readouts := make(map[string]bool, 0)
	for _, row := range data {
		// Check if score passes.
		score, _ := strconv.ParseFloat(row["score"], 64)
		passes := filterFunc(score, primaryFilter)

		// Get abundance value. Could be list of pipe-separated values, so
		// get average for comparing against the minimum requried.
		abundance := strings.Split(row["abundance"], "|")
		avgAbundance := float64(0)
		for _, abdValue := range abundance {
			abdFloat, _ := strconv.ParseFloat(abdValue, 64)
			avgAbundance += abdFloat
		}
		avgAbundance /= float64(len(abundance))
		if passes && avgAbundance > minimum {
			readout := row["readout"]
			if _, ok := readouts[readout]; !ok { // only add readouts not already present
				readouts[readout] = true
			}
		}
	}

	// Remove readouts not passing score in at least one condition.
	filtered = data
	filteredlen := len(data)
	for i := filteredlen - 1; i >= 0; i-- {
		readout := data[i]["readout"]
		if _, ok := readouts[readout]; !ok {
			filtered = append(filtered[:i], filtered[i+1:]...)
		}
	}

	return
}
