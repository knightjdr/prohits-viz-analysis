package circheatmap

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write circles", func() {
	It("should reformat circheatmap data to individual circles", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		c := &CircHeatmapSVG{
			Dimensions: CircHeatmapDimensions{
				Radius:    200,
				Thickness: 50,
			},
			Legend: types.CircHeatmapLegend{
				{Attribute: "foldchange", Color: "red", Max: 5, Min: 0},
				{Attribute: "abundance", Color: "blue", Max: 4, Min: 0},
			},
			Plot: types.CircHeatmap{
				Name: "conditionA",
				Readouts: []types.CircHeatmapReadout{
					{
						Label: "readoutY",
						Segments: map[string]types.RoundedSegment{
							"abundance":  4,
							"foldchange": 5,
						},
					},
					{
						Label: "readoutX",
						Segments: map[string]types.RoundedSegment{
							"abundance":  3,
							"foldchange": 4,
						},
					},
					{
						Label: "readoutA",
						Segments: map[string]types.RoundedSegment{
							"abundance":  2,
							"foldchange": 3,
						},
					},
					{
						Label: "readoutB",
						Segments: map[string]types.RoundedSegment{
							"abundance":  1,
							"foldchange": 2,
						},
					},
				},
			},
		}

		expected := "\t\t<g>\n" +
			"\t\t\t<g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M 200 0 A 200 200 0 0 1 0 200 L 0 150 A 150 150 0 0 0 150 0 Z\" fill=\"#000000\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M 0 200 A 200 200 0 0 1 -200 0 L -150 0 A 150 150 0 0 0 0 150 Z\" fill=\"#660000\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M -200 0 A 200 200 0 0 1 -0 -200 L -0 -150 A 150 150 0 0 0 -150 0 Z\" fill=\"#cc0000\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M -0 -200 A 200 200 0 0 1 200 -0 L 150 -0 A 150 150 0 0 0 -0 -150 Z\" fill=\"#ff3333\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t</g>\n" +
			"\t\t\t<g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M 137.5 0 A 137.5 137.5 0 0 1 0 137.5 L 0 87 A 87 87 0 0 0 87 0 Z\" fill=\"#000000\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M 0 137.5 A 137.5 137.5 0 0 1 -137.5 0 L -87 0 A 87 87 0 0 0 0 87 Z\" fill=\"#002080\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M -137.5 0 A 137.5 137.5 0 0 1 -0 -137.5 L -0 -87 A 87 87 0 0 0 -87 0 Z\" fill=\"#0040ff\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M -0 -137.5 A 137.5 137.5 0 0 1 137.5 -0 L 87 -0 A 87 87 0 0 0 -0 -87 Z\" fill=\"#809fff\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t</g>\n" +
			"\t\t</g>\n"

		writeCircles(c, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})

var _ = Describe("Reformat circheatmap data", func() {
	It("should reformat circheatmap data to individual circles", func() {
		c := &CircHeatmapSVG{
			Dimensions: CircHeatmapDimensions{
				Radius:    360,
				Thickness: 50,
			},
			Legend: types.CircHeatmapLegend{
				{Attribute: "foldchange", Color: "red", Max: 5, Min: 0},
				{Attribute: "abundance", Color: "blue", Max: 4, Min: 0},
			},
			Plot: types.CircHeatmap{
				Name: "conditionA",
				Readouts: []types.CircHeatmapReadout{
					{
						Label: "readoutY",
						Segments: map[string]types.RoundedSegment{
							"abundance":  4,
							"foldchange": 5,
						},
					},
					{
						Label: "readoutX",
						Segments: map[string]types.RoundedSegment{
							"abundance":  3,
							"foldchange": 4,
						},
					},
					{
						Label: "readoutA",
						Segments: map[string]types.RoundedSegment{
							"abundance":  2,
							"foldchange": 3,
						},
					},
					{
						Label: "readoutB",
						Segments: map[string]types.RoundedSegment{
							"abundance":  1,
							"foldchange": 2,
						},
					},
				},
			},
		}

		expected := []Circle{
			{
				Attribute: "foldchange",
				Color:     "red",
				Max:       5,
				Min:       0,
				Radius:    360,
				Thickness: 50,
				Values:    []float64{5, 4, 3, 2},
			},
			{
				Attribute: "abundance",
				Color:     "blue",
				Max:       4,
				Min:       0,
				Radius:    297.5,
				Thickness: 50,
				Values:    []float64{4, 3, 2, 1},
			},
		}

		Expect(reformatCircHeatmapData(c)).To(Equal(expected))
	})
})

var _ = Describe("Calculate radii", func() {
	It("should determine the inner and outer radius for segments", func() {
		radius := float64(100)
		thickness := float64(30.5)

		expected := map[string]float64{
			"inner": 69,
			"outer": 100,
		}

		Expect(calculateRadii(radius, thickness)).To(Equal(expected))
	})
})

var _ = Describe("Create color range", func() {
	It("should return an array of hex colours", func() {
		circle := Circle{
			Attribute: "abundance",
			Color:     "blue",
			Max:       4,
			Min:       0,
			Values:    []float64{4, 3, 2, 1},
		}

		expected := []string{"#000000", "#002080", "#0040ff", "#809fff"}

		Expect(createColourRange(circle)).To(Equal(expected))
	})
})

var _ = Describe("Define segments", func() {
	It("should return an array segments", func() {
		colorGradient := []string{"#000000", "#002080", "#0040ff", "#809fff"}
		radii := map[string]float64{
			"inner": 70,
			"outer": 100,
		}

		expected := []Segment{
			{
				A:    SegmentPath{X: 100, Y: 0},
				B:    SegmentPath{Arc: 0, X: 0, Y: 100},
				C:    SegmentPath{X: 0, Y: 70},
				D:    SegmentPath{Arc: 0, X: 70, Y: 0},
				Fill: "#000000",
			},
			{
				A:    SegmentPath{X: 0, Y: 100},
				B:    SegmentPath{Arc: 0, X: -100, Y: 0},
				C:    SegmentPath{X: -70, Y: 0},
				D:    SegmentPath{Arc: 0, X: 0, Y: 70},
				Fill: "#002080",
			},
			{
				A:    SegmentPath{X: -100, Y: 0},
				B:    SegmentPath{Arc: 0, X: -0, Y: -100},
				C:    SegmentPath{X: -0, Y: -70},
				D:    SegmentPath{Arc: 0, X: -70, Y: 0},
				Fill: "#0040ff",
			},
			{
				A:    SegmentPath{X: -0, Y: -100},
				B:    SegmentPath{Arc: 0, X: 100, Y: -0},
				C:    SegmentPath{X: 70, Y: -0},
				D:    SegmentPath{Arc: 0, X: -0, Y: -70},
				Fill: "#809fff",
			},
		}

		actual := defineSegments(colorGradient, radii)
		for i, segment := range actual {
			Expect(segment.A.X).To(BeNumerically("~", expected[i].A.X, 0.1))
			Expect(segment.A.Y).To(BeNumerically("~", expected[i].A.Y, 0.1))
			Expect(segment.B.Arc).To(BeNumerically("~", expected[i].B.Arc, 0.1))
			Expect(segment.B.X).To(BeNumerically("~", expected[i].B.X, 0.1))
			Expect(segment.B.Y).To(BeNumerically("~", expected[i].B.Y, 0.1))
			Expect(segment.C.X).To(BeNumerically("~", expected[i].C.X, 0.1))
			Expect(segment.C.Y).To(BeNumerically("~", expected[i].C.Y, 0.1))
			Expect(segment.D.Arc).To(BeNumerically("~", expected[i].D.Arc, 0.1))
			Expect(segment.D.X).To(BeNumerically("~", expected[i].D.X, 0.1))
			Expect(segment.D.Y).To(BeNumerically("~", expected[i].D.Y, 0.1))
			Expect(segment.Fill).To(Equal(expected[i].Fill))
		}
	})
})

var _ = Describe("Draw segment", func() {
	It("should write a segment to buffer", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		radii := map[string]float64{
			"inner": 70,
			"outer": 100,
		}
		segment := Segment{
			A:    SegmentPath{X: 100, Y: 0},
			B:    SegmentPath{Arc: 0, X: 0, Y: 100},
			C:    SegmentPath{X: 0, Y: 70},
			D:    SegmentPath{Arc: 0, X: 70, Y: 0},
			Fill: "#000000",
		}

		expected := "\t\t\t\t<g transform=\"scale(0.85)\">\n" +
			"\t\t\t\t\t<path d=\"M 100 0 A 100 100 0 0 1 0 100 L 0 70 A 70 70 0 0 0 70 0 Z\" " +
			"fill=\"#000000\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n" +
			"\t\t\t\t</g>\n"
		drawSegment(segment, radii, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
