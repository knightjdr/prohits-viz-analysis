package filter

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func filterByConditionsAndReadouts(analysis *types.Analysis) {
	if analysis.Settings.ConditionClustering == "none" || analysis.Settings.ReadoutClustering == "none" {
		filterDataByConditionsAndReadouts(analysis)
	}
}

func filterDataByConditionsAndReadouts(analysis *types.Analysis) {
	filterFunc := getConditionAndReadoutFilter(analysis.Settings)

	dataLength := len(analysis.Data)
	for i := dataLength - 1; i >= 0; i-- {
		condition := analysis.Data[i]["condition"]
		readout := analysis.Data[i]["readout"]

		if !filterFunc(condition, readout) {
			analysis.Data = append(analysis.Data[:i], analysis.Data[i+1:]...)
		}
	}
}
