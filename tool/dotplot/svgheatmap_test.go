package dotplot

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestSvgHeatmap(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("svg", 0755)

	// Data.
	matrix := [][]float64{
		{5, 55, 15},
		{3, 7, 1},
		{75, 0.2, 0.5},
	}
	sortedColumns := []string{"col1", "col3", "col2"}
	sortedRows := []string{"row2", "row3", "row1"}

	// TEST1: condition-readout heatmap
	want := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"105\" height=\"105\" viewBox=\"0 0 105 105\">\n" +
		"\t<g transform=\"translate(45)\">\n" +
		"\t\t<text y=\"43\" x=\"6\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 6, 43)\">col1</text>\n" +
		"\t\t<text y=\"43\" x=\"26\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 26, 43)\">col3</text>\n" +
		"\t\t<text y=\"43\" x=\"46\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 46, 43)\">col2</text>\n" +
		"\t</g>\n" +
		"\t<g transform=\"translate(0, 45)\">\n" +
		"\t\t<text y=\"15\" x=\"43\" font-size=\"12\" text-anchor=\"end\">row2</text>\n" +
		"\t\t<text y=\"35\" x=\"43\" font-size=\"12\" text-anchor=\"end\">row3</text>\n" +
		"\t\t<text y=\"55\" x=\"43\" font-size=\"12\" text-anchor=\"end\">row1</text>\n" +
		"\t</g>\n" +
		"\t<g id=\"minimap\" transform=\"translate(45, 45)\">\n" +
		"\t\t<rect fill=\"#ccd9ff\" y=\"0\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#000000\" y=\"0\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#668cff\" y=\"0\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#e0e8ff\" y=\"20\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#b8c9ff\" y=\"20\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#f5f7ff\" y=\"20\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#000000\" y=\"40\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#ffffff\" y=\"40\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#fafbff\" y=\"40\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t</g>\n" +
		"\t<text y=\"10\" x=\"75\" font-size=\"12\" text-anchor=\"middle\">Conditions</text>\n" +
		"\t<text y=\"75\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 75)\">Readouts</text>\n" +
		"</svg>\n"
	SvgHeatmap(matrix, sortedColumns, sortedRows, "blueBlack", 50, false)
	svg, _ := afero.ReadFile(fs.Instance, "svg/heatmap.svg")
	assert.Equal(t, want, string(svg), "Heatmap svg not generated correctly")
}
