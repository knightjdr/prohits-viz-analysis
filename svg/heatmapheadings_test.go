package svg

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestHeatmapHeadings(t *testing.T) {
	dims := HDimensions{
		svgHeight: 250,
		svgWidth:  150,
	}
	parameters := typedef.Parameters{
		Condition: "Bait",
		Readout:   "Prey",
	}
	expected := "\t<text y=\"10\" x=\"75\" font-size=\"12\" text-anchor=\"middle\">Bait</text>\n" +
		"\t<text y=\"125\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 125)\">Prey</text>\n"
	assert.Equal(t, expected, HeatmapHeadings(dims, parameters), "Headings svg element does not match expected string")
}
