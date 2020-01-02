package heatmap

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write headings", func() {
	It("should write column and row headings", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Heatmap{
			LeftMargin: 50,
			SvgHeight:  250,
			SvgWidth:   150,
			TopMargin:  50,
			XLabel:     "Bait",
			YLabel:     "Prey",
		}

		expected := "\t<text y=\"10\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">Bait</text>\n" +
			"\t<text y=\"150\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 150)\">Prey</text>\n"
		writeHeadings(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})

	It("should write column heading", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Heatmap{
			SvgWidth: 150,
			XLabel:   "Bait",
		}

		expected := "\t<text y=\"10\" x=\"75\" font-size=\"12\" text-anchor=\"middle\">Bait</text>\n"
		writeHeadings(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})

	It("should write row heading", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Heatmap{
			SvgHeight: 250,
			YLabel:    "Prey",
		}

		expected := "\t<text y=\"125\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 125)\">Prey</text>\n"
		writeHeadings(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})

	It("should write nothing", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Heatmap{}

		expected := ""
		writeHeadings(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
