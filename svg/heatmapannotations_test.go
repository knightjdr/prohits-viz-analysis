package svg

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestHeatmapAnnotations(t *testing.T) {
	annotations := typedef.Annotations{
		FontSize: 15,
		List: []typedef.Annotation{
			{Text: "a", X: 0.5, Y: 0.2},
		},
	}
	dims := HDimensions{
		leftMargin: 50,
		plotWidth:  100,
		plotHeight: 200,
		topMargin:  50,
	}

	// TEST
	expected := "\t<g transform=\"translate(50, 50)\">\n" +
		"\t\t<text y=\"40\" x=\"50\" font-size=\"15\" text-anchor=\"middle\">a</text>\n" +
		"\t</g>\n"
	assert.Equal(t, expected, HeatmapAnnotations(annotations, dims), "Annotations svg element is not correct")
}
