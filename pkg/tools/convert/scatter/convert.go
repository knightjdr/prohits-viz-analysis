// Package scatter converts an interactive file from ProHits-viz V1 to V2 format
package scatter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

type jsonSettings struct {
	Bait            string  `json:"bait,omitempty"`
	Plot            string  `json:"tool,omitempty"`
	PrimaryFilter   float64 `json:"primary,omitempty"`
	Score           string  `json:"score,omitempty"`
	ScoreType       int     `json:"filter,omitempty"`
	SecondaryFilter float64 `json:"secondary,omitempty"`
	XLabel          string  `json:"xAxis,omitempty"`
	YLabel          string  `json:"yAxis,omitempty"`
}

// Convert a heatmap or dotplot file to json format.
func Convert(filename string) {
	format := determineFormat(filename)

	files.CreateFolders([]string{"interactive"})
	fileid := strings.Split(filename, ".txt")[0]

	legend := []map[string]string{}
	plots := []types.ScatterPlot{}
	settings := types.Settings{}
	if format == 1 {
		plots, settings, legend = readFormat1(filename)
	}
	if format == 2 {
		plots, settings, legend = readFormat2(filename)
	}

	createInteractive(plots, settings, legend, fileid)
}

func determineFormat(filename string) int {
	file, err := fs.Instance.Open(filename)
	log.CheckError(err, true)
	defer file.Close()
	reader := createReader(file)

	header, err := reader.Read()
	log.CheckError(err, true)

	var format int
	if header[0] == "entry" {
		format = 1
	}
	if header[0] == "details:" {
		format = 2
	}

	return format
}

func createReader(file afero.File) *csv.Reader {
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	return reader
}

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

func determineAnalysisType(plotType string) string {
	if plotType == "Binary" {
		return "condition-condition"
	}
	if plotType == "Binary-biDist" {
		return "condition-condition"
	}
	if plotType == "Specificity" {
		return "specificity"
	}
	return ""
}

func defineLegend(plotSettings *jsonSettings, format int) []map[string]string {
	if plotSettings.Plot == "Binary" {
		return []map[string]string{
			{
				"color": "#ff0000",
				"text":  "Infinite fold change",
			},
		}
	}
	if plotSettings.Plot == "Binary-biDist" {
		scoreSymbols := createLegendScoreSymbol(plotSettings.ScoreType)
		return []map[string]string{
			{
				"color": "#0066cc",
				"text": fmt.Sprintf(
					"%s %s %s",
					plotSettings.Score,
					scoreSymbols[0],
					float.RemoveTrailingZeros(plotSettings.PrimaryFilter),
				),
			},
			{
				"color": "#99ccff",
				"text": fmt.Sprintf(
					"%s %s %s %s %s",
					float.RemoveTrailingZeros(plotSettings.PrimaryFilter),
					scoreSymbols[1],
					plotSettings.Score,
					scoreSymbols[2],
					float.RemoveTrailingZeros(plotSettings.SecondaryFilter),
				),
			},
		}
	}
	if plotSettings.Plot == "Specificity" && format == 1 {
		return []map[string]string{
			{
				"color": "#ff0000",
				"text":  "Infinite specificity",
			},
		}
	}
	if plotSettings.Plot == "Specificity" && format == 2 {
		return []map[string]string{
			{
				"color": "#0066cc",
				"text":  "Infinite specificity",
			},
		}
	}
	return []map[string]string{}
}

func parseJSONSettings(settings *jsonSettings, jsonString string) {
	err := json.Unmarshal([]byte(jsonString), settings)
	log.CheckError(err, false)
}

func createLegendScoreSymbol(scoreType int) []string {
	if scoreType == 1 {
		return []string{"≥", ">", "≥"}
	}
	return []string{"≤", "<", "≤"}
}

func inferSettings(plotSettings *jsonSettings) types.Settings {
	settings := types.Settings{
		Score: plotSettings.Score,
		Type:  determineAnalysisType(plotSettings.Plot),
	}

	if plotSettings.Plot == "Binary-biDist" {
		settings.ScoreType = convertScoreTypeFromInt(plotSettings.ScoreType)
	}
	if plotSettings.Plot == "Specificity" {
		settings.Abundance = plotSettings.XLabel
	}

	return settings
}

func convertScoreTypeFromInt(scoreType int) string {
	if scoreType == 1 {
		return "gte"
	}
	return "lte"
}

func createInteractive(plots []types.ScatterPlot, settings types.Settings, legend []map[string]string, filename string) {
	interactiveData := &interactive.ScatterData{
		AnalysisType: settings.Type,
		Filename:     fmt.Sprintf("interactive/%s.json", filename),
		Legend:       legend,
		Parameters:   settings,
		Plots:        plots,
		Settings: map[string]interface{}{
			"xFilter": 0,
			"yFilter": 0,
		},
	}

	interactive.CreateScatter(interactiveData)
}
