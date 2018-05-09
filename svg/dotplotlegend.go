package svg

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// DotplotLegend creates a legend for a dotplot.
func DotplotLegend(
	colorSpace, title string,
	numColors int,
	min, max, primary, secondary float64,
	score, scoreType string,
) (svg string) {
	// Get color gradient.
	gradient := ColorGradient(colorSpace, numColors, false)

	// Define svg.
	svgSlice := make([]string, 0)
	svgInit := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"" +
		" xml:space=\"preserve\" width=\"200\" height=\"240\" viewBox=\"0 0 200 240\">\n"
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

	// Draw relative abundance graphic.
	svgSlice = append(svgSlice, "\t<g>\n")
	smallCirc := "\t\t<circle fill=\"#000000\" cy=\"100\" cx=\"60\" r=\"6\" />\n"
	svgSlice = append(svgSlice, smallCirc)
	largeCirc := "\t\t<circle fill=\"#000000\" cy=\"100\" cx=\"135\" r=\"12\" />\n"
	svgSlice = append(svgSlice, largeCirc)
	arrow := "\t\t<line fill=\"none\" stroke=\"#000000\" stroke-width=\"1\" x1=\"70\" y1=\"100\"" +
		" x2=\"119\" y2=\"100\"/>\n" +
		"\t\t<polygon fill=\"#000000\" points=\"110,96 112,100 110,104 119,100\"/>\n"
	svgSlice = append(svgSlice, arrow)
	relText := "\t\t<text y=\"130\" x=\"100\" font-size=\"12\"" +
		" text-anchor=\"middle\">Relative abundance</text>\n"
	svgSlice = append(svgSlice, relText)
	svgSlice = append(svgSlice, "\t</g>\n")

	// Draw score graphic.
	svgSlice = append(svgSlice, "\t<g>\n")
	scoreText := fmt.Sprintf("\t\t<text y=\"220\" x=\"100\" font-size=\"12\""+
		" text-anchor=\"middle\">%s</text>\n",
		score,
	)
	svgSlice = append(svgSlice, scoreText)

	// Create primary circle.
	primaryCirc := fmt.Sprintf("\t\t<circle fill=\"none\" cy=\"165\" cx=\"50\" r=\"12\""+
		" stroke=\"%s\" stroke-width=\"2\" />\n",
		gradient[numColors-1],
	)
	svgSlice = append(svgSlice, primaryCirc)

	// Create primary text based on scoretype.
	var primaryText string
	if scoreType == "gte" {
		primaryText = fmt.Sprintf("\t\t<text y=\"195\" x=\"50\" font-size=\"12\""+
			" text-anchor=\"middle\">≥ %s</text>\n",
			strconv.FormatFloat(primary, 'f', -1, 64),
		)
	} else {
		primaryText = fmt.Sprintf("\t\t<text y=\"195\" x=\"50\" font-size=\"12\""+
			" text-anchor=\"middle\">≤ %s</text>\n",
			strconv.FormatFloat(primary, 'f', -1, 64),
		)
	}
	svgSlice = append(svgSlice, primaryText)

	// Create secondary circle.
	secondaryCirc := fmt.Sprintf("\t\t<circle fill=\"none\" cy=\"165\" cx=\"100\" r=\"12\""+
		" stroke=\"%s\" stroke-width=\"2\" />\n",
		gradient[(numColors-1)/2],
	)
	svgSlice = append(svgSlice, secondaryCirc)

	// Create secondary text based on scoretype.
	var secondaryText string
	if scoreType == "gte" {
		secondaryText = fmt.Sprintf("\t\t<text y=\"195\" x=\"100\" font-size=\"12\""+
			" text-anchor=\"middle\">≥ %s</text>\n",
			strconv.FormatFloat(secondary, 'f', -1, 64),
		)
	} else {
		secondaryText = fmt.Sprintf("\t\t<text y=\"195\" x=\"100\" font-size=\"12\""+
			" text-anchor=\"middle\">≤ %s</text>\n",
			strconv.FormatFloat(secondary, 'f', -1, 64),
		)
	}
	svgSlice = append(svgSlice, secondaryText)

	// Create tertiaryCirc circle.
	tertiaryCirc := fmt.Sprintf("\t\t<circle fill=\"none\" cy=\"165\" cx=\"150\" r=\"12\""+
		" stroke=\"%s\" stroke-width=\"2\" />\n",
		gradient[(numColors-1)/4],
	)
	svgSlice = append(svgSlice, tertiaryCirc)

	// Create tertiary text based on scoretype.
	var tertiaryText string
	if scoreType == "gte" {
		tertiaryText = fmt.Sprintf("\t\t<text y=\"195\" x=\"150\" font-size=\"12\""+
			" text-anchor=\"middle\">< %s</text>\n",
			strconv.FormatFloat(secondary, 'f', -1, 64),
		)
	} else {
		tertiaryText = fmt.Sprintf("\t\t<text y=\"195\" x=\"150\" font-size=\"12\""+
			" text-anchor=\"middle\">> %s</text>\n",
			strconv.FormatFloat(secondary, 'f', -1, 64),
		)
	}
	svgSlice = append(svgSlice, tertiaryText)
	svgSlice = append(svgSlice, "\t</g>\n")

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
