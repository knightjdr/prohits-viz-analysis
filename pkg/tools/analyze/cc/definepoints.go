package cc

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/pkg/mapf"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func definePoints(analysis *types.Analysis) []types.ScatterPoint {
	readoutData, readouts := createDictOfReadoutsToData(analysis)
	return convertDataToPoints(readoutData, readouts, analysis.Settings)
}

func createDictOfReadoutsToData(analysis *types.Analysis) (map[string]map[string]float64, []string) {
	conditionToAxis := map[string]string{
		analysis.Settings.ConditionX: "x",
		analysis.Settings.ConditionY: "y",
	}

	initializeDatum := defineDataInitializer(analysis.Settings.ScoreType, analysis.Settings.SecondaryFilter)

	data := make(map[string]map[string]float64, 0)
	readoutDict := make(map[string]bool, 0)
	for _, datum := range analysis.Data {
		abundance := datum["abundance"]
		condition := datum["condition"]
		readout := datum["readout"]
		score := datum["score"]

		if _, ok := data[readout]; !ok {
			data[readout] = initializeDatum()
		}

		data[readout][conditionToAxis[condition]], _ = strconv.ParseFloat(abundance, 64)
		data[readout][fmt.Sprintf("score%s", conditionToAxis[condition])], _ = strconv.ParseFloat(score, 64)
		readoutDict[readout] = true
	}

	readouts := mapf.KeysStringBool(readoutDict)
	sort.Strings(readouts)

	return data, readouts
}

func defineDataInitializer(scoreType string, scoreFilter float64) func() map[string]float64 {
	var defaultScore float64
	if scoreType == "gte" {
		defaultScore = scoreFilter - 0.01
	} else {
		defaultScore = scoreFilter + 0.01
	}

	return func() map[string]float64 {
		return map[string]float64{
			"scorex": defaultScore,
			"scorey": defaultScore,
			"x":      0,
			"y":      0,
		}
	}
}

func convertDataToPoints(data map[string]map[string]float64, readouts []string, settings types.Settings) []types.ScatterPoint {
	numReadouts := len(readouts)
	points := make([]types.ScatterPoint, numReadouts)
	getColor := defineColorGetter(settings)
	for i, readout := range readouts {
		points[i] = types.ScatterPoint{
			Label: readout,
			X:     data[readout]["x"],
			Y:     data[readout]["y"],
		}

		points[i].Color = getColor([]float64{data[readout]["scorex"], data[readout]["scorey"]})
	}

	return points
}

func defineColorGetter(settings types.Settings) func([]float64) string {
	primaryColor := "#0066cc"
	secondaryColor := "#99ccff"

	scoreType := settings.ScoreType
	primaryFilter := settings.PrimaryFilter

	if scoreType == "gte" {
		return func(scores []float64) string {
			score := math.Max(scores[0], scores[1])
			if score >= primaryFilter {
				return primaryColor
			}
			return secondaryColor
		}
	}

	return func(scores []float64) string {
		score := math.Min(scores[0], scores[1])
		if score <= primaryFilter {
			return primaryColor
		}
		return secondaryColor
	}
}
