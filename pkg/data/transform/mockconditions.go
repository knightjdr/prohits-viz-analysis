package transform

import (
	"sort"

	"github.com/knightjdr/prohits-viz-analysis/pkg/data/filter"
	"github.com/knightjdr/prohits-viz-analysis/pkg/parse"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/mapf"
)

func mockConditionAbundance(analysis *types.Analysis) {
	if analysis.Settings.MockConditionAbundance {
		maxAbundance := getMaxAbundance(analysis.Data, analysis.Settings)
		addAbundance(analysis, maxAbundance)
	}
}

func getMaxAbundance(data []map[string]string, settings types.Settings) map[string]map[string]float64 {
	filterByScore := filter.DefineScoreFilter(settings)

	maxAbundance := make(map[string]map[string]float64, 0)
	for _, row := range data {
		score := parse.Score(row["score"])
		if filterByScore(score) {
			abundance := parse.PipeSeparatedStringToMean(row["abundance"])
			condition := row["condition"]
			if _, ok := maxAbundance[condition]; !ok || maxAbundance[condition]["abundance"] < abundance {
				maxAbundance[condition] = map[string]float64{
					"abundance": abundance,
					"score":     score,
				}
			}
		}
	}

	return maxAbundance
}

func addAbundance(analysis *types.Analysis, maxAbundance map[string]map[string]float64) {
	uniqueConditionsWithAbundances := make(map[string]bool, 0)
	for _, row := range analysis.Data {
		condition := row["condition"]
		readout := row["readout"]
		if condition == readout {
			uniqueConditionsWithAbundances[condition] = true
		}
		if condition != readout && uniqueConditionsWithAbundances[condition] != true {
			uniqueConditionsWithAbundances[condition] = false
		}
	}

	conditions := mapf.KeysStringBool(uniqueConditionsWithAbundances)
	sort.Strings(conditions)
	for _, condition := range conditions {
		values := maxAbundance[condition]
		if !uniqueConditionsWithAbundances[condition] {
			row := map[string]string{
				"abundance": float.RemoveTrailingZeros(values["abundance"]),
				"condition": condition,
				"readout":   condition,
				"score":     float.RemoveTrailingZeros(values["score"]),
			}
			analysis.Data = append(analysis.Data, row)
		}
	}
}
