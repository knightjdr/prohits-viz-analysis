// Package main takes a stringified JSON and generates a minimap.
package main

import (
	"flag"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/image/interactive"
	"github.com/knightjdr/prohits-viz-analysis/image/minimap"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/parse"
)

func main() {
	// Parse flags.
	jsonFile := flag.String("json", "", "JSON file")
	flag.Parse()

	data, err := parse.HeatmapJSON(*jsonFile)
	logmessage.CheckError(err, true)

	// Format dataset for svg creator.
	matrices := parse.FormatMatrix(data)

	// Format parameters for svg.
	parameters := FormatParams(data)

	// Creat dummy row and column names.
	matrices.Conditions, matrices.Readouts = Dummy(len(data.Rows[0].Data), len(data.Rows))

	// Create folders
	helper.CreateFolders([]string{"minimap"})

	// Create minimap.
	mapData := minimap.Data{
		Filename:   "minimap/minimap",
		ImageType:  data.ImageType,
		Matrices:   matrices,
		Parameters: parameters,
	}
	minimap.Write(&mapData)

	// Generate URL.
	url := interactive.Pngurl("minimap/minimap.png")
	fmt.Println(url)
}
