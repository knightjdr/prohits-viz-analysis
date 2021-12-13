package scv

import (
	"encoding/json"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/read/csv"
	"github.com/knightjdr/prohits-viz-analysis/pkg/slice"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

func defineKnown(data map[string]map[string]map[string]float64, idMaps map[string]map[string]string, settings types.Settings) map[string]map[string]bool {
	if settings.Known == "interaction" {
		knownness := readKnownInteractionFromJSON(settings.KnownFile)
		return defineKnownInteraction(data, idMaps, knownness)
	}
	if settings.Known == "custom" && settings.KnownFile != "" {
		knownness, updatedIDMap := readKnownInteractionFromTXT(settings.KnownFile, idMaps)
		return defineKnownInteraction(data, updatedIDMap, knownness)
	}

	return nil
}

func readKnownInteractionFromJSON(filename string) map[string][]string {
	file, err := fs.Instance.Open(filename)
	log.CheckError(err, true)

	bytes, err := afero.ReadAll(file)
	log.CheckError(err, true)

	var interactions map[string][]string
	json.Unmarshal(bytes, &interactions)

	return interactions
}

func readKnownInteractionFromTXT(filename string, idMaps map[string]map[string]string) (map[string][]string, map[string]map[string]string) {
	interactions := csv.ReadToSliceMap(filename, '\t')

	updatedIDMap := make(map[string]map[string]string)
	for key, mapping := range idMaps {
		updatedIDMap[key] = make(map[string]string)
		for symbol := range mapping {
			updatedIDMap[key][symbol] = symbol
		}
	}

	return interactions, updatedIDMap
}

func defineKnownInteraction(data map[string]map[string]map[string]float64, idMaps map[string]map[string]string, interactions map[string][]string) map[string]map[string]bool {
	known := make(map[string]map[string]bool)
	for condition, conditionData := range data {
		known[condition] = make(map[string]bool)

		mappedCondition := idMaps["condition"][condition]
		_, conditionHasInteractionData := interactions[mappedCondition]

		lookup := make(map[string]bool)
		if conditionHasInteractionData {
			lookup = slice.ConvertToBoolMap(interactions[mappedCondition])
		}

		for readout := range conditionData {
			mappedReadout := idMaps["readout"][readout]
			_, interactionIsKnown := lookup[mappedReadout]
			if conditionHasInteractionData && interactionIsKnown {
				known[condition][readout] = true
			} else {
				known[condition][readout] = false
			}
		}
	}

	return known
}
