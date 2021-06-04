package circheatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Draw circular heatmap", func() {
	It("should draw it with known indicator", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		c := Initialize()
		c.Legend = types.CircHeatmapLegend{
			{Attribute: "foldchange", Color: "red", Max: 5, Min: 0},
			{Attribute: "abundance", Color: "blue", Max: 4, Min: 0},
		}
		c.Plot = types.CircHeatmap{
			Name: "conditionA",
			Readouts: []types.CircHeatmapReadout{
				{
					Known: true,
					Label: "readoutX",
					Segments: map[string]types.RoundedSegment{
						"abundance":  3,
						"foldchange": 4,
					},
				},
				{
					Known: true,
					Label: "readoutA",
					Segments: map[string]types.RoundedSegment{
						"abundance":  2,
						"foldchange": 3,
					},
				},
				{
					Known: false,
					Label: "readoutY",
					Segments: map[string]types.RoundedSegment{
						"abundance":  4,
						"foldchange": 5,
					},
				},
				{
					Known: false,
					Label: "readoutB",
					Segments: map[string]types.RoundedSegment{
						"abundance":  1,
						"foldchange": 2,
					},
				},
			},
		}
		c.ShowKnown = true

		c.Draw("test/circheatmap.svg")

		actual, _ := afero.ReadFile(fs.Instance, "test/circheatmap.svg")
		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"750\" height=\"750\" viewBox=\"-375 -375 750 750\">\n" +
			"\t<rect width=\"100%\" height=\"100%\" fill=\"white\" transform=\"translate(-375 -375)\"/>\n" +
			"\t<g transform=\"rotate(-90)\">\n" +
			"\t\t<path d=\"M 360.00 0 A 360.00 360.00 0 0 1 -360.00 0.00\" fill=\"none\" stroke=\"#333\" stroke-width=\"5\" transform=\"scale(0.9 0.9)\"/>\n" +
			"\t\t<g>\n" +
			"\t\t\t<g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M 360 0 A 360 360 0 0 1 0 360 L 0 253 A 253 253 0 0 0 253 0 Z\" fill=\"#660000\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M 0 360 A 360 360 0 0 1 -360 0 L -253 0 A 253 253 0 0 0 0 253 Z\" fill=\"#cc0000\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M -360 0 A 360 360 0 0 1 -0 -360 L -0 -253 A 253 253 0 0 0 -253 0 Z\" fill=\"#000000\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M -0 -360 A 360 360 0 0 1 360 -0 L 253 -0 A 253 253 0 0 0 -0 -253 Z\" fill=\"#ff3333\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t</g>\n" +
			"\t\t\t<g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M 226.25 0 A 226.25 226.25 0 0 1 0 226.25 L 0 119 A 119 119 0 0 0 119 0 Z\" fill=\"#002080\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M 0 226.25 A 226.25 226.25 0 0 1 -226.25 0 L -119 0 A 119 119 0 0 0 0 119 Z\" fill=\"#0040ff\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M -226.25 0 A 226.25 226.25 0 0 1 -0 -226.25 L -0 -119 A 119 119 0 0 0 -119 0 Z\" fill=\"#000000\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M -0 -226.25 A 226.25 226.25 0 0 1 226.25 -0 L 119 -0 A 119 119 0 0 0 -0 -119 Z\" fill=\"#809fff\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t</g>\n" +
			"\t\t</g>\n" +
			"\t</g>\n" +
			"</svg>\n"

		Expect(string(actual)).To(Equal(expected))
	})

	It("should draw it without known indicator", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		c := Initialize()
		c.Legend = types.CircHeatmapLegend{
			{Attribute: "foldchange", Color: "red", Max: 5, Min: 0},
			{Attribute: "abundance", Color: "blue", Max: 4, Min: 0},
		}
		c.Plot = types.CircHeatmap{
			Name: "conditionA",
			Readouts: []types.CircHeatmapReadout{
				{
					Known: false,
					Label: "readoutY",
					Segments: map[string]types.RoundedSegment{
						"abundance":  4,
						"foldchange": 5,
					},
				},
				{
					Known: false,
					Label: "readoutX",
					Segments: map[string]types.RoundedSegment{
						"abundance":  3,
						"foldchange": 4,
					},
				},
				{
					Known: false,
					Label: "readoutA",
					Segments: map[string]types.RoundedSegment{
						"abundance":  2,
						"foldchange": 3,
					},
				},
				{
					Known: false,
					Label: "readoutB",
					Segments: map[string]types.RoundedSegment{
						"abundance":  1,
						"foldchange": 2,
					},
				},
			},
		}
		c.ShowKnown = false

		c.Draw("test/circheatmap.svg")

		actual, _ := afero.ReadFile(fs.Instance, "test/circheatmap.svg")
		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"750\" height=\"750\" viewBox=\"-375 -375 750 750\">\n" +
			"\t<rect width=\"100%\" height=\"100%\" fill=\"white\" transform=\"translate(-375 -375)\"/>\n" +
			"\t<g transform=\"rotate(-90)\">\n" +
			"\t\t<g>\n" +
			"\t\t\t<g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M 360 0 A 360 360 0 0 1 0 360 L 0 253 A 253 253 0 0 0 253 0 Z\" fill=\"#000000\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M 0 360 A 360 360 0 0 1 -360 0 L -253 0 A 253 253 0 0 0 0 253 Z\" fill=\"#660000\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M -360 0 A 360 360 0 0 1 -0 -360 L -0 -253 A 253 253 0 0 0 -253 0 Z\" fill=\"#cc0000\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M -0 -360 A 360 360 0 0 1 360 -0 L 253 -0 A 253 253 0 0 0 -0 -253 Z\" fill=\"#ff3333\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t</g>\n" +
			"\t\t\t<g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M 226.25 0 A 226.25 226.25 0 0 1 0 226.25 L 0 119 A 119 119 0 0 0 119 0 Z\" fill=\"#000000\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M 0 226.25 A 226.25 226.25 0 0 1 -226.25 0 L -119 0 A 119 119 0 0 0 0 119 Z\" fill=\"#002080\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M -226.25 0 A 226.25 226.25 0 0 1 -0 -226.25 L -0 -119 A 119 119 0 0 0 -119 0 Z\" fill=\"#0040ff\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"scale(0.85)\">\n\t\t\t\t\t<path d=\"M -0 -226.25 A 226.25 226.25 0 0 1 226.25 -0 L 119 -0 A 119 119 0 0 0 -0 -119 Z\" fill=\"#809fff\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n\t\t\t\t</g>\n" +
			"\t\t\t</g>\n" +
			"\t\t</g>\n" +
			"\t</g>\n" +
			"</svg>\n"

		Expect(string(actual)).To(Equal(expected))
	})
})
