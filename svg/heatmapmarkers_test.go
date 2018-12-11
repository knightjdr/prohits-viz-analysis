package svg

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestHeatmapMarkers(t *testing.T) {
	dims := HDimensions{
		cellSize:   20,
		leftMargin: 50,
		topMargin:  50,
	}
	markers := typedef.Markers{
		Color: "#000000",
		List: []typedef.Marker{
			{Height: 2, Width: 2, X: 0, Y: 1},
		},
	}

	// TEST
	expected := "\t<g transform=\"translate(50, 50)\">\n" +
		"\t\t<rect y=\"20\" x=\"0\" width=\"40\" height=\"40\" stroke=\"#000000\" stroke-width=\"1\" fill=\"none\"/>\n" +
		"\t</g>\n"
	assert.Equal(t, expected, HeatmapMarkers(markers, dims), "Markers svg element is not correct")
}
