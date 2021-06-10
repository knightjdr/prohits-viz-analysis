package scatter

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func readFormat2(filename string) ([]types.ScatterPlot, types.Settings, []map[string]string) {
	file, err := fs.Instance.Open(filename)
	log.CheckError(err, true)
	defer file.Close()
	reader := createReader(file)

	plots := make([]types.ScatterPlot, 0)

	var plot types.ScatterPlot
	plotSettings := &jsonSettings{}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		log.CheckError(err, true)

		if line[0] == "details:" {
			if plot.Name != "" {
				plots = append(plots, plot)
			}
			parseJSONSettings(plotSettings, line[1])

			plot = types.ScatterPlot{
				Labels: types.ScatterAxesLabels{
					X: plotSettings.XLabel,
					Y: plotSettings.YLabel,
				},
				Name:   plotSettings.Bait,
				Points: make([]types.ScatterPoint, 0),
			}
		} else {
			x, _ := strconv.ParseFloat(line[1], 64)
			y, _ := strconv.ParseFloat(line[2], 64)
			point := types.ScatterPoint{
				Color: line[4],
				Label: line[0],
				X:     x,
				Y:     y,
			}
			plot.Points = append(plot.Points, point)
		}
	}
	plots = append(plots, plot)

	legend := defineLegend(plotSettings, 2)
	settings := inferSettings(plotSettings)

	return plots, settings, legend
}

func parseJSONSettings(settings *jsonSettings, jsonString string) {
	err := json.Unmarshal([]byte(jsonString), settings)
	log.CheckError(err, false)
}
