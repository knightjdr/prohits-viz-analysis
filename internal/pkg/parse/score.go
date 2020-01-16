package parse

import "strconv"

// Score from a string value.
func Score(score string) float64 {
	parsedScore, err := strconv.ParseFloat(score, 64)
	if err != nil {
		return 0
	}

	return parsedScore
}
