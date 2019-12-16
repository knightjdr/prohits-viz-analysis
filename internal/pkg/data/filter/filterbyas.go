package filter

import (
	"errors"
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/stats"
)

func filterByAbundanceAndScore(analysis *types.Analysis) {
	filterByCriteria := getFilterCriteria(analysis.Settings)

	passingReadouts := make(map[string]bool, 0)
	for _, row := range analysis.Data {
		abundance := parseAbundance(row["abundance"])
		score := parseScore(row["score"])

		if filterByCriteria(abundance, score) {
			readout := row["readout"]
			if _, ok := passingReadouts[readout]; !ok {
				passingReadouts[readout] = true
			}
		}
	}

	removeReadouts(analysis, passingReadouts)

	if len(analysis.Data) == 0 {
		err := errors.New("No parsed results matching filter criteria")
		log.CheckError(err, true)
	}
}

func parseAbundance(abundanceString string) float64 {
	abundances := strings.Split(abundanceString, "|")

	parsedAbundances := make([]float64, 0)
	for _, str := range abundances {
		value, err := strconv.ParseFloat(str, 64)
		if err == nil {
			parsedAbundances = append(parsedAbundances, value)
		}
	}

	return stats.MeanFloat(parsedAbundances)
}

func parseScore(score string) float64 {
	parsedScore, err := strconv.ParseFloat(score, 64)
	if err != nil {
		err = errors.New("score column is not numeric")
		log.CheckError(err, true)
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
