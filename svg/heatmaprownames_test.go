package svg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeatmapRowNames(t *testing.T) {
	dims := HDimensions{
		cellSize:   20,
		fontSize:   12,
		leftMargin: 50,
		topMargin:  50,
	}
	rows := []string{"prey1", "prey2", "prey3"}

	// TEST
	expected := "\t<g transform=\"translate(0, 50)\">\n" +
		"\t\t<text y=\"15\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey1</text>\n" +
		"\t\t<text y=\"35\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey2</text>\n" +
		"\t\t<text y=\"55\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey3</text>\n" +
		"\t</g>\n"
	assert.Equal(t, expected, HeatmapRowNames(dims, rows), "Row svg element is not correct")
}
