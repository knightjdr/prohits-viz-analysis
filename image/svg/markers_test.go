package svg

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestMarkers(t *testing.T) {
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
	markers := typedef.Markers{
		Color: "#000000",
		List: []typedef.Marker{
			{Height: 2, Width: 2, X: 0, Y: 1},
		},
	}

	// TEST: with markers
	expected := "\t<g transform=\"translate(50, 50)\">\n" +
		"\t\t<rect y=\"20\" x=\"0\" width=\"40\" height=\"40\" stroke=\"#000000\" stroke-width=\"1\" fill=\"none\"/>\n" +
		"\t</g>\n"
	Markers(markers, dims, writeString)
	assert.Equal(t, expected, helper.StringConcat(svg), "Markers svg element is not correct")

	// TEST: without markers
	svg = make([]string, 0)
	expected = ""
	Markers(typedef.Markers{}, dims, writeString)
	assert.Equal(t, expected, helper.StringConcat(svg), "Markers svg element should be empty")
}
