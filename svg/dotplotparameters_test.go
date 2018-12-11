package svg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDotplotParameters(t *testing.T) {
	dims := HDimensions{
		cellSize: 20,
		ratio:    1,
	}

	// TEST
	params := DotplotParameters(dims)
	expected := DParameters{
		cellSizeHalf: 10,
		edgeWidth:    2,
		maxRadius:    8.5,
	}
	assert.Equal(t, expected, params, "Heatmap dimensions are not correct")
}
