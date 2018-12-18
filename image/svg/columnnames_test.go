package svg

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/stretchr/testify/assert"
)

func TestColumnNames(t *testing.T) {
	// Mock writeString.
	svg := make([]string, 0)
	writeString := func(str string) {
		svg = append(svg, str)
	}

	columns := []string{"bait1", "bait2", "bait3"}
	dims := HeatmapDimensions{
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
	ColumnNames(dims, columns, writeString)
	assert.Equal(t, expected, helper.StringConcat(svg), "Column svg element is not correct")
}
