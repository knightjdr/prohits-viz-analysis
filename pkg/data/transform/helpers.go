package transform

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/parse"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	customMath "github.com/knightjdr/prohits-viz-analysis/pkg/math"
	"github.com/knightjdr/prohits-viz-analysis/pkg/stats"
)

func adjustAbundanceByMultiplier(analysis *types.Analysis, column string, multiplier map[string]float64) {
	for i, row := range analysis.Data {
		multiplierKey := row[column]
		abundances := parse.PipeSeparatedStringToArray(row["abundance"])

		lengthAdjustedAbundances := make([]float64, 0)
		for _, abundance := range abundances {
			adjustedAbundance := abundance * multiplier[multiplierKey]
			rounded := customMath.Round(adjustedAbundance, 0.01)
			lengthAdjustedAbundances = append(lengthAdjustedAbundances, rounded)
		}
		analysis.Data[i]["abundance"] = float.Join(lengthAdjustedAbundances[:], "|")
	}
}

func calculateMedian(dict map[string]float64) float64 {
	values := make([]float64, len(dict))

	i := 0
	for _, value := range dict {
		values[i] = value
		i++
	}

	return stats.MedianFloat(values)
}

func calculateMultipliers(dict map[string]float64, median float64) map[string]float64 {
	multiplier := make(map[string]float64, 0)

	for readout, value := range dict {
		if value != 0 {
			multiplier[readout] = median / value
		} else {
			multiplier[readout] = 1
		}
	}

	return multiplier
}
