// Package main takes input files and Parameters and runs prohits-viz analysis
package main

import (
	"os"

	"github.com/knightjdr/prohits-viz-analysis/errorcheck"
	"github.com/knightjdr/prohits-viz-analysis/filereader/columnparser"
	"github.com/knightjdr/prohits-viz-analysis/filereader/filter"
	"github.com/knightjdr/prohits-viz-analysis/tool"
	"github.com/knightjdr/prohits-viz-analysis/transform"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

func main() {
	// Parse flags.
	columnMap, parameters, err := ParseFlags()
	if err != nil {
		os.Exit(1)
	}

	// Create dataset.
	dataset := typedef.Dataset{Parameters: parameters}

	// Read needed columns from files.
	dataset.FileData = columnparser.ReadFile(parameters.Files, columnMap, false)

	// Filter rows.
	dataset.FileData = filter.Data(parsedColumns, dataset.Parameters)

	// Check for common errors in filtered data that result from incorrect input format.
	err = errorcheck.Required(dataset)
	if err != nil {
		os.Exit(1)
	}

	// Transform readout abundances.
	dataset.FileData = transform.Readouts(dataset)

	// Perform analysis
	tool.Start(&dataset)
}
