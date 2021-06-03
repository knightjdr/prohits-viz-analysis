package circheatmap

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write known circle", func() {
	It("should create circular element when all readouts known", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		c := &CircHeatmapSVG{
			Dimensions: CircHeatmapDimensions{
				Radius: 200,
			},
			Plot: types.CircHeatmap{
				Name: "conditionA",
				Readouts: []types.CircHeatmapReadout{
					{
						Known: true,
						Label: "readoutY",
						Segments: map[string]types.RoundedSegment{
							"abundance": 4,
						},
					},
					{
						Known: true,
						Label: "readoutX",
						Segments: map[string]types.RoundedSegment{
							"abundance": 3,
						},
					},
					{
						Known: true,
						Label: "readoutA",
						Segments: map[string]types.RoundedSegment{
							"abundance": 2,
						},
					},
					{
						Known: true,
						Label: "readoutB",
						Segments: map[string]types.RoundedSegment{
							"abundance": 1,
						},
					},
				},
			},
			ShowKnown: true,
		}

		expected := "\t\t<circle cx=\"0\" cy=\"0\" fill=\"none\" r=\"200.00\" stroke=\"#333\" stroke-width=\"5\" transform=\"scale(0.9 0.9)\"/>\n"
		writeKnown(c, writeString)
		Expect(svg.String()).To(Equal(expected))
	})

	It("should create path element when some readouts known", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		c := &CircHeatmapSVG{
			Dimensions: CircHeatmapDimensions{
				Radius: 200,
			},
			Plot: types.CircHeatmap{
				Name: "conditionA",
				Readouts: []types.CircHeatmapReadout{
					{
						Known: true,
						Label: "readoutX",
						Segments: map[string]types.RoundedSegment{
							"abundance": 3,
						},
					},
					{
						Known: true,
						Label: "readoutA",
						Segments: map[string]types.RoundedSegment{
							"abundance": 2,
						},
					},
					{
						Known: true,
						Label: "readoutB",
						Segments: map[string]types.RoundedSegment{
							"abundance": 1,
						},
					},
					{
						Known: false,
						Label: "readoutY",
						Segments: map[string]types.RoundedSegment{
							"abundance": 4,
						},
					},
				},
			},
			ShowKnown: true,
		}

		expected := "\t\t<path d=\"M 200.00 0 A 200.00 200.00 0 1 1 -0.00 -200.00\" " +
			"fill=\"none\" stroke=\"#333\" stroke-width=\"5\" transform=\"scale(0.9 0.9)\"/>\n"

		writeKnown(c, writeString)
		Expect(svg.String()).To(Equal(expected))
	})

	It("should write nothing when known circle not requested", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		c := &CircHeatmapSVG{
			Dimensions: CircHeatmapDimensions{
				Radius: 200,
			},
			Plot: types.CircHeatmap{
				Name: "conditionA",
				Readouts: []types.CircHeatmapReadout{
					{
						Known: false,
						Label: "readoutY",
						Segments: map[string]types.RoundedSegment{
							"abundance": 4,
						},
					},
					{
						Known: false,
						Label: "readoutX",
						Segments: map[string]types.RoundedSegment{
							"abundance": 3,
						},
					},
					{
						Known: false,
						Label: "readoutA",
						Segments: map[string]types.RoundedSegment{
							"abundance": 2,
						},
					},
					{
						Known: false,
						Label: "readoutB",
						Segments: map[string]types.RoundedSegment{
							"abundance": 1,
						},
					},
				},
			},
			ShowKnown: false,
		}

		expected := ""

		writeKnown(c, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
