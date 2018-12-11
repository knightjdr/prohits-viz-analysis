package svg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeatmapDimensions(t *testing.T) {
	columns := []string{"bait1", "bait2", "bait3"}
	abundance := [][]float64{
		{25, 5, 50.2},
		{100, 30, 7},
		{5, 2.3, 8},
	}
	rows := []string{"prey1", "prey2", "prey3"}

	// TEST: Heatmap dimension not for minimap
	dims := HeatmapDimensions(abundance, columns, rows, false)
	expected := HDimensions{
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

	// TEST: Heatmap dimension as minimap
	dims = HeatmapDimensions(abundance, columns, rows, true)
	expected = HDimensions{
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
