package svg

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestAnnotations(t *testing.T) {
	// Mock writeString.
	svg := make([]string, 0)
	writeString := func(str string) {
		svg = append(svg, str)
	}

	annotations := typedef.Annotations{
		FontSize: 15,
		List: []typedef.Annotation{
			{Text: "a", X: 0.5, Y: 0.2},
		},
	}
	dims := HeatmapDimensions{
		leftMargin: 50,
		plotWidth:  100,
		plotHeight: 200,
		topMargin:  50,
	}

	// TEST: with annotations
	expected := "\t<g transform=\"translate(50, 50)\">\n" +
		"\t\t<text y=\"40\" x=\"50\" font-size=\"15\" text-anchor=\"middle\">a</text>\n" +
		"\t</g>\n"
	Annotations(annotations, dims, writeString)
	assert.Equal(t, expected, helper.StringConcat(svg), "Annotations svg element is not correct")

	// TEST: without annotations
	svg = make([]string, 0)
	expected = ""
	Annotations(typedef.Annotations{}, dims, writeString)
	assert.Equal(t, expected, helper.StringConcat(svg), "Annotations svg element should be empty")
}
