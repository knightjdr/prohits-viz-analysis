package dotplot

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write bounding box", func() {
	It("should write bounding box", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		d := &Dotplot{
			BoundingBox: true,
			LeftMargin:  50,
			SvgHeight:   250,
			SvgWidth:    150,
			TopMargin:   50,
		}

		expected := "\t<rect fill=\"none\" y=\"50\" x=\"50\" width=\"100\" height=\"200\"" +
			" stroke=\"#000000\" stroke-width=\"0.5\" />\n"
		writeBoundingBox(d, writeString)
		Expect(svg.String()).To(Equal(expected))
	})

	It("should not write bounding box", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		d := &Dotplot{
			LeftMargin: 50,
			SvgHeight:  250,
			SvgWidth:   150,
			TopMargin:  50,
		}

		expected := ""
		writeBoundingBox(d, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
