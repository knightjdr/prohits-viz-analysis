package svg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoundingBox(t *testing.T) {
	dims := HDimensions{
		leftMargin: 50,
		svgHeight:  250,
		svgWidth:   150,
		topMargin:  50,
	}
	expected := "\t<rect fill=\"none\" y=\"50\" x=\"50\" width=\"100\" height=\"200\"" +
		" stroke=\"#000000\" stroke-width=\"0.5\" />\n"
	assert.Equal(t, expected, BoundingBox(dims), "Bounding box does not match expected string")
}
