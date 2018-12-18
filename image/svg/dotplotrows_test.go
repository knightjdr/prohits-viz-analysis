package svg

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestScoreColorFunc(t *testing.T) {
	// TEST: when larger scores are better
	scoreFunc := scoreColorFunc("gte", float64(1), float64(0.5), 100)
	scores := []float64{0.1, 0.4, 0.5, 0.75, 1, 1.5}
	expected := []int{25, 25, 50, 50, 100, 100}
	for i, value := range scores {
		assert.Equal(t, expected[i], scoreFunc(value), "gte score function not returning correct gradient index")
	}

	// TEST: when smaller scores are better
	scoreFunc = scoreColorFunc("lte", float64(0.5), float64(1), 100)
	scores = []float64{0.1, 0.4, 0.5, 0.75, 1, 1.5}
	expected = []int{100, 100, 100, 50, 50, 25}
	for i, value := range scores {
		assert.Equal(t, expected[i], scoreFunc(value), "lte score function not returning correct gradient index")
	}
}

func TestDotplotRows(t *testing.T) {
	// Mock writeString.
	svg := make([]string, 0)
	writeString := func(str string) {
		svg = append(svg, str)
	}

	dims := HeatmapDimensions{
		cellSize:   20,
		leftMargin: 50,
		topMargin:  50,
	}
	dotplotParameters := DotplotParameters{
		cellSizeHalf: 10,
		edgeWidth:    2,
		maxRadius:    8.5,
	}
	matrices := new(typedef.Matrices)
	matrices.Abundance = [][]float64{
		{25, 5, 50.2},
		{100, 30, 7},
		{5, 2.3, 8},
	}
	matrices.Ratio = [][]float64{
		{1, 0.5, 0.3},
		{1, 0.3, 0.1},
		{0.5, 0.25, 1},
	}
	matrices.Score = [][]float64{
		{0.01, 0, 0.02},
		{0, 0.01, 0.01},
		{0.02, 0.1, 0.01},
	}
	parameters := typedef.Parameters{
		EdgeColor:       "blueBlack",
		FillColor:       "blueBlack",
		AbundanceCap:    50,
		InvertColor:     false,
		PrimaryFilter:   0.01,
		SecondaryFilter: 0.05,
		ScoreType:       "lte",
	}

	// TEST: create svg.
	DotplotRows(matrices, dims, dotplotParameters, parameters, writeString)
	expected := "\t<g id=\"minimap\" transform=\"translate(50, 50)\">\n" +
		"\t\t<circle fill=\"#0040ff\" cy=\"10\" cx=\"10\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#ccd9ff\" cy=\"10\" cx=\"30\" r=\"4.250000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#000000\" cy=\"10\" cx=\"50\" r=\"2.550000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#000000\" cy=\"30\" cx=\"10\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#0033cc\" cy=\"30\" cx=\"30\" r=\"2.550000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#b8c9ff\" cy=\"30\" cx=\"50\" r=\"0.850000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#ccd9ff\" cy=\"50\" cx=\"10\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#e6ecff\" cy=\"50\" cx=\"30\" r=\"2.130000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
		"\t\t<circle fill=\"#adc2ff\" cy=\"50\" cx=\"50\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
		"\t</g>\n"
	assert.Equal(t, expected, helper.StringConcat(svg), "Dotplot rows svg element is not correct")
}
