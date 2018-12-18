// Package main takes a file from ProHits-viz V1 and converts it to V2 JSON.
package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/filereader/columnparser"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/image/interactive"
	"github.com/knightjdr/prohits-viz-analysis/image/minimap"
)

func main() {
	// Get file to convert.
	file := flag.String("file", "", "ProHits-viz V1 file")
	flag.Parse()

	// Set file type and column mapping.
	fileType := []string{"text/tab-separated-values"}
	columnMap := map[string]string{
		"condition": "column",
		"params":    "params",
		"readout":   "row",
		"score":     "score",
		"abundance": "value",
	}

	// Parse file and parameters.
	csv := columnparser.ParseCsv([]string{*file}, fileType, columnMap, true)
	imageType, parameters := parseParams(csv)

	// Convert csv to matrices
	matrices := helper.ConditionReadoutMatrix(&csv, parameters.ScoreType, false, true)

	// Create folders
	helper.CreateFolders([]string{"interactive", "minimap"})

	// Generate minimap
	mapData := minimap.Data{
		ImageType:  imageType,
		Matrices:   matrices,
		Parameters: parameters,
		Path:       "minimap",
	}
	minimap.Write(&mapData)

	// Generate interactive file
	interactiveData := interactive.Data{
		ImageType:  imageType,
		Matrices:   matrices,
		Minimap:    "minimap/minimap.png",
		Parameters: parameters,
		Path:       fmt.Sprintf("interactive/%s.json", imageType),
	}
	interactive.ParseHeatmap(&interactiveData)

	// Remove minimap folder.
	mapFolder := filepath.Join(".", "minimap")
	fs.Instance.RemoveAll(mapFolder)
}
