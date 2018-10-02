package dotplot

import (
	"sort"
	"strconv"
	"strings"
)

type conditionReadout struct {
	readout, condition string
}

// Data holds information about the input data table.
type Data struct {
	Abundance, Score     [][]float64
	Conditions, Readouts []string
}

type readoutData struct {
	abundance, score float64
}

// ScoreFunc returns a function for determining if a score is worse than treshold.
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

// ConditionReadoutMatrix generates a 2D matrix with rows (first dimensions) equal to readouts
// sorted alphabetically and columns (2nd dimension) equal to conditions sorted
// alphabetically. It does this with both abundance and the score as the matrix
// value. It also returns lists of the conditions and readouts.
func ConditionReadoutMatrix(table []map[string]interface{}, scoreType string) (data Data) {
	// Get scoring function to use for finding the worst score.
	scoreCompare := scoreFunc(scoreType)
	worstScore := float64(0)

	// Put each condition readout pair into a 2D map. Also find the worst score.
	conditions := make(map[string]bool, 0)
	readouts := make(map[string]bool, 0)
	readoutCondition := make(map[conditionReadout]readoutData)
	for _, row := range table {
		conditionName := row["condition"].(string)
		readoutName := row["readout"].(string)

		// Abundance could be a pipe separated list. Split and sum to accomodate.
		abundance := strings.Split(row["abundance"].(string), "|")
		abundanceSum := float64(0)
		for _, value := range abundance {
			abdFloat, _ := strconv.ParseFloat(value, 64)
			abundanceSum += abdFloat
		}

		// Set readout-condition map value.
		score, _ := row["score"].(float64)
		readoutCondition[conditionReadout{readoutName, conditionName}] = readoutData{abundance: abundanceSum, score: score}

		// Set worst score.
		worstScore = scoreCompare(score, worstScore)

		// Add condition if unique.
		if _, ok := conditions[conditionName]; !ok {
			conditions[conditionName] = true
		}

		// Add readout if unique.
		if _, ok := readouts[readoutName]; !ok {
			readouts[readoutName] = true
		}
	}

	// Sort conditions and readouts.
	for condition := range conditions {
		data.Conditions = append(data.Conditions, condition)
	}
	for readout := range readouts {
		data.Readouts = append(data.Readouts, readout)
	}
	sort.Strings(data.Conditions)
	sort.Strings(data.Readouts)

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
