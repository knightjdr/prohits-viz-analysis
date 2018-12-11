// Package main takes a file from ProHits-viz V1 and converts it to V2 JSON.
package main

import (
	"flag"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/filereader/columnparser"
	"github.com/knightjdr/prohits-viz-analysis/helper"
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
		"value":     "value",
	}

	// Parse file and parameters.
	csv := columnparser.ParseCsv([]string{*file}, fileType, columnMap)
	imageType, parameters := parseParams(csv)

	// Convert csv to matrices
	matrices := helper.ConditionReadoutMatrix(csv, parameters.ScoreType, false)

	// Generate minimap
	writeMinimap()

	fmt.Println(matrices, imageType)
}
