package svg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeatmapColumnNames(t *testing.T) {
	columns := []string{"bait1", "bait2", "bait3"}
	dims := HDimensions{
		cellSize:   20,
		fontSize:   12,
		leftMargin: 50,
		topMargin:  50,
	}

	// TEST
	expected := "\t<g transform=\"translate(50)\">\n" +
		"\t\t<text y=\"48\" x=\"6\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 6, 48)\">bait1</text>\n" +
		"\t\t<text y=\"48\" x=\"26\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 26, 48)\">bait2</text>\n" +
		"\t\t<text y=\"48\" x=\"46\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 46, 48)\">bait3</text>\n" +
		"\t</g>\n"
	assert.Equal(t, expected, HeatmapColumnNames(dims, columns), "Column svg element is not correct")
}
