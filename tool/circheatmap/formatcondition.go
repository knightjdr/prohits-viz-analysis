package circheatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

func formatReadout(known bool) func(name string, readoutData map[string]float64) map[string]interface{} {
	if known {
		return func(name string, readoutData map[string]float64) map[string]interface{} {
			return map[string]interface{}{
				"name":  name,
				"known": readoutData["known"] != 0,
			}
		}
	}
	return func(name string, readoutData map[string]float64) map[string]interface{} {
		return map[string]interface{}{
			"name": name,
		}
	}
}

func readoutKeys(hash map[string]map[string]float64) []string {
	keys := make([]string, len(hash))

	i := 0
	for key := range hash {
		keys[i] = key
		i++
	}
	return keys
}

// formatCondition converts data for a condition (containing readouts with metrics)
// to structure for plotting
func formatCondition(
	name string,
	data map[string]map[string]float64,
	known bool,
	metricOrder []string,
	readoutMetrics map[string]string,
) typedef.CircHeatmapPlot {
	segments := make([]typedef.CircHeatmapSegments, len(metricOrder))
	for index, metricName := range metricOrder {
		segments[index] = typedef.CircHeatmapSegments{
			Name:   readoutMetrics[metricName],
			Values: make([]float64, len(data)),
		}
	}

	readoutFunc := formatReadout(known)
	sortedReadouts := helper.SortStringSlice(readoutKeys(data), "asc")
	readouts := make([]map[string]interface{}, len(sortedReadouts))
	for readoutIndex, readout := range sortedReadouts {
		readouts[readoutIndex] = readoutFunc(readout, data[readout])
		for metricIndex, metricName := range metricOrder {
			segments[metricIndex].Values[readoutIndex] = helper.TruncateFloat(data[readout][metricName], 2)
		}
	}

	plot := typedef.CircHeatmapPlot{
		Name:     name,
		Readouts: readouts,
		Segments: segments,
	}

	return plot
}
