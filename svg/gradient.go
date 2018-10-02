package svg

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// Gradient draws a color gradient for a legend. It can either be a complete svg
// or just the elements for the gradient (i.e. not wrapped by <svg>)
func Gradient(colorSpace, title string, numColors int, min, max float64) (svg string) {
	// Get color gradient.
	gradient := colorGradient(colorSpace, numColors, false)

	// Define svg.
	svgSlice := make([]string, 0)
	svgInit := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"" +
		" xml:space=\"preserve\" width=\"200\" height=\"80\" viewBox=\"0 0 200 80\">\n"
	svgSlice = append(svgSlice, svgInit)

	// Add title
	titleText := fmt.Sprintf("\t<text y=\"20\" x=\"100\" font-size=\"12\""+
		" text-anchor=\"middle\">%s</text>\n",
		title,
	)
	svgSlice = append(svgSlice, titleText)

	// Create gradient. CellWidth is the width of each gradient cell.
	cellWidth := helper.Round(float64(150)/float64(numColors), 0.01)
	svgSlice = append(svgSlice, "\t<g>\n")
	for i, color := range gradient {
		xPos := (float64(i) * cellWidth) + 25
		rect := fmt.Sprintf(
			"\t\t<rect fill=\"%s\" y=\"30\" x=\"%f\" width=\"%f\" height=\"20\" />\n",
			color, xPos, cellWidth,
		)
		svgSlice = append(svgSlice, rect)
	}
	svgSlice = append(svgSlice, "\t</g>\n")

	// Draw border around gradient.
	border := "\t<rect fill=\"none\" y=\"30\" x=\"25\" width=\"150\" height=\"20\"" +
		" stroke=\"#000000\" stroke-width=\"1\"/>\n"
	svgSlice = append(svgSlice, border)

	// Add min and max labels.
	maxLabel := fmt.Sprintf("\t<text y=\"65\" x=\"175\" font-size=\"12\""+
		" text-anchor=\"middle\">%s</text>\n",
		strconv.FormatFloat(max, 'f', -1, 64),
	)
	svgSlice = append(svgSlice, maxLabel)
	minLabel := fmt.Sprintf("\t<text y=\"65\" x=\"25\" font-size=\"12\""+
		" text-anchor=\"middle\">%s</text>\n",
		strconv.FormatFloat(min, 'f', -1, 64),
	)
	svgSlice = append(svgSlice, minLabel)

	// Terminate svg.
	svgSlice = append(svgSlice, "</svg>\n")

	// Create svg string.
	var buffer bytes.Buffer
	for _, value := range svgSlice {
		buffer.WriteString(value)
	}
	svg = buffer.String()

	return
}
