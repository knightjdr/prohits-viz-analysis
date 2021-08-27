package settings

import (
	"math"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func inferDotplotSettings(csv []map[string]string, settings *types.Settings) {
	settings.EdgeColor = "blue"

	min, max := findMinMax(csv)
	setAbundanceType(min, settings)
	setFillParameters(min, max, settings)
}

func inferHeatmapSettings(csv []map[string]string, settings *types.Settings) {
	min, max := findMinMax(csv)
	setAbundanceType(min, settings)
	setFillParameters(min, max, settings)
}

func findMinMax(csv []map[string]string) (float64, float64) {
	max := -math.MaxFloat64
	min := math.MaxFloat64
	for _, entry := range csv {
		value, _ := strconv.ParseFloat(entry["abundance"], 64)
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return min, max
}

func setAbundanceType(min float64, settings *types.Settings) {
	if min >= 0 {
		settings.AbundanceType = "positive"
		settings.FillColor = "blue"
	} else {
		settings.AbundanceType = "bidirectional"
		settings.FillColor = "blueRed"
	}
}

func setFillParameters(min, max float64, settings *types.Settings) {
	settings.MinAbundance = 0
	if min >= 0 && max > 1 {
		settings.AbundanceCap = 50
		settings.FillMax = 50
		settings.FillMin = 0
		return
	}
	if min >= 0 && max <= 1 {
		settings.AbundanceCap = 1
		settings.FillMax = 1
		settings.FillMin = 0
		return
	}
	ceiling := math.Ceil(max)
	settings.AbundanceCap = ceiling
	settings.FillMax = ceiling
	settings.FillMin = math.Floor(min)
}
