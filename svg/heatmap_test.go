package svg

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestHeatmap(t *testing.T) {
	columns := []string{"bait1", "bait2", "bait3"}
	matrix := [][]float64{
		{25, 5, 50.2},
		{100, 30, 7},
		{5, 2.3, 8},
	}
	annotations := typedef.Annotations{
		FontSize: 15,
		List: []typedef.Annotation{
			{Text: "a", X: 0.5, Y: 0.2},
		},
	}
	markers := typedef.Markers{
		Color: "#000000",
		List: []typedef.Marker{
			{Height: 2, Width: 2, X: 0, Y: 1},
		},
	}
	options := map[string]interface{}{
		"annotationFontSize": int(15),
		"colLabel":           "Baits",
		"fillColor":          "blueBlack",
		"markerColor":        "#000000",
		"maximumAbundance":   float64(50),
		"invert":             false,
		"rowLabel":           "Preys",
	}
	rows := []string{"prey1", "prey2", "prey3"}

	// TEST1: create svg.
	want := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"117\" height=\"117\" viewBox=\"0 0 117 117\">\n" +
		"\t<g transform=\"translate(57)\">\n" +
		"\t\t<text y=\"55\" x=\"6\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 6, 55)\">bait1</text>\n" +
		"\t\t<text y=\"55\" x=\"26\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 26, 55)\">bait2</text>\n" +
		"\t\t<text y=\"55\" x=\"46\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 46, 55)\">bait3</text>\n" +
		"\t</g>\n" +
		"\t<g transform=\"translate(0, 57)\">\n" +
		"\t\t<text y=\"15\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey1</text>\n" +
		"\t\t<text y=\"35\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey2</text>\n" +
		"\t\t<text y=\"55\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey3</text>\n" +
		"\t</g>\n" +
		"\t<g id=\"minimap\" transform=\"translate(57, 57)\">\n" +
		"\t\t<rect fill=\"#0040ff\" y=\"0\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#ccd9ff\" y=\"0\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#000000\" y=\"0\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#000000\" y=\"20\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#0033cc\" y=\"20\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#b8c9ff\" y=\"20\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#ccd9ff\" y=\"40\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#e6ecff\" y=\"40\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#adc2ff\" y=\"40\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t</g>\n" +
		"\t<g transform=\"translate(57, 57)\">\n" +
		"\t\t<rect y=\"20\" x=\"0\" width=\"40\" height=\"40\" stroke=\"#000000\" stroke-width=\"1\" fill=\"none\"/>\n" +
		"\t</g>\n" +
		"\t<g transform=\"translate(57, 57)\">\n" +
		"\t\t<text y=\"12\" x=\"30\" font-size=\"15\" text-anchor=\"middle\">a</text>\n" +
		"\t</g>\n" +
		"\t<text y=\"10\" x=\"87\" font-size=\"12\" text-anchor=\"middle\">Baits</text>\n" +
		"\t<text y=\"87\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 87)\">Preys</text>\n" +
		"</svg>\n"
	svg := Heatmap(matrix, annotations, markers, columns, rows, options)
	assert.Equal(t, want, svg, "Heatmap svg is not correct")
}
