package filter

// returns a function for comparing scores against a filter
func Passes(scoreType string) func(float64, float64) bool {
	gteFilter := func(score float64, filter float64) bool {
		return score >= filter
	}
	lteFilter := func(score float64, filter float64) bool {
		return score <= filter
	}
	if scoreType == "gte" {
		return gteFilter
	}
	return lteFilter
}
