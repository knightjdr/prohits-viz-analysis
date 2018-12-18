package svg

import (
	"math"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// HeatmapDimensions contains heatmap plot dimensions.
type HeatmapDimensions struct {
	cellSize   int
	fontSize   int
	leftMargin int
	plotHeight int
	plotWidth  int
	ratio      float64
	svgHeight  int
	svgWidth   int
	topMargin  int
}

// Dimensions calculates the dimensions to use for a heatmap.
func Dimensions(matrices *typedef.Matrices, minimap bool) (dims HeatmapDimensions) {
	// Check row and column size and adjust plot parameters as needed. The parameter
	// adjustment is done based on whichever dimension exceeds the limits by
	// a greater amount.
	colNum := len(matrices.Abundance[0])
	colRatio := float64(colNum*idealCellSize) / float64(maxImageWidth)
	rowNum := len(matrices.Abundance)
	rowRatio := float64(rowNum*idealCellSize) / float64(maxImageHeight)

	// Set parameters based on ratios. If there are more columns or rows than would
	// fit with the ideal cell size, get the ratio to adjust down by.
	dims.ratio = float64(1)
	if colRatio > 1 || rowRatio > 1 {
		dims.ratio = math.Max(colRatio, rowRatio)
		dims.ratio = 1 / dims.ratio
	}
	if dims.ratio < minRatio {
		dims.ratio = minRatio
	}
	dims.cellSize = int(math.Floor(dims.ratio * float64(idealCellSize)))

	// For minimaps, return plot area with no margin. Otherwise calculate
	// margins needed for column and row labels.
	if minimap {
		dims.topMargin = 0
		dims.leftMargin = 0
		dims.svgHeight = rowNum * dims.cellSize
		dims.svgWidth = colNum * dims.cellSize
		dims.plotHeight = dims.svgHeight
		dims.plotWidth = dims.svgWidth
	} else {
		dims.fontSize = int(math.Floor(dims.ratio * float64(idealFontSize)))

		// Calculate required top margin. Find the longest column name and assume it
		// is made entirely of the "W" character (which has a width of 11.33px
		// in arial with a 12pt fontsize).
		longestColumnName := 0
		for _, colName := range matrices.Conditions {
			nameLength := len([]rune(colName))
			if nameLength > longestColumnName {
				longestColumnName = nameLength
			}
		}
		dims.topMargin = int(math.Round(float64(longestColumnName) * 11.33 * dims.ratio))

		// Calculate required left margin. Find the longest row name and assume it
		// is made entirely of the "W" character (which has a width of 11.33px at
		// in arial with a 12pt fontsize).
		longestRowName := 0
		for _, rowName := range matrices.Readouts {
			nameLength := len([]rune(rowName))
			if nameLength > longestRowName {
				longestRowName = nameLength
			}
		}
		dims.leftMargin = int(math.Round(float64(longestRowName) * 11.33 * dims.ratio))

		// Set plot dimensions.
		dims.svgHeight = dims.topMargin + (rowNum * dims.cellSize)
		dims.svgWidth = dims.leftMargin + (colNum * dims.cellSize)
		dims.plotHeight = dims.svgHeight - dims.topMargin
		dims.plotWidth = dims.svgWidth - dims.leftMargin
	}

	return
}
