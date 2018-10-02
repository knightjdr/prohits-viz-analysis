// Package interactive generates interactive files for prohits viz.
package interactive

import (
	"encoding/json"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// HeatmapJSON stores the JSON structure for the interactive heatmap.
type HeatmapJSON struct {
	Columns    []string                 `json:"columns"`
	Parameters map[string]interface{}   `json:"parameters"`
	Settings   map[string]interface{}   `json:"settings"`
	Rows       []map[string]interface{} `json:"rows"`
	URI        string                   `json:"minimap"`
}

// Heatmap creates an interactive heatmap as json. The data matrix, row and column
// names should already be sorted. UserParams has the parameters used for the
// analysis.
func Heatmap(
	data []map[string]interface{},
	columns []string,
	parameters map[string]interface{},
	settings map[string]interface{},
	uri string,
) (jsonString string) {
	var jsonStruct HeatmapJSON

	jsonStruct.Columns = columns
	jsonStruct.Parameters = parameters
	jsonStruct.Rows = data
	jsonStruct.Settings = settings
	jsonStruct.URI = uri

	// Convert struct to json.
	byte, err := json.Marshal(jsonStruct)
	// Log message but do not panic.
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}
	jsonString = string(byte)
	return
}
