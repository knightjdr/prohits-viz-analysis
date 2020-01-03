// Package heatmap exports images in png or svg format.
package heatmap

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// Export image.
func Export(filename, format string) {
	data := readJSON(filename)

	matrix := createMatrix(data)
	columns, rows := getColumnsAndRows(data)

	matrices := &types.Matrices{
		Abundance:  matrix,
		Conditions: columns,
		Readouts:   rows,
	}
	fmt.Println(matrices)
}
