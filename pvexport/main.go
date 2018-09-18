// Package main takes a stringified JSON and generates a minimap.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/knightjdr/prohits-viz-analysis/parse"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/tool/dotplot"
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
	params := FormatParams(data)

	rowNames := RowNames(data.Rows)

	// Create svg.
	image := "heatmap"
	if data.ImageType == "dotplot" {
		image = "dotplot"
		dotplot.SvgDotplot(
			abundance,
			ratios,
			scores,
			data.Columns,
			rowNames,
			data.Invert,
			params,
		)
	} else {
		dotplot.SvgHeatmap(
			abundance,
			data.Columns,
			rowNames,
			data.FillColor,
			data.MaximumAbundance,
			data.Invert,
		)
	}

	// Create additional output type if needed.
	imageName := fmt.Sprintf("%s.svg", image)
	if *outputType == "pdf" {
		svg.ConvertPdf([]string{imageName})
	} else if *outputType == "png" {
		svg.ConvertPng([]string{imageName})
	}
}
