// Package interactive generates interactive files for prohits viz.
package interactive

import (
	"encoding/json"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// HeatmapJson stores the JSON structure for the interactive heatmap.
type HeatmapJson struct {
	Columns []string                 `json:"columns"`
	Params  map[string]interface{}   `json:"params"`
	Rows    []map[string]interface{} `json:"rows"`
	Uri     string                   `json:"minimap"`
}

// Heatmap creates an interactive heatmap as json. The data matrix, row and column
// names should already be sorted. UserParams has the parameters used for the
// analysis.
func Heatmap(
	data []map[string]interface{},
	columns []string,
	params map[string]interface{},
	uri string,
) (jsonString string) {
	var jsonStruct HeatmapJson

	// Add header.
	jsonStruct.Columns = columns

	// Set Params.
	jsonStruct.Params = params

	// Add rows.
	jsonStruct.Rows = data

	// Add uri.
	jsonStruct.Uri = uri

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
