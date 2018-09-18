package main

import (
	"github.com/knightjdr/prohits-viz-analysis/parse"
)

// RowNames taks an array or Rows and returns just the names.
func RowNames(rows []parse.Row) []string {
	rowNum := len(rows)
	rowNames := make([]string, rowNum)

	// Define matrices.
	for i, row := range rows {
		rowNames[i] = row.Name
	}
	return rowNames
}
