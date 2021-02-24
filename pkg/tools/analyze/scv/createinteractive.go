package scv

import (
	"sort"

	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createInteractive(data map[string]map[string]map[string]float64, known map[string]map[string]bool, legend types.CircHeatmapLegend, settings types.Settings) {
	conditions := getAndSortConditions(data)

	writeKnownness := getKnownessInteractiveWriter(settings.Known, known)

	plots := make([]types.CircHeatmap, len(conditions))
	for i, condition := range conditions {
		readouts := getAndSortReadouts(data[condition])
		plots[i] = types.CircHeatmap{
			Name:     condition,
			Readouts: make([]types.CircHeatmapReadout, len(readouts)),
		}
		for j, readout := range readouts {
			plots[i].Readouts[j] = types.CircHeatmapReadout{
				Known:    writeKnownness(condition, readout),
				Label:    readout,
				Segments: make(map[string]types.RoundedSegment, len(legend)),
			}
			for _, legendElement := range legend {
				plots[i].Readouts[j].Segments[legendElement.Attribute] = types.RoundedSegment(data[condition][readout][legendElement.Attribute])
			}
		}
	}

	interactiveData := &interactive.CircHeatmapData{
		Filename:   "interactive/scv.json",
		Legend:     legend,
		Parameters: settings,
		Plots:      plots,
		Settings: map[string]interface{}{
			"sortByKnown": shownKnownElement(settings.Known),
		},
	}
	interactive.CreateCircHeatmap(interactiveData)
}

func getAndSortConditions(data map[string]map[string]map[string]float64) []string {
	conditions := make([]string, len(data))
	i := 0
	for condition := range data {
		conditions[i] = condition
		i++
	}
	sort.Strings(conditions)
	return conditions
}

func getKnownessInteractiveWriter(known string, knownMap map[string]map[string]bool) func(string, string) bool {
	if known != "" && knownMap != nil {
		return func(condition string, readout string) bool {
			return knownMap[condition][readout]
		}
	}
	return func(condition string, readout string) bool {
		return false
	}
}

func getAndSortReadouts(data map[string]map[string]float64) []string {
	readouts := make([]string, len(data))
	i := 0
	for readout := range data {
		readouts[i] = readout
		i++
	}
	sort.Strings(readouts)
	return readouts
}

func shownKnownElement(known string) bool {
	if known != "" {
		return true
	}
	return false
}
