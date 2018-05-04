package dotplot

import (
	"sort"
	"strconv"
	"strings"
)

type BaitPrey struct {
	Prey, Bait string
}

// BaitPreyMatrix generates a 2D matrix with rows (first dimensions) equal to preys
// sorted alphabetically and columns (2nd dimension) equal to baits sorted
// alphabetically.
func BaitPreyMatrix(data []map[string]interface{}) (matrix [][]float64, baitList []string, preyList []string) {
	// Put each bait prey pair into a 2D map.
	baits := make(map[string]bool, 0)
	preys := make(map[string]bool, 0)
	preyBait := make(map[BaitPrey]float64)
	for _, row := range data {
		baitName := row["bait"].(string)
		preyName := row["prey"].(string)

		// Abundance could be a pipe separated list. Split and sum to accomodate.
		abundance := strings.Split(row["abundance"].(string), "|")
		abundanceSum := float64(0)
		for _, value := range abundance {
			abdFloat, _ := strconv.ParseFloat(value, 64)
			abundanceSum += abdFloat
		}

		// Set prey-bait map value
		preyBait[BaitPrey{preyName, baitName}] = abundanceSum

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
	for bait, _ := range baits {
		baitList = append(baitList, bait)
	}
	for prey, _ := range preys {
		preyList = append(preyList, prey)
	}
	sort.Strings(baitList)
	sort.Strings(preyList)

	// Iterate over bait and prey lists to create matrix. Missing values are
	// set to zero.
	matrix = make([][]float64, len(preyList)) // Set row capacity.
	for i, prey := range preyList {
		matrix[i] = make([]float64, len(baitList)) // Set column capacity.
		for j, bait := range baitList {
			if value, ok := preyBait[BaitPrey{prey, bait}]; ok {
				matrix[i][j] = value
			} else {
				matrix[i][j] = 0
			}
		}
	}

	return
}
