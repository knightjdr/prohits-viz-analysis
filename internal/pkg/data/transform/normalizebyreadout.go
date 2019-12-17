package transform

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/parse"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"strings"
)

func normalizeByReadout(analysis *types.Analysis) {
	readoutAbundanceByCondition := getReadoutAbundance(analysis.Data, analysis.Settings.NormalizationReadout)
	median := calculateMedian(readoutAbundanceByCondition)
	multipliers := calculateMultipliers(readoutAbundanceByCondition, median)
	adjustAbundanceByMultiplier(analysis, "condition", multipliers)
}

func getReadoutAbundance(data []map[string]string, normalizationReadout string) map[string]float64 {
	abundanceByCondition := make(map[string]float64, 0)

	for _, row := range data {
		condition := row["condition"]
		if _, ok := abundanceByCondition[condition]; !ok {
			abundanceByCondition[condition] = 0
		}

		readout := row["readout"]
		if strings.EqualFold(readout, normalizationReadout) {
			abundance := parse.PipeSeparatedStringToMean(row["abundance"])
			abundanceByCondition[condition] = abundance
		}
	}

	return abundanceByCondition
}
