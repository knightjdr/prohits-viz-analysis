package heatmap

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write annotations", func() {
	It("should write annotations", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Heatmap{
			Annotations: types.Annotations{
				FontSize: 15,
				List: map[string]types.Annotation{
					"a": {
						Position: types.AnnotationPosition{X: 0.5, Y: 0.2},
						Text:     "a",
					},
				},
			},
			LeftMargin: 50,
			PlotWidth:  100,
			PlotHeight: 200,
			TopMargin:  50,
		}

		expected := "\t<g transform=\"translate(50, 50)\">\n" +
			"\t\t<text y=\"40\" x=\"50\" font-size=\"15\" text-anchor=\"middle\">a</text>\n" +
			"\t</g>\n"
		writeAnnotations(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})

	It("should write nothing when no annotations supplied", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Heatmap{
			Annotations: types.Annotations{
				FontSize: 15,
				List:     map[string]types.Annotation{},
			},
			LeftMargin: 50,
			PlotWidth:  100,
			PlotHeight: 200,
			TopMargin:  50,
		}

		expected := ""
		writeAnnotations(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})

	It("should write nothing when annotations are nil", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Heatmap{
			LeftMargin: 50,
			PlotWidth:  100,
			PlotHeight: 200,
			TopMargin:  50,
		}

		expected := ""
		writeAnnotations(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
