package circheatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

func metricKeys(hash map[string]string) []string {
	keys := make([]string, len(hash))

	i := 0
	for key := range hash {
		keys[i] = key
		i++
	}
	return keys
}

func segmentSettings(metrics map[string]string, parameters typedef.Parameters) ([]string, []typedef.CircHeatmapSetttings) {
	// Order metrics so that abundance is first and others are sorted alphabetically.
	metricReverseMap := make(map[string]string, len(metrics)-1)
	otherMetrics := make([]string, len(metrics)-1)
	numOtherMetrics := 0
	for key, metricName := range metrics {
		if key != "abundance" {
			metricReverseMap[metricName] = key
			otherMetrics[numOtherMetrics] = metricName
			numOtherMetrics++
		}
	}
	otherMetrics = helper.SortStringSlice(otherMetrics, "asc")
	sortedMetricKeys := make([]string, len(metrics))
	sortedMetricKeys[0] = "abundance"
	numSortedMetricKeys := 1
	for _, metricName := range otherMetrics {
		sortedMetricKeys[numSortedMetricKeys] = metricReverseMap[metricName]
		numSortedMetricKeys++
	}

	// Create settings.
	settings := make([]typedef.CircHeatmapSetttings, len(metrics))
	for index, key := range sortedMetricKeys {
		settings[index] = typedef.CircHeatmapSetttings{
			AbundanceCap: parameters.AbundanceCap,
			Color:        parameters.FillColor,
			MinAbundance: parameters.MinAbundance,
			Name:         metrics[key],
		}
	}

	return sortedMetricKeys, settings
}
