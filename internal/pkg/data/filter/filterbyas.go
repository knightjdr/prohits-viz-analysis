package filter

import (
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/parse"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func filterByAbundanceAndScore(analysis *types.Analysis) {
	filterByCriteria := getAbundanceAndScoreFilter(analysis.Settings)

	passingReadouts := make(map[string]bool, 0)
	for _, row := range analysis.Data {
		abundance := parse.PipeSeparatedFloat(row["abundance"])
		score := parseScore(row["score"])

		if filterByCriteria(abundance, score) {
			readout := row["readout"]
			passingReadouts[readout] = true
		}
	}

	removeReadouts(analysis, passingReadouts)
}

func parseScore(score string) float64 {
	parsedScore, err := strconv.ParseFloat(score, 64)
	if err != nil {
		return 0
	}

	return parsedScore
}

func removeReadouts(analysis *types.Analysis, passingReadouts map[string]bool) {
	dataLength := len(analysis.Data)
	for i := dataLength - 1; i >= 0; i-- {
		readout := analysis.Data[i]["readout"]
		if _, ok := passingReadouts[readout]; !ok {
			analysis.Data = append(analysis.Data[:i], analysis.Data[i+1:]...)
		}
	}
}
