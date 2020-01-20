package heatmap

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/knightjdr/prohits-viz-analysis/pkg/mapf"
)

func parseSettings(settings map[string]interface{}) string {
	imageSettings := map[string]map[string]map[string]interface{}{
		"main": map[string]map[string]interface{}{
			"current": map[string]interface{}{},
		},
	}

	keys := mapf.KeysStringInterface(settings)
	sort.Strings(keys)
	for _, key := range keys {
		imageSettings["main"]["current"][key] = settings[key]
	}

	jsonString, _ := json.Marshal(imageSettings)
	return fmt.Sprintf("\"settings\": %s", string(jsonString))
}
