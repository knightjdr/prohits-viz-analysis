package svg

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/stretchr/testify/assert"
)

func TestBoundingBox(t *testing.T) {
	// Mock writeString.
	svg := make([]string, 0)
	writeString := func(str string) {
		svg = append(svg, str)
	}

	dims := HeatmapDimensions{
		leftMargin: 50,
		svgHeight:  250,
		svgWidth:   150,
		topMargin:  50,
	}
	expected := "\t<rect fill=\"none\" y=\"50\" x=\"50\" width=\"100\" height=\"200\"" +
		" stroke=\"#000000\" stroke-width=\"0.5\" />\n"
	BoundingBox(dims, writeString)
	assert.Equal(t, expected, helper.StringConcat(svg), "Bounding box does not match expected string")
}
