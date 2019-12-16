package filter

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func getFilterCriteria(settings types.Settings) func(float64, float64) bool {
	filterByAbundance := defineAbundanceFilter(settings.MinAbundance)
	filterByScore := defineScoreFilter(settings.ScoreType, settings.PrimaryFilter)

	return func(abundance, score float64) bool {
		return filterByAbundance(abundance) && filterByScore(score)
	}
}

func defineAbundanceFilter(minAbundance float64) func(float64) bool {
	return func(abundance float64) bool {
		return abundance >= minAbundance
	}
}

func defineScoreFilter(scoreType string, filter float64) func(float64) bool {
	gteFilter := func(score float64) bool {
		return score >= filter
	}
	lteFilter := func(score float64) bool {
		return score <= filter
	}

	if scoreType == "gte" {
		return gteFilter
	}
	return lteFilter
}
