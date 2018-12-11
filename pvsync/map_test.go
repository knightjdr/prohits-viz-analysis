package main

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("minimap", 0755)

	// Data.
	abundance := [][]float64{
		{5, 55, 15},
		{3, 7, 1},
		{75, 0.2, 0.5},
	}
	parameters := typedef.Parameters{
		AbundanceCap:    50,
		EdgeColor:       "blueBlack",
		FillColor:       "blueBlack",
		PrimaryFilter:   0.01,
		SecondaryFilter: 0.05,
		ScoreType:       "lte",
	}
	ratios := [][]float64{
		{0.2, 0.5, 1},
		{0.7, 0.8, 1},
		{1, 0.2, 0.5},
	}
	scores := [][]float64{
		{0.01, 0.05, 0.08},
		{1, 0, 0.5},
		{0.2, 0.7, 0.01},
	}
	sortedColumns := []string{"col1", "col3", "col2"}
	sortedRows := []string{"row2", "row3", "row1"}

	// TEST: dotplot
	want := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"60\" height=\"60\" viewBox=\"0 0 60 60\">\n" +
		"\t<g id=\"minimap\" transform=\"translate(0, 0)\">\n" +
		"\t\t<circle fill=\"#ccd9ff\" cy=\"10\" cx=\"10\" r=\"1.700000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#000000\" cy=\"10\" cx=\"30\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#668cff\" cy=\"10\" cx=\"50\" r=\"8.500000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#e0e8ff\" cy=\"30\" cx=\"10\" r=\"5.950000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#b8c9ff\" cy=\"30\" cx=\"30\" r=\"6.800000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#f5f7ff\" cy=\"30\" cx=\"50\" r=\"8.500000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#000000\" cy=\"50\" cx=\"10\" r=\"8.500000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#ffffff\" cy=\"50\" cx=\"30\" r=\"1.700000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#fafbff\" cy=\"50\" cx=\"50\" r=\"4.250000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
		"\t</g>\n" +
		"</svg>\n"
	Map("dotplot", abundance, ratios, scores, sortedColumns, sortedRows, parameters)
	svg, _ := afero.ReadFile(fs.Instance, "minimap/dotplot.svg")
	assert.Equal(t, want, string(svg), "Dotplot svg not generated correctly")

	// TEST: heatmap
	want = "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"60\" height=\"60\" viewBox=\"0 0 60 60\">\n" +
		"\t<g id=\"minimap\" transform=\"translate(0, 0)\">\n" +
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
		"</svg>\n"
	Map("heatmap", abundance, ratios, scores, sortedColumns, sortedRows, parameters)
	svg, _ = afero.ReadFile(fs.Instance, "minimap/heatmap.svg")
	assert.Equal(t, want, string(svg), "Heatmap svg not generated correctly")
}
