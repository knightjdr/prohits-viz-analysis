package file

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("svg", 0755)

	// Data.
	dist := [][]float64{
		{0.2, 0.5, 1},
		{0.7, 0.8, 1},
		{1, 0.2, 0.5},
	}
	sorted := []string{"row2", "row3", "row1"}
	parameters := typedef.Parameters{
		FillColor: "blueBlack",
		Readout:   "Readouts",
	}

	// TEST: readout-readout heatmap
	want := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"105\" height=\"105\" viewBox=\"0 0 105 105\">\n" +
		"\t<g transform=\"translate(45)\">\n" +
		"\t\t<text y=\"43\" x=\"6\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 6, 43)\">row2</text>\n" +
		"\t\t<text y=\"43\" x=\"26\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 26, 43)\">row3</text>\n" +
		"\t\t<text y=\"43\" x=\"46\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 46, 43)\">row1</text>\n" +
		"\t</g>\n" +
		"\t<g transform=\"translate(0, 45)\">\n" +
		"\t\t<text y=\"15\" x=\"43\" font-size=\"12\" text-anchor=\"end\">row2</text>\n" +
		"\t\t<text y=\"35\" x=\"43\" font-size=\"12\" text-anchor=\"end\">row3</text>\n" +
		"\t\t<text y=\"55\" x=\"43\" font-size=\"12\" text-anchor=\"end\">row1</text>\n" +
		"\t</g>\n" +
		"\t<g id=\"minimap\" transform=\"translate(45, 45)\">\n" +
		"\t\t<rect fill=\"#001966\" y=\"0\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#0040ff\" y=\"0\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#ffffff\" y=\"0\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#668cff\" y=\"20\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#99b3ff\" y=\"20\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#ffffff\" y=\"20\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#ffffff\" y=\"40\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#001966\" y=\"40\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#0040ff\" y=\"40\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t</g>\n" +
		"\t<text y=\"10\" x=\"75\" font-size=\"12\" text-anchor=\"middle\">Readouts</text>\n" +
		"\t<text y=\"75\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 75)\">Readouts</text>\n" +
		"</svg>\n"
	Distance(dist, sorted, "Readouts", "readout-readout", parameters)
	svg, _ := afero.ReadFile(fs.Instance, "svg/readout-readout.svg")
	assert.Equal(t, want, string(svg), "Readout-readout svg not generated correctly")
}
