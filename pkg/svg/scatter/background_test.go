package scatter

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write background", func() {
	It("should write rect for white background", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		expected := "\t<rect width=\"100%\" height=\"100%\" fill=\"white\" />\n"
		writeBackground(writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
