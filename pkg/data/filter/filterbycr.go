package filter

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func byConditionsAndReadouts(analysis *types.Analysis) {
	if analysis.Settings.Type == "condition-condition" {
		filterDataByConditions(analysis)
	}
	if analysis.Settings.Type == "dotplot" &&
		analysis.Settings.Clustering == "none" &&
		(analysis.Settings.ConditionClustering == "none" || analysis.Settings.ReadoutClustering == "none") {
		filterDataByConditionsAndReadouts(analysis)
	}
}

func filterDataByConditions(analysis *types.Analysis) {
	dataLength := len(analysis.Data)
	for i := dataLength - 1; i >= 0; i-- {
		condition := analysis.Data[i]["condition"]

		if condition != analysis.Settings.ConditionX && condition != analysis.Settings.ConditionY {
			analysis.Data = append(analysis.Data[:i], analysis.Data[i+1:]...)
		}
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
