package filter

import (
	"errors"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/slice"
)

func filterByConditionsAndReadouts(analysis *types.Analysis) {
	if analysis.Settings.ConditionClustering == "none" || analysis.Settings.ReadoutClustering == "none" {
		filterDataByConditionsAndReadouts(analysis)
	}
}

func filterDataByConditionsAndReadouts(analysis *types.Analysis) {
	conditionMap := slice.ConvertToMap(analysis.Settings.ConditionList)
	readoutMap := slice.ConvertToMap(analysis.Settings.ReadoutList)

	dataLength := len(analysis.Data)
	for i := dataLength - 1; i >= 0; i-- {
		condition := analysis.Data[i]["condition"]
		readout := analysis.Data[i]["readout"]

		_, okCondition := conditionMap[condition]
		_, okReadout := readoutMap[readout]
		if !okCondition || !okReadout {
			analysis.Data = append(analysis.Data[:i], analysis.Data[i+1:]...)
		}
	}

	if len(analysis.Data) == 0 {
		err := errors.New("No parsed results matching condition and readout criteria")
		log.CheckError(err, true)
	}
}
