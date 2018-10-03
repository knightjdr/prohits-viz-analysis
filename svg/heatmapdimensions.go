package svg

import (
	"math"
)

// HDimensions contains heatmap plot dimensions.
type HDimensions struct {
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

// HeatmapDimensions calculates the dimensions to use for the heatmap.
func HeatmapDimensions(matrix [][]float64, columns, rows []string) (dims HDimensions) {
	// Check row and column size and adjust plot parameters as needed. The parameter
	// adjustment is done based on whichever dimension exceeds the limits by
	// a greater amount.
	colNum := len(matrix[0])
	colRatio := float64(colNum*idealCellSize) / float64(maxImageWidth)
	rowNum := len(matrix)
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
	dims.fontSize = int(math.Floor(dims.ratio * float64(idealFontSize)))

	// Calculate required top margin. Find the longest column name and assume it
	// is made entirely of the "W" character (which has a width of 11.33px
	// in arial with a 12pt fontsize).
	longestColumnName := 0
	for _, colName := range columns {
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
	for _, rowName := range rows {
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

	return
}
