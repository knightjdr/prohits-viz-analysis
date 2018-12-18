package svg

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestHeadings(t *testing.T) {
	// Mock writeString.
	svg := make([]string, 0)
	writeString := func(str string) {
		svg = append(svg, str)
	}

	dims := HeatmapDimensions{
		svgHeight: 250,
		svgWidth:  150,
	}

	// TEST: both x and y label supplied
	parameters := typedef.Parameters{
		XLabel: "Bait",
		YLabel: "Prey",
	}
	expected := "\t<text y=\"10\" x=\"75\" font-size=\"12\" text-anchor=\"middle\">Bait</text>\n" +
		"\t<text y=\"125\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 125)\">Prey</text>\n"
	Headings(parameters, dims, writeString)
	assert.Equal(t, expected, helper.StringConcat(svg), "Headings svg element does not match expected string")

	// TEST: both only x label supplied
	svg = make([]string, 0)
	parameters = typedef.Parameters{
		XLabel: "Bait",
	}
	expected = "\t<text y=\"10\" x=\"75\" font-size=\"12\" text-anchor=\"middle\">Bait</text>\n"
	Headings(parameters, dims, writeString)
	assert.Equal(t, expected, helper.StringConcat(svg), "Headings svg element does not match x label only expected string")

	// TEST: both only y label supplied
	svg = make([]string, 0)
	parameters = typedef.Parameters{
		YLabel: "Prey",
	}
	expected = "\t<text y=\"125\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 125)\">Prey</text>\n"
	Headings(parameters, dims, writeString)
	assert.Equal(t, expected, helper.StringConcat(svg), "Headings svg element does not match y label only expected string")

	// TEST: no labels supplied
	svg = make([]string, 0)
	parameters = typedef.Parameters{}
	expected = ""
	Headings(parameters, dims, writeString)
	assert.Equal(t, expected, helper.StringConcat(svg), "Headings svg element does not match missing labels expected string")
}
