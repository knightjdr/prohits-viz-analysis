// Package main takes a stringified JSON and generates a minimap.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/parse"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/spf13/afero"
)

func main() {
	// Parse flags.
	jsonFile := flag.String("json", "", "JSON file")
	outputType := flag.String("type", "", "Type of file export")
	flag.Parse()

	data, err := parse.HeatmapJSON(*jsonFile)
	if err != nil {
		os.Exit(1)
	}

	// Format dataset for svg creator.
	heatmap := parse.FormatMatrix(data)

	// Format parameters and generate for names for svg.
	parameters := FormatParams(data)
	rowNames := RowNames(data.Rows)

	// Create svg.
	content := svg.Heatmap(
		data.ImageType,
		heatmap,
		data.Annotations,
		data.Markers,
		data.Columns,
		rowNames,
		false,
		parameters,
	)
	filename := fmt.Sprintf("svg/%s.svg", data.ImageType)
	afero.WriteFile(fs.Instance, filename, []byte(content), 0644)

	// Create additional output type if needed.
	imageName := fmt.Sprintf("%s.svg", data.ImageType)
	if *outputType == "pdf" {
		svg.ConvertPdf([]string{imageName})
	} else if *outputType == "png" {
		svg.ConvertPng([]string{imageName})
	}
}
