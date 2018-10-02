// Package main takes a stringified JSON and generates a minimap.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/knightjdr/prohits-viz-analysis/parse"
	"github.com/knightjdr/prohits-viz-analysis/svg"
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
	abundance, ratios, scores := parse.FormatMatrix(data)

	// Format parameters for svg.
	parameters := FormatParams(data)

	rowNames := RowNames(data.Rows)

	// Create svg.
	Export(
		data.ImageType,
		abundance,
		ratios,
		scores,
		data.Annotations,
		data.Markers,
		data.Columns,
		rowNames,
		parameters,
	)

	// Create additional output type if needed.
	imageName := fmt.Sprintf("%s.svg", data.ImageType)
	if *outputType == "pdf" {
		svg.ConvertPdf([]string{imageName})
	} else if *outputType == "png" {
		svg.ConvertPng([]string{imageName})
	}
}
