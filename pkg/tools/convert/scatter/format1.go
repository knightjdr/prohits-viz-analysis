package scatter

import (
	"io"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func readFormat1(filename string) ([]types.ScatterPlot, types.Settings, []map[string]string) {
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

		if line[0] == "entry" {
			if plot.Name != "" {
				plots = append(plots, plot)
			}

			plotSettings.Plot = line[4]
			plotSettings.Score = line[5]
			plotSettings.XLabel = line[2]
			plotSettings.YLabel = line[3]

			plot = types.ScatterPlot{
				Labels: types.ScatterAxesLabels{
					X: plotSettings.XLabel,
					Y: plotSettings.YLabel,
				},
				Name:   line[1],
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

	legend := defineLegend(plotSettings, 1)
	settings := inferSettings(plotSettings)

	return plots, settings, legend
}
