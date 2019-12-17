package transform

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/parse"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func normalizeByTotalSum(analysis *types.Analysis) {
	abundanceByCondition := sumAbundanceByCondition(analysis.Data)
	median := calculateMedian(abundanceByCondition)
	multipliers := calculateMultipliers(abundanceByCondition, median)
	adjustAbundanceByMultiplier(analysis, "condition", multipliers)
}

func sumAbundanceByCondition(data []map[string]string) map[string]float64 {
	abundanceByCondition := make(map[string]float64, 0)

	for _, row := range data {
		condition := row["condition"]
		if _, ok := abundanceByCondition[condition]; !ok {
			abundanceByCondition[condition] = 0
		}

		abundance := parse.PipeSeparatedStringToMean(row["abundance"])
		abundanceByCondition[condition] += abundance
	}

	return abundanceByCondition
}
