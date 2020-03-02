package convert


// getScoreTest returns a function for determining if a score is worse than treshold.
// This is used for finding the worst score to use for missing readouts in the table.
func getScoreTest(scoreType string) func(float64, float64) float64 {
	if scoreType == "gte" {
		return func(score, currentThreshold float64) float64 {
			if score < currentThreshold {
				return score
			}
			return currentThreshold
		}
	}
	return func(score, currentThreshold float64) float64 {
		if score > currentThreshold {
			return score
		}
		return currentThreshold
	}
}