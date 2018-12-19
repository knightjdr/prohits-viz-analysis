// Package main takes a stringified JSON and generates a minimap.
package main

import (
	"flag"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/image/svg"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/parse"
)

func main() {
	// Parse flags.
	jsonFile := flag.String("json", "", "JSON file")
	outputType := flag.String("type", "", "Type of file export")
	flag.Parse()

	data, err := parse.HeatmapJSON(*jsonFile)
	logmessage.CheckError(err, true)

	// Format dataset for svg creator.
	matrices := parse.FormatMatrix(data)

	// Format parameters and generate readout names for svg.
	parameters := FormatParams(data)
	matrices.Readouts = RowNames(data.Rows)

	// Create folders
	folders := make([]string, 2)
	folders[0] = "svg"
	if *outputType != "svg" {
		folders[1] = *outputType
	}
	helper.CreateFolders(folders)

	// Create svg.
	svgData := svg.Data{
		Annotations: data.Annotations,
		Filename:    fmt.Sprintf("svg/%s.svg", data.ImageType),
		ImageType:   data.ImageType,
		Markers:     data.Markers,
		Matrices:    matrices,
		Minimap:     false,
		Parameters:  parameters,
	}
	svg.Heatmap(&svgData)

	// Create additional output type if needed.
	imageName := fmt.Sprintf("%s.svg", data.ImageType)
	if *outputType == "pdf" {
		svg.ConvertPdf([]string{imageName})
	} else if *outputType == "png" {
		svg.ConvertPng([]string{imageName})
	}
}
