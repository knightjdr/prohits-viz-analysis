package heatmap

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write markers", func() {
	It("should write markers", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Heatmap{
			CellSize:   20,
			LeftMargin: 50,
			Markers: types.Markers{
				Color: "#000000",
				List: map[string]types.Marker{
					"a": {Height: 2, Width: 2, X: 0.25, Y: 0.2},
				},
			},
			PlotHeight: 100,
			PlotWidth:  200,
			TopMargin:  50,
		}

		expected := "\t<g transform=\"translate(50, 50)\">\n" +
			"\t\t<rect y=\"20\" x=\"50\" width=\"40\" height=\"40\" stroke=\"#000000\" stroke-width=\"1\" fill=\"none\"/>\n" +
			"\t</g>\n"
		writeMarkers(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})

	It("should write nothing when no markers supplied", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Heatmap{
			CellSize:   20,
			LeftMargin: 50,
			Markers: types.Markers{
				Color: "#000000",
				List:  map[string]types.Marker{},
			},
			TopMargin: 50,
		}

		expected := ""
		writeMarkers(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})

	It("should write nothing when markers are nil", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Heatmap{
			CellSize:   20,
			LeftMargin: 50,
			TopMargin:  50,
		}

		expected := ""
		writeMarkers(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
