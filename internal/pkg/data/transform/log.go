package transform

import (
	"math"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/parse"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	customMath "github.com/knightjdr/prohits-viz-analysis/pkg/math"
)

func logTransform(analysis *types.Analysis) {
	base := analysis.Settings.LogBase
	validLogValues := map[string]bool{
		"2":  true,
		"e":  true,
		"10": true,
	}
	if _, ok := validLogValues[base]; !ok {
		return
	}

	logFunc := getLogFunction(base)
	adjustAbundanceByLog(analysis, logFunc)
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
