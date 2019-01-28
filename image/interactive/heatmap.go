// Package interactive generates interactive files for prohits viz.
package interactive

import (
	"encoding/json"

	"github.com/knightjdr/prohits-viz-analysis/image/file"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// HeatmapJSON stores the JSON structure for the interactive heatmap.
type HeatmapJSON struct {
	Columns    columnObj              `json:"columns"`
	Parameters map[string]interface{} `json:"parameters"`
	Settings   map[string]interface{} `json:"settings"`
	Rows       rowObj                 `json:"rows"`
	Minimap    mapObj                 `json:"minimap"`
}

type columnObj struct {
	Names []string `json:"names"`
	Ref   *string  `json:"ref"`
}

type mapObj struct {
	Image string `json:"image"`
}

type rowObj struct {
	List []map[string]interface{} `json:"list"`
}

// Heatmap creates an interactive heatmap as json. The data matrix, row and column
// names should already be sorted.
func Heatmap(
	columns []string,
	rowdata []map[string]interface{},
	parameters map[string]interface{},
	settings map[string]interface{},
	uri, outfile string,
) {
	var jsonStruct HeatmapJSON

	jsonStruct.Columns = columnObj{
		Names: columns,
	}
	jsonStruct.Parameters = parameters
	jsonStruct.Rows = rowObj{
		List: rowdata,
	}
	jsonStruct.Settings = map[string]interface{}{
		"current": settings,
	}
	jsonStruct.Minimap = mapObj{
		Image: uri,
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
