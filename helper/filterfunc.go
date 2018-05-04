// Package helper contains misc. methods used throughout the analysis package.
package helper

// FilterFunc returns a function for comparing scores against a filter.
func FilterFunc(scoreType string) func(float64, float64) bool {
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
