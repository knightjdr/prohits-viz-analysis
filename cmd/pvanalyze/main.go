package main

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/analyze"
)

func main() {
	analyze.Run()
	/* // Parse flags.
	columnMap, parameters, err := parseFlags()
	if err != nil {
		os.Exit(1)
	}

	// Create dataset.
	dataset := typedef.Dataset{Parameters: parameters}

	// Read needed columns from files.
	dataset.FileData = columnparser.ReadFile(parameters.Files, columnMap, false)

	// Filter rows.
	dataset.FileData = filter.Data(dataset.FileData, dataset.Parameters)

	// Check for common errors in filtered data that result from incorrect input format.
	err = errorcheck.Required(dataset)
	if err != nil {
		os.Exit(1)
	}

	// Transform readout abundances.
	dataset.FileData = transform.Readouts(dataset)

	// Perform analysis
	tool.Start(&dataset) */
}
