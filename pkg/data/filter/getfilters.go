package filter

import (
	"math"

	"github.com/knightjdr/prohits-viz-analysis/pkg/slice"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func getAbundanceAndScoreFilter(settings types.Settings) func(float64, float64) bool {
	filterByAbundance := defineAbundanceFilter(settings)
	filterByScore := DefineScoreFilter(settings)

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

func defineAbundanceFilter(settings types.Settings) func(float64) bool {
	minAbundance := defineMinAbundance(settings)
	return func(abundance float64) bool {
		return math.Abs(abundance) >= minAbundance
	}
}

func defineMinAbundance(settings types.Settings) float64 {
	if settings.Type == "correlation" {
		if settings.ConditionAbundanceFilter >= settings.ReadoutAbundanceFilter {
			return settings.ConditionAbundanceFilter
		}
		return settings.ReadoutAbundanceFilter

	}
	return settings.MinAbundance
}

// DefineScoreFilter returns a function for filtering by score.
func DefineScoreFilter(settings types.Settings) func(float64) bool {
	scoreFilter := defineUltimateFilter(settings)
	scoreType := settings.ScoreType
	if scoreType == "gte" {
		return func(score float64) bool {
			return score >= scoreFilter
		}
	}
	return func(score float64) bool {
		return score <= scoreFilter
	}
}

func defineUltimateFilter(settings types.Settings) float64 {
	if settings.Type == "condition-condition" {
		return settings.SecondaryFilter
	}
	if settings.Type == "correlation" {
		if (settings.ScoreType == "lte" && settings.ConditionScoreFilter <= settings.ReadoutScoreFilter) ||
			(settings.ScoreType == "gte" && settings.ConditionScoreFilter >= settings.ReadoutScoreFilter) {
			return settings.ReadoutScoreFilter
		}
		return settings.ConditionScoreFilter

	}
	return settings.PrimaryFilter
}

func defineNameFilter(clusteringType string, names []string) func(string) bool {
	if clusteringType != "none" {
		return func(name string) bool {
			return true
		}
	}

	dict := slice.ConvertToBoolMap(names)
	return func(name string) bool {
		if _, ok := dict[name]; ok {
			return true
		}
		return false
	}
}
