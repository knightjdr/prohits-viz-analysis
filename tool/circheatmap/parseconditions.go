package circheatmap

import (
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

func checkScore(scoreType string, cutoff float64) func(float64) bool {
	if scoreType == "lte" {
		return func(value float64) bool {
			return value <= cutoff
		}
	}
	return func(value float64) bool {
		return value >= cutoff
	}
}

func parseConditions(
	table []map[string]string,
	parameters typedef.Parameters,
	readoutMetrics map[string]string,
) ([]string, map[string]bool, map[string]map[string]map[string]float64) {
	scoreFunc := checkScore(parameters.ScoreType, parameters.PrimaryFilter)

	// Get condition and readout names as slice by first creating a hash. Convert
	// condition hash to an array and sort alphabetically.
	conditionMap := make(map[string]bool)
	readoutMap := make(map[string]bool)
	for _, row := range table {
		conditionMap[row["condition"]] = true
		readoutMap[row["readout"]] = true
	}
	conditionNames := make([]string, len(conditionMap))
	i := 0
	for key := range conditionMap {
		conditionNames[i] = key
		i++
	}
	conditionNames = helper.SortStringSlice(conditionNames, "asc")

	// Create map for storing readout data.
	conditionData := make(map[string]map[string]map[string]float64, len(conditionNames))
	for _, name := range conditionNames {
		conditionData[name] = make(map[string]map[string]float64)
	}

	// Iterate over table, only keeping readouts that pass score threshold
	for _, row := range table {
		score, _ := strconv.ParseFloat(row["score"], 64)
		if scoreFunc(score) {
			readoutValues := make(map[string]float64, len(readoutMetrics))
			for metric := range readoutMetrics {
				readoutValues[metric], _ = strconv.ParseFloat(row[metric], 64)
			}
			conditionData[row["condition"]][row["readout"]] = readoutValues
		}
	}

	return conditionNames, readoutMap, conditionData
}
