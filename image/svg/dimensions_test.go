package svg

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestDimensions(t *testing.T) {
	data := new(typedef.Matrices)
	data.Abundance = [][]float64{
		{25, 5, 50.2},
		{100, 30, 7},
		{5, 2.3, 8},
	}
	data.Conditions = []string{"bait1", "bait2", "bait3"}
	data.Readouts = []string{"prey1", "prey2", "prey3"}

	// TEST: dimensions for full image with column and row names
	dims := Dimensions(data, false)
	expected := HeatmapDimensions{
		cellSize:   20,
		fontSize:   12,
		leftMargin: 57,
		plotHeight: 60,
		plotWidth:  60,
		ratio:      1,
		svgHeight:  117,
		svgWidth:   117,
		topMargin:  57,
	}
	assert.Equal(t, expected, dims, "Heatmap dimensions are not correct")

	// TEST: dimensions for minimap alone
	dims = Dimensions(data, true)
	expected = HeatmapDimensions{
		cellSize:   20,
		fontSize:   0,
		leftMargin: 0,
		plotHeight: 60,
		plotWidth:  60,
		ratio:      1,
		svgHeight:  60,
		svgWidth:   60,
		topMargin:  0,
	}
	assert.Equal(t, expected, dims, "Heatmap dimensions are not correct")
}
