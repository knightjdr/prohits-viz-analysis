package scv

import (
	"sort"

	"github.com/knightjdr/prohits-viz-analysis/pkg/slice"
	customSort "github.com/knightjdr/prohits-viz-analysis/pkg/sort"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createPlotData(data map[string]map[string]map[string]float64, known map[string]map[string]bool, legend types.CircHeatmapLegend, settings types.Settings) []types.CircHeatmap {
	conditions := getAndSortConditions(data)

	writeKnownness := getKnownessInteractiveWriter(settings.Known, known)

	plots := make([]types.CircHeatmap, len(conditions))
	for i, condition := range conditions {
		readouts := getAndSortReadouts(data[condition], known[condition], settings.Known, legend[0].Attribute)
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

	return plots
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

func getAndSortReadouts(data map[string]map[string]float64, knownMap map[string]bool, sortByKnown string, attribute string) []string {
	readouts := make([]string, 0)
	values := make([]float64, 0)

	if sortByKnown != "" && knownMap != nil {
		knownReadouts := make([]string, 0)
		knownValues := make([]float64, 0)
		unknownReadouts := make([]string, 0)
		unknownValues := make([]float64, 0)

		for readout, readoutData := range data {
			value := readoutData[attribute]
			if knownMap[readout] {
				knownReadouts = append(knownReadouts, readout)
				knownValues = append(knownValues, value)
			} else {
				unknownReadouts = append(unknownReadouts, readout)
				unknownValues = append(unknownValues, value)
			}
		}
		readouts = append(readouts, knownReadouts...)
		readouts = append(readouts, unknownReadouts...)
		values = append(values, knownValues...)
		values = append(values, unknownValues...)

		sorted := sortReadouts(knownValues, knownReadouts)
		return append(sorted, sortReadouts(unknownValues, unknownReadouts)...)
	}

	for readout, readoutData := range data {
		readouts = append(readouts, readout)
		values = append(values, readoutData[attribute])
	}

	return sortReadouts(values, readouts)
}

func sortReadouts(values []float64, readouts []string) []string {
	indices := customSort.ArgsortFloat(values)
	indices = slice.ReverseInt(indices)

	sorted := make([]string, 0)
	for _, index := range indices {
		sorted = append(sorted, readouts[index])
	}

	return sorted
}
