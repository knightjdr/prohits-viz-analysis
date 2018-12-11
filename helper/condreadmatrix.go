package helper

import (
	"sort"
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

type conditionReadout struct {
	readout, condition string
}

type readoutData struct {
	abundance, score float64
}

// scoreFunc returns a function for determining if a score is worse than treshold.
// This is used for finding the worst score to use for missing readouts in the table.
func scoreFunc(scoreType string) func(score, threshold float64) float64 {
	if scoreType == "gte" {
		return func(score, threshold float64) float64 {
			if score < threshold {
				return score
			}
			return threshold
		}
	}
	return func(score, threshold float64) float64 {
		if score > threshold {
			return score
		}
		return threshold
	}
}

// sortLabels sorts a map of labels either alphabetically or by the index
// stored as each entry's value.
// input = []map{
//    {"C label": 2}
//		{"B label": 0},
//		{"A label": 1}
// output alphabetically: []string{"A label", "B label", "C label"}
// output by index: []string{"B label", "A label", "C label"}
func sortLabels(labels map[string]int, alphabetically bool) []string {
	sortedLabels := make([]string, len(labels))
	if alphabetically {
		index := 0
		for label := range labels {
			sortedLabels[index] = label
			index++
		}
		sort.Strings(sortedLabels)
	} else {
		for label, value := range labels {
			sortedLabels[value] = label
		}
	}
	return sortedLabels
}

// ConditionReadoutMatrix generates a 2D matrix with rows (first dimensions) equal to readouts.
// If the rows and columns should be sorted alphabetically, set resort to true. It does this with
// both abundance and the score as the matrix value. It also returns lists of the conditions and
// readouts.
func ConditionReadoutMatrix(table []map[string]string, scoreType string, resort bool) (data typedef.Matrices) {
	// Get scoring function to use for finding the worst score.
	scoreCompare := scoreFunc(scoreType)
	worstScore := float64(0)

	// Put each condition readout pair into a 2D map. Also find the worst score.
	conditions := make(map[string]int, 0)
	readouts := make(map[string]int, 0)
	readoutCondition := make(map[conditionReadout]readoutData)
	for _, row := range table {
		conditionName := row["condition"]
		readoutName := row["readout"]

		// Abundance could be a pipe separated list. Split and sum to accomodate.
		abundance := strings.Split(row["abundance"], "|")
		abundanceSum := float64(0)
		for _, value := range abundance {
			abdFloat, _ := strconv.ParseFloat(value, 64)
			abundanceSum += abdFloat
		}

		// Set readout-condition map value.
		score, _ := strconv.ParseFloat(row["score"], 64)
		readoutCondition[conditionReadout{readoutName, conditionName}] = readoutData{abundance: abundanceSum, score: score}

		// Set worst score.
		worstScore = scoreCompare(score, worstScore)

		// Add condition if unique.
		if _, ok := conditions[conditionName]; !ok {
			conditions[conditionName] = len(conditions)
		}

		// Add readout if unique.
		if _, ok := readouts[readoutName]; !ok {
			readouts[readoutName] = len(readouts)
		}
	}

	// Sort conditions and readouts.
	data.Conditions = sortLabels(conditions, resort)
	data.Readouts = sortLabels(readouts, resort)

	// Iterate over condition and readout lists to create matrix. Missing values are
	// set to zero.
	data.Abundance = make([][]float64, len(data.Readouts)) // Set row capacity.
	data.Score = make([][]float64, len(data.Readouts))     // Set row capacity.
	for i, readout := range data.Readouts {
		data.Abundance[i] = make([]float64, len(data.Conditions)) // Set column capacity.
		data.Score[i] = make([]float64, len(data.Conditions))     // Set column capacity.
		for j, condition := range data.Conditions {
			if value, ok := readoutCondition[conditionReadout{readout, condition}]; ok {
				data.Abundance[i][j] = value.abundance
				data.Score[i][j] = value.score
			} else {
				data.Abundance[i][j] = 0
				data.Score[i][j] = worstScore
			}
		}
	}

	return
}
