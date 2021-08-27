package filter

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/parse"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func byAbundanceAndScore(analysis *types.Analysis) {
	doesReadoutPassFilters := getAbundanceAndScoreFilter(analysis.Settings)

	passingReadouts := make(map[string]map[string]bool)
	for _, row := range analysis.Data {
		abundance := parse.PipeSeparatedStringToMean(row["abundance"])
		score := parse.Score(row["score"])

		if doesReadoutPassFilters(abundance, score) {
			addReadout(&passingReadouts, row)
		}
	}

	filterFailingReadouts(analysis, passingReadouts)
}

func addReadout(passingReadouts *map[string]map[string]bool, row map[string]string) {
	condition := row["condition"]
	readout := row["readout"]

	if _, ok := (*passingReadouts)[readout]; !ok {
		(*passingReadouts)[readout] = make(map[string]bool)
	}

	(*passingReadouts)[readout][condition] = true
}

func filterFailingReadouts(analysis *types.Analysis, passingReadouts map[string]map[string]bool) {
	shouldRemoveReadout := filterByReadout(passingReadouts, analysis.Settings)
	dataLength := len(analysis.Data)
	for i := dataLength - 1; i >= 0; i-- {
		condition := analysis.Data[i]["condition"]
		readout := analysis.Data[i]["readout"]
		if shouldRemoveReadout(condition, readout) {
			analysis.Data = append(analysis.Data[:i], analysis.Data[i+1:]...)
		}
	}
}

func filterByReadout(passingReadouts map[string]map[string]bool, settings types.Settings) func(string, string) bool {
	if settings.ParsimoniousReadoutFiltering {
		return func(condition, readout string) bool {
			_, okReadout := passingReadouts[readout]
			_, okCondition := passingReadouts[readout][condition]
			if okReadout &&
				okCondition &&
				len(passingReadouts[readout]) >= settings.MinConditions {
				return false
			}
			return true
		}
	}
	return func(condition, readout string) bool {
		if _, ok := passingReadouts[readout]; ok &&
			len(passingReadouts[readout]) >= settings.MinConditions {
			return false
		}
		return true
	}
}
