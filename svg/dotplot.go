// Package svg creates svg files for various image types.
package svg

import (
	"bytes"
	"fmt"
	"math"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// ScoreColorFunc returns a function for determining the gradient index to use
// for the score color.
func ScoreColorFunc(scoretype string, primary, secondary float64, numColors int) func(score float64) int {
	if scoretype == "gte" {
		return func(score float64) int {
			if score >= primary {
				return numColors
			} else if score < primary && score >= secondary {
				return numColors / 2
			}
			return numColors / 4
		}
	} else {
		return func(score float64) int {
			if score <= primary {
				return numColors
			} else if score > primary && score <= secondary {
				return numColors / 2
			}
			return numColors / 4
		}
	}
}

// Dotplot creates a dotplot from an input matrices of abundance, abundance
// ratios and score.
func Dotplot(
	matrix, ratios, scores [][]float64,
	columns, rows []string,
	options map[string]interface{},
) (svg string) {
	svgSlice := make([]string, 0)

	// Check row and column size and adjust plot params as needed. The parameter
	// adjustment is done based on whichever dimension exceeds the limits by
	// a greater amount.
	colNum := len(matrix[0])
	colRatio := float64(colNum*idealCellSize) / float64(maxImageWidth)
	rowNum := len(matrix)
	rowRatio := float64(rowNum*idealCellSize) / float64(maxImageHeight)

	// Set parameters based on ratios. If there are more columns or rows than would
	// fit with the ideal cell size, get the ratio to adjust down by.
	ratio := float64(1)
	if colRatio > 1 || rowRatio > 1 {
		ratio = math.Max(colRatio, rowRatio)
		ratio = 1 / ratio
	}

	// Sit the size of each cell on the dotplot.
	cellSize := int(math.Floor(ratio * float64(idealCellSize)))
	cellSizeHalf := int(math.Round(float64(cellSize) / float64(2)))

	// Determine the amount to subtract for the maximum circle radius. We do this
	// so that there is some space between circles. For the ideal cell size of 20
	// (half size 10), I subtract 1 pixel from the radius to create 2 pixels of
	// padding between circles, but this can go down as the size of the cell
	// decreases.
	circleSpace := ratio * idealCircleSpace

	// Edgewidth can shrink as cell size gets smaller.
	edgeWidth := helper.Round(ratio*float64(idealEdgeWidth), 0.01)

	// Font size.
	fontSize := int(math.Floor(ratio * float64(idealFontSize)))

	// Determine the maxium circle radius. Will be 9px for the default cell size
	// of 20px.
	maxRadius := helper.Round(float64(cellSizeHalf)-circleSpace, 0.01)

	// Calculate required top margin. Find the longest column name and assume it
	// is made entirely of the "W" character (which has a width of 11.33px at
	// in arial with a 12pt fontsize).
	longestColumnName := 0
	for _, colName := range columns {
		nameLength := len([]rune(colName))
		if nameLength > longestColumnName {
			longestColumnName = nameLength
		}
	}
	topMargin := int(math.Round(float64(longestColumnName) * 11.33 * ratio))

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
	leftMargin := int(math.Round(float64(longestRowName) * 11.33 * ratio))

	// Set plot dimensions.
	plotHeight := topMargin + (rowNum * cellSize)
	plotWidth := leftMargin + (colNum * cellSize)

	// Define svg.
	svgInit := fmt.Sprintf(
		"<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\""+
			" xml:space=\"preserve\" width=\"%d\" height=\"%d\" viewBox=\"0 0 %d %d\">\n",
		plotWidth, plotHeight, plotWidth, plotHeight,
	)
	svgSlice = append(svgSlice, svgInit)

	// Write column names.
	xOffset := fontSize / 2
	yOffset := topMargin - 2
	svgSlice = append(svgSlice, fmt.Sprintf("\t<g transform=\"translate(%d)\">\n", leftMargin))
	for i, colName := range columns {
		xPos := (i * cellSize) + xOffset
		text := fmt.Sprintf(
			"\t\t<text y=\"%d\" x=\"%d\" font-size=\"%d\""+
				" text-anchor=\"end\" transform=\"rotate(90, %d, %d)\">%s</text>\n",
			yOffset, xPos, fontSize, xPos, yOffset, colName,
		)
		svgSlice = append(svgSlice, text)
	}
	svgSlice = append(svgSlice, "\t</g>\n")

	// Write row names.
	xOffset = leftMargin - 2
	yOffset = cellSizeHalf + ((fontSize - 2) / 2)
	svgSlice = append(svgSlice, fmt.Sprintf("\t<g transform=\"translate(0, %d)\">\n", topMargin))
	for i, rowName := range rows {
		yPos := (i * cellSize) + yOffset
		text := fmt.Sprintf(
			"\t\t<text y=\"%d\" x=\"%d\" font-size=\"%d\" text-anchor=\"end\">%s</text>\n",
			yPos, xOffset, fontSize, rowName,
		)
		svgSlice = append(svgSlice, text)
	}
	svgSlice = append(svgSlice, "\t</g>\n")

	// Get color gradient.
	colorGradient := ColorGradient(options["colorSpace"].(string), 101, options["invert"].(bool))

	// Get function for determining score edge color.
	edgeColorFunc := ScoreColorFunc(options["scoreType"].(string), options["primary"].(float64), options["secondary"].(float64), 100)

	// Write rows.
	svgSlice = append(svgSlice, fmt.Sprintf("\t<g transform=\"translate(%d, %d)\">\n", leftMargin, topMargin))
	for i, row := range matrix {
		// Set x position.
		iPos := (i * cellSize) + cellSizeHalf

		// Draw dots.
		for j, value := range row {
			if value > 0 {
				// Get fill color.
				var fill string
				if value > options["maximumAbundance"].(float64) {
					fill = colorGradient[100]
				} else {
					index := int(math.Round(value / options["maximumAbundance"].(float64) * 100))
					fill = colorGradient[index]
				}

				// Edge color
				edgeColorIndex := edgeColorFunc(scores[i][j])
				edgeColor := colorGradient[edgeColorIndex]

				// Get circle radius.
				radius := helper.Round(ratios[i][j]*maxRadius, 0.01)

				// Draw circle.
				circle := fmt.Sprintf(
					"\t\t<circle fill=\"%s\" cy=\"%d\" cx=\"%d\" r=\"%f\""+
						" stroke=\"%s\" stroke-width=\"%f\"/>\n",
					fill, iPos, (j*cellSize)+cellSizeHalf, radius, edgeColor, edgeWidth,
				)
				svgSlice = append(svgSlice, circle)
			}
		}
	}
	svgSlice = append(svgSlice, "\t</g>\n")

	// Add column label.
	xOffset = leftMargin + ((plotWidth - leftMargin) / 2)
	text := fmt.Sprintf(
		"\t<text y=\"10\" x=\"%d\" font-size=\"12\""+
			" text-anchor=\"middle\">%s</text>\n",
		xOffset, options["colLabel"],
	)
	svgSlice = append(svgSlice, text)

	// Add row label.
	yOffset = topMargin + ((plotHeight - topMargin) / 2)
	text = fmt.Sprintf(
		"\t<text y=\"%d\" x=\"10\" font-size=\"12\""+
			" text-anchor=\"middle\" transform=\"rotate(-90, 10, %d)\">%s</text>\n",
		yOffset, yOffset, options["rowLabel"],
	)
	svgSlice = append(svgSlice, text)

	// Add bounding box.
	rect := fmt.Sprintf(
		"\t<rect fill=\"none\" y=\"%d\" x=\"%d\" width=\"%d\" height=\"%d\""+
			" stroke=\"#000000\" stroke-width=\"0.5\" />\n",
		topMargin, leftMargin, plotWidth-leftMargin, plotHeight-topMargin,
	)
	svgSlice = append(svgSlice, rect)

	// Add end element wrapper for svg.
	svgSlice = append(svgSlice, "</svg>\n")

	// Create svg string.
	var buffer bytes.Buffer
	for _, value := range svgSlice {
		buffer.WriteString(value)
	}
	svg = buffer.String()

	return
}
