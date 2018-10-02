package dotplot

import (
	"sort"
	"strconv"
	"strings"
)

type baitPrey struct {
	prey, bait string
}

// Data holds information about the input data table.
type Data struct {
	Abundance, Score [][]float64
	Baits, Preys     []string
}

type preyData struct {
	abundance, score float64
}

// ScoreFunc returns a function for determining if a score is worse than treshold.
// This is used for finding the worst score to use for missing preys in the table.
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

// BaitPreyMatrix generates a 2D matrix with rows (first dimensions) equal to preys
// sorted alphabetically and columns (2nd dimension) equal to baits sorted
// alphabetically. It does this with both abundance and the score as the matrix
// value. It also returns lists of the baits and preys.
func BaitPreyMatrix(table []map[string]interface{}, scoreType string) (data Data) {
	// Get scoring function to use for finding the worst score.
	scoreCompare := scoreFunc(scoreType)
	worstScore := float64(0)

	// Put each bait prey pair into a 2D map. Also find the worst score.
	baits := make(map[string]bool, 0)
	preys := make(map[string]bool, 0)
	preyBait := make(map[baitPrey]preyData)
	for _, row := range table {
		baitName := row["bait"].(string)
		preyName := row["prey"].(string)

		// Abundance could be a pipe separated list. Split and sum to accomodate.
		abundance := strings.Split(row["abundance"].(string), "|")
		abundanceSum := float64(0)
		for _, value := range abundance {
			abdFloat, _ := strconv.ParseFloat(value, 64)
			abundanceSum += abdFloat
		}

		// Set prey-bait map value.
		score, _ := row["score"].(float64)
		preyBait[baitPrey{preyName, baitName}] = preyData{abundance: abundanceSum, score: score}

		// Set worst score.
		worstScore = scoreCompare(score, worstScore)

		// Add bait if unique.
		if _, ok := baits[baitName]; !ok {
			baits[baitName] = true
		}

		// Add prey if unique.
		if _, ok := preys[preyName]; !ok {
			preys[preyName] = true
		}
	}

	// Sort baits and preys.
	for bait := range baits {
		data.Baits = append(data.Baits, bait)
	}
	for prey := range preys {
		data.Preys = append(data.Preys, prey)
	}
	sort.Strings(data.Baits)
	sort.Strings(data.Preys)

	// Iterate over bait and prey lists to create matrix. Missing values are
	// set to zero.
	data.Abundance = make([][]float64, len(data.Preys)) // Set row capacity.
	data.Score = make([][]float64, len(data.Preys))     // Set row capacity.
	for i, prey := range data.Preys {
		data.Abundance[i] = make([]float64, len(data.Baits)) // Set column capacity.
		data.Score[i] = make([]float64, len(data.Baits))     // Set column capacity.
		for j, bait := range data.Baits {
			if value, ok := preyBait[baitPrey{prey, bait}]; ok {
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
