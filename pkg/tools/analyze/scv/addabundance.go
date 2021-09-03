package scv

import (
	"math"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/pkg/data/filter"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func addAbundance(data map[string]map[string]map[string]float64, analysis *types.Analysis) {
	doesScorePassFilter := filter.DefineScoreFilter(analysis.Settings)

	for _, row := range analysis.Data {
		abundance, _ := strconv.ParseFloat(row["abundance"], 64)
		condition := row["condition"]
		readout := row["readout"]
		score, _ := strconv.ParseFloat(row["score"], 64)

		initializeCondition(data, condition)

		if doesScorePassFilter(score) && math.Abs(abundance) >= analysis.Settings.MinAbundance {
			initializeReadout(data[condition], readout)
			data[condition][readout][analysis.Settings.Abundance] = abundance
			for _, column := range analysis.Settings.OtherAbundance {
				data[condition][readout][column], _ = strconv.ParseFloat(row[column], 64)
			}
		}
	}
}

func initializeCondition(data map[string]map[string]map[string]float64, condition string) {
	if _, ok := data[condition]; !ok {
		data[condition] = make(map[string]map[string]float64)
	}
}

func initializeReadout(data map[string]map[string]float64, readout string) {
	if _, ok := data[readout]; !ok {
		data[readout] = make(map[string]float64)
	}
}
