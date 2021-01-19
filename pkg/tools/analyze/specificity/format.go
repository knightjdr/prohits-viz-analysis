package specificity

import (
	"math"
	"sort"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func formatDataForPlot(data map[string]map[string]map[string]float64, settings types.Settings) map[string][]types.ScatterPoint {
	filtered := filterData(data, settings)
	maxPerCondition := defineMaxPerCondition(filtered)
	formatted := formatAsScatterPoints(filtered, maxPerCondition)

	return formatted
}

func filterData(data map[string]map[string]map[string]float64, settings types.Settings) map[string]map[string]map[string]float64 {
	filtered := make(map[string]map[string]map[string]float64, 0)

	passesFilter := getScoreFilterer(settings.ScoreType, settings.PrimaryFilter)
	for condition, conditionData := range data {
		filtered[condition] = make(map[string]map[string]float64, 0)
		for readout, readoutData := range conditionData {
			if passesFilter(readoutData["score"]) {
				filtered[condition][readout] = readoutData
			}
		}
	}

	return filtered
}

func getScoreFilterer(scoreType string, filter float64) func(score float64) bool {
	if scoreType == "gte" {
		return func(score float64) bool {
			return score >= filter
		}
	}
	return func(score float64) bool {
		return score <= filter
	}
}

func defineMaxPerCondition(data map[string]map[string]map[string]float64) map[string]float64 {
	maxPerCondition := make(map[string]float64, len(data))

	infinity := math.Inf(1)
	negInfinity := math.Inf(-1)
	for condition, conditionData := range data {
		maxPerCondition[condition] = negInfinity
		for _, readoutData := range conditionData {
			if readoutData["specificity"] > maxPerCondition[condition] && readoutData["specificity"] != infinity {
				maxPerCondition[condition] = readoutData["specificity"]
			}
		}

		if maxPerCondition[condition] == negInfinity {
			maxPerCondition[condition] = 100
		}
	}

	return maxPerCondition
}

func formatAsScatterPoints(data map[string]map[string]map[string]float64, maxPerCondition map[string]float64) map[string][]types.ScatterPoint {
	formatted := make(map[string][]types.ScatterPoint, len(data))

	infinity := math.Inf(1)
	for condition, conditionData := range data {
		formatted[condition] = make([]types.ScatterPoint, len(conditionData))
		i := 0
		for readout, readoutData := range conditionData {
			color := "#dfcd06"
			y := readoutData["specificity"]
			if y == infinity {
				color = "#6e97ff"
				y = maxPerCondition[condition]
			}

			formatted[condition][i] = types.ScatterPoint{
				Color: color,
				Label: readout,
				X:     readoutData["abundance"],
				Y:     y,
			}
			i++
		}

		sort.Slice(formatted[condition], func(i, j int) bool {
			return formatted[condition][i].Label < formatted[condition][j].Label
		})
	}

	return formatted
}
