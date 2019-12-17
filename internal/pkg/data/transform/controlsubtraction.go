package transform

import (
	"math"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/parse"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	customMath "github.com/knightjdr/prohits-viz-analysis/pkg/math"
)

func controlSubtract(analysis *types.Analysis) {
	if analysis.Settings.Control == "" {
		return
	}

	for i, row := range analysis.Data {
		abundances := parse.PipeSeparatedStringToArray(row["abundance"])
		controlAverage := parse.PipeSeparatedStringToMean(row["control"])

		controlSubstractedAbundances := make([]float64, 0)
		for _, abundance := range abundances {
			substractedAbundance := math.Max(0, abundance-controlAverage)
			rounded := customMath.Round(substractedAbundance, 0.01)
			controlSubstractedAbundances = append(controlSubstractedAbundances, rounded)
		}
		analysis.Data[i]["abundance"] = float.Join(controlSubstractedAbundances[:], "|")
	}
}
