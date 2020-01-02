package dotplot

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write headings", func() {
	It("should write svg tag and dimensions", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Dotplot{
			SvgHeight: 250,
			SvgWidth:  150,
		}

		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"" +
			" xml:space=\"preserve\" width=\"150\" height=\"250\" viewBox=\"0 0 150 250\">\n"
		writeHeader(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
