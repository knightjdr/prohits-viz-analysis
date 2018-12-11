package svg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeatmapHeader(t *testing.T) {
	dims := HDimensions{
		svgHeight: 250,
		svgWidth:  150,
	}
	expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"" +
		" xml:space=\"preserve\" width=\"150\" height=\"250\" viewBox=\"0 0 150 250\">\n"
	assert.Equal(t, expected, HeatmapHeader(dims), "Header does not match expected string")
}
