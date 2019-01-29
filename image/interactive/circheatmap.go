package interactive

import (
	"encoding/json"

	"github.com/knightjdr/prohits-viz-analysis/image/file"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// circHeatmapJSON stores the JSON structure for the interactive circular heatmap.
type circHeatmapJSON struct {
	AvailablePlots      []typedef.CircHeatmapPlot      `json:"availablePlots"`
	CircHeatmapSettings []typedef.CircHeatmapSetttings `json:"circHeatmapSettings"`
	Parameters          map[string]interface{}         `json:"parameters"`
	Settings            map[string]interface{}         `json:"settings"`
}

// CircHeatmap creates an interactive circular heatmap as json.
func CircHeatmap(
	plots []typedef.CircHeatmapPlot,
	parameters, settings map[string]interface{},
	segmentSettings []typedef.CircHeatmapSetttings,
	outfile string,
) {
	jsonStruct := circHeatmapJSON{
		AvailablePlots:      plots,
		CircHeatmapSettings: segmentSettings,
		Parameters:          parameters,
		Settings: map[string]interface{}{
			"current": settings,
		},
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
