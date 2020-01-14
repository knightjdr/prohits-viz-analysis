package filter

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/slice"
)

func getAbundanceAndScoreFilter(settings types.Settings) func(float64, float64) bool {
	filterByAbundance := defineAbundanceFilter(settings.MinAbundance)
	filterByScore := DefineScoreFilter(settings.ScoreType, settings.PrimaryFilter)

	return func(abundance, score float64) bool {
		return filterByAbundance(abundance) && filterByScore(score)
	}
}

func getConditionAndReadoutFilter(settings types.Settings) func(string, string) bool {
	filterByCondition := defineNameFilter(settings.ConditionClustering, settings.ConditionList)
	filterByReadout := defineNameFilter(settings.ReadoutClustering, settings.ReadoutList)

	return func(condition, readout string) bool {
		return filterByCondition(condition) && filterByReadout(readout)
	}
}

func defineAbundanceFilter(minAbundance float64) func(float64) bool {
	return func(abundance float64) bool {
		return abundance >= minAbundance
	}
}

// DefineScoreFilter returns a function for filtering by score.
func DefineScoreFilter(scoreType string, filter float64) func(float64) bool {
	if scoreType == "gte" {
		return func(score float64) bool {
			return score >= filter
		}
	}
	return func(score float64) bool {
		return score <= filter
	}
}

func defineNameFilter(clusteringType string, names []string) func(string) bool {
	if clusteringType != "none" {
		return func(name string) bool {
			return true
		}
	}

	dict := slice.ConvertToMap(names)
	return func(name string) bool {
		if _, ok := dict[name]; ok {
			return true
		}
		return false
	}
}
