package dotplot

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestSvgDotplot(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("svg", 0755)

	// Data.
	abundance := [][]float64{
		{5, 55, 15},
		{3, 7, 1},
		{75, 0.2, 0.5},
	}
	params := typedef.Parameters{
		ColorSpace:       "blueBlack",
		MaximumAbundance: float64(50),
		PrimaryFilter:    0.01,
		SecondaryFilter:  0.05,
		ScoreType:        "lte",
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

	// TEST1: bait-prey dotplot
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
		"\t<text y=\"10\" x=\"75\" font-size=\"12\" text-anchor=\"middle\">Baits</text>\n" +
		"\t<text y=\"75\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 75)\">Preys</text>\n" +
		"\t<rect fill=\"none\" y=\"45\" x=\"45\" width=\"60\" height=\"60\" stroke=\"#000000\" stroke-width=\"0.5\" />\n" +
		"</svg>\n"
	SvgDotplot(abundance, ratios, scores, sortedColumns, sortedRows, params)
	svg, _ := afero.ReadFile(fs.Instance, "svg/dotplot.svg")
	assert.Equal(t, want, string(svg), "Dotplot svg not generated correctly")
}
