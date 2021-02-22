package scv

import (
	"encoding/json"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/slice"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

func defineKnown(data map[string]map[string]map[string]float64, idMaps map[string]map[string]string, settings types.Settings) map[string]map[string]bool {
	if settings.Known == "interaction" {
		return defineKnownInteraction(data, idMaps, settings.KnownFile)
	}

	return nil
}

func defineKnownInteraction(data map[string]map[string]map[string]float64, idMaps map[string]map[string]string, filename string) map[string]map[string]bool {
	file, err := fs.Instance.Open(filename)
	log.CheckError(err, true)

	bytes, err := afero.ReadAll(file)
	log.CheckError(err, true)

	var interactions map[string][]string
	json.Unmarshal(bytes, &interactions)

	known := make(map[string]map[string]bool, 0)
	for condition, conditionData := range data {
		known[condition] = make(map[string]bool, 0)

		mappedCondition := idMaps["condition"][condition]
		_, conditionHasInteractionData := interactions[mappedCondition]

		lookup := make(map[string]bool, 0)
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
