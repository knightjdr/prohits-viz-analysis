package transform

import (
	"math"

	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	customMath "github.com/knightjdr/prohits-viz-analysis/pkg/math"
	"github.com/knightjdr/prohits-viz-analysis/pkg/parse"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func logTransform(analysis *types.Analysis) {
	logBase := analysis.Settings.LogBase
	tool := analysis.Settings.Type
	if !shouldLogTransform(logBase, tool) {
		return
	}

	logFunc := getLogFunction(logBase)
	adjustAbundanceByLog(analysis, logFunc)
}

func shouldLogTransform(logBase, tool string) bool {
	if (logBase == "2" || logBase == "e" || logBase == "10") &&
		(tool == "correlation" || tool == "dotplot") {
		return true
	}
	return false
}

func getLogFunction(base string) func(float64) float64 {
	if base == "2" {
		return func(num float64) float64 {
			if num <= 0 {
				return 0
			}
			return math.Log2(num)
		}
	}
	if base == "10" {
		return func(num float64) float64 {
			if num <= 0 {
				return 0
			}
			return math.Log10(num)
		}
	}
	return func(num float64) float64 {
		if num <= 0 {
			return 0
		}
		return math.Log(num)
	}
}

func adjustAbundanceByLog(analysis *types.Analysis, logfunc func(float64) float64) {
	for i, row := range analysis.Data {
		abundances := parse.PipeSeparatedStringToArray(row["abundance"])
		logAdjustedAbundances := make([]float64, 0)
		for _, abundance := range abundances {
			adjustedAbundance := logfunc(abundance)
			rounded := customMath.Round(adjustedAbundance, 0.01)
			logAdjustedAbundances = append(logAdjustedAbundances, rounded)
		}
		analysis.Data[i]["abundance"] = float.Join(logAdjustedAbundances[:], "|")
	}
}
