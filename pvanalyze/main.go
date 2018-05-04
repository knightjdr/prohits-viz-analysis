// Package main takes input files and Parameters and runs prohits-viz analysis
package main

import (
	"os"

	"github.com/knightjdr/prohits-viz-analysis/errorcheck"
	"github.com/knightjdr/prohits-viz-analysis/filereader/columnparser"
	"github.com/knightjdr/prohits-viz-analysis/filereader/filter"
	"github.com/knightjdr/prohits-viz-analysis/tool"
	"github.com/knightjdr/prohits-viz-analysis/transform"
	"github.com/knightjdr/prohits-viz-analysis/types"
)

func main() {
	// Parse flags.
	columnMap, params, err := ParseFlags()
	if err != nil {
		os.Exit(1)
	}

	// Create dataset.
	dataset := types.Dataset{Params: params}

	// Read needed columns from files.
	parsedColumns := columnparser.ReadFile(params.Files, columnMap)

	// Filter rows.
	dataset.Data = filter.Data(parsedColumns, dataset.Params)

	// Check for common errors in filtered data that result from incorrect input format.
	err = errorcheck.Required(dataset)
	if err != nil {
		os.Exit(1)
	}

	// Transform prey abundances.
	dataset.Data = transform.Preys(dataset)

	// Perform analysis
	tool.Start(dataset)
}