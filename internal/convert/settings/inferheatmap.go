package settings

import (
	"math"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func inferDotplotSettings(csv []map[string]string, settings *types.Settings) {
	settings.EdgeColor = "blue"

	min, max := findMinMax(csv)
	setFillAndMinAbundance(min, settings)
	setAbundanceCap(min, max, settings, true)
}

func inferHeatmapSettings(csv []map[string]string, settings *types.Settings) {
	min, max := findMinMax(csv)
	setFillAndMinAbundance(min, settings)
	setAbundanceCap(min, max, settings, false)
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

func setFillAndMinAbundance(min float64, settings *types.Settings) {
	if min >= 0 {
		settings.FillColor = "blue"
		settings.MinAbundance = 0
	} else {
		settings.FillColor = "blueRed"
		settings.MinAbundance = math.Floor(min)
	}
}

func setAbundanceCap(min, max float64, settings *types.Settings, isDotplot bool) {
	if isDotplot && min >= 0 {
		settings.AbundanceCap = float64(50)
	} else {
		settings.AbundanceCap = float64(math.Ceil(max))
	}
}
