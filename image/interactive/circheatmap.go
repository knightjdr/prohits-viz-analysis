package interactive

import (
	"encoding/json"

	"github.com/knightjdr/prohits-viz-analysis/helper"

	"github.com/knightjdr/prohits-viz-analysis/image/file"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// circHeatmapJSON stores the JSON structure for the interactive circular heatmap.
type circHeatmapJSON struct {
	AvailablePlots      []circularPlot           `json:"availablePlots"`
	CircHeatmapSettings []map[string]interface{} `json:"circHeatmapSettings"`
	Settings            map[string]interface{}   `json:"settings"`
}

// circularPlot describes an individual circular heatmap
type circularPlot struct {
	Name     string                   `json:"name"`
	Readouts []map[string]interface{} `json:"readouts"`
	Segments []cicularSegments        `json:"segments"`
}

// cicularSegments describes a segment on a circular heatmap
type cicularSegments struct {
	Name   string    `json:"name"`
	Values []float64 `json:"values"`
}

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

// CircHeatmap creates an interactive circular heatmap as json. The data matrix, row and column
// names should already be sorted
func CircHeatmap(
	conditionOrder []string,
	data map[string]map[string]map[string]float64,
	settings map[string]interface{},
	readoutMetrics map[string]string,
	outfile string,
) {
	readoutFunc := formatReadout(settings["known"].(bool))
	var jsonStruct circHeatmapJSON

	// Convert condition data to plot format.
	availablePlots := make([]circularPlot, len(conditionOrder))
	for index, condition := range conditionOrder {
		sortedReadouts := helper.SortStringSlice(readoutKeys(data[condition]), "asc")
		readouts := make([]map[string]interface{}, len(sortedReadouts))
		segments := make([]cicularSegments, len(readoutMetrics))
		segmentNum := 0
		for _, metricName := range readoutMetrics {
			segments[segmentNum] = cicularSegments{
				Name:   metricName,
				Values: make([]float64, len(sortedReadouts)),
			}
			segmentNum++
		}
		for readoutIndex, readout := range sortedReadouts {
			readouts[readoutIndex] = readoutFunc(readout, data[condition][readout])
			segmentNum := 0
			for metric := range readoutMetrics {
				segments[segmentNum].Values[readoutIndex] = data[condition][readout][metric]
				segmentNum++
			}
		}
		availablePlots[index] = circularPlot{
			Name:     condition,
			Readouts: readouts,
			Segments: segments,
		}
	}
	jsonStruct.AvailablePlots = availablePlots
	jsonStruct.Settings = map[string]interface{}{
		"current": settings,
	}

	// Open file for writing
	file, err := file.Create(outfile)
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}

	// Stream json to file.
	enc := json.NewEncoder(file)
	enc.SetIndent("", "\t")
	if err := enc.Encode(jsonStruct); err != nil {
		logmessage.CheckError(err, false)
		return
	}
}
