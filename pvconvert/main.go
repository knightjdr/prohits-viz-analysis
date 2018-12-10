// Package main takes file from ProHits-viz V1 and converts it to V2 JSON.
package main

import (
	"flag"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/filereader/columnparser"
)

func main() {
	// Get file to convert.
	file := flag.String("file", "", "ProHits-viz V1 file")
	flag.Parse()

	// Set file type and column mapping.
	fileType := []string{"text/tab-separated-values"}
	columnMap := map[string]string{
		"abundance": "column",
		"params":    "params",
		"row":       "row",
		"score":     "score",
		"value":     "value",
	}

	// Parse file and parameters.
	data := columnparser.ParseCsv([]string{*file}, fileType, columnMap)
	params := parseParams(data)
	fmt.Println(params)
}
