package circheatmap

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

		c := &CircHeatmapSVG{
			Dimensions: CircHeatmapDimensions{
				Center: 500,
			},
		}

		expected := "\t<rect width=\"100%\" height=\"100%\" fill=\"white\" transform=\"translate(-500 -500)\"/>\n"
		writeBackground(c, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
