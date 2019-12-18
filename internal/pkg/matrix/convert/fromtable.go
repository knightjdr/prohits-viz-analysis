// Package convert has functions for convert to matrix format.
package convert

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// ConversionSettings stores parameters for converting from a table to a matrix.
type ConversionSettings struct {
	CalculateRatios bool
	Resort          bool
	ScoreType       string
}

// FromTable generates a 2D matrix of conditions v readouts (conditions = columns, readouts = rows).
// If the rows and columns should be sorted alphabetically, set resort to true. It generates matrices for
// both the abundance and score column, and can generate row ratios as well if requested.
// It also returns lists of the conditions and readouts.
func FromTable(table *[]map[string]string, settings ConversionSettings) *types.Matrices {
	data := parseTable(table, settings.ScoreType)

	matrices := &types.Matrices{
		Conditions: sortLabels(data.conditions, settings.Resort),
		Readouts:   sortLabels(data.readouts, settings.Resort),
	}
	createMatrices(matrices, data, settings)

	return matrices
}
