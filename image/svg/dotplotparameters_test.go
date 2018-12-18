package svg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDotplotParametersDefine(t *testing.T) {
	dims := HeatmapDimensions{
		cellSize: 20,
		ratio:    1,
	}

	// TEST
	params := DotplotParametersDefine(dims)
	expected := DotplotParameters{
		cellSizeHalf: 10,
		edgeWidth:    2,
		maxRadius:    8.5,
	}
	assert.Equal(t, expected, params, "Heatmap dimensions are not correct")
}
