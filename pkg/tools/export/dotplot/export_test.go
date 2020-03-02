package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var jsonText = `{
	"annotations": {
		"fontSize": 16,
		"list": {
			"a": {
				"position": {"x": 0.5, "y": 0.25},
				"text": "a"
			}
		}
	},
	"markers": {
		"color": "#000000",
		"list": {
			"a": {
				"height": 1,
				"width": 2,
				"x": 0.25,
				"y": 0.5
			}
		}
	},
	"settings": {"abundanceCap":4,"edgeColor":"blue","fillColor":"blue","imageType":"dotplot","minAbundance":0,"primaryFilter":0.01,"secondaryFilter":0.05,"scoreType":"lte","xLabel":"Bait","yLabel":"Prey"},
	"columnDB": ["bait1", "bait2", "bait3", "bait4"],
	"columnOrder": [0, 1, 2, 3],
	"rowOrder": [0, 1, 2, 3],
	"rowDB": [
		{
			"name": "prey1",
			"data": [
				{"ratio": 0.5, "score": 0.05, "value": 1},
				{"ratio": 0.5, "score": 0.05, "value": 1},
				{"ratio": 1, "score": 0.01, "value": 2},
				{"ratio": 1, "score": 0.01, "value": 2}
			]
		},
		{
			"name": "prey2",
			"data": [
				{"ratio": 0.67, "score": 0.01, "value": 2},
				{"ratio": 0.67, "score": 0.01, "value": 2},
				{"ratio": 0.67, "score": 0.01, "value": 2},
				{"ratio": 1, "score": 0, "value": 3}
			]
		},
		{
			"name": "prey3",
			"data": [
				{"ratio": 0.25, "score": 0.1, "value": 1},
				{"ratio": 0.25, "score": 0.1, "value": 1},
				{"ratio": 1, "score": 0, "value": 4},
				{"ratio": 0.25, "score": 0.1, "value": 1}
			]
		},
		{
			"name": "prey4",
			"data": [
				{"ratio": 1, "score": 0, "value": 4},
				{"ratio": 0.5, "score": 0.05, "value": 2},
				{"ratio": 0.5, "score": 0.05, "value": 2},
				{"ratio": 0.75, "score": 0.01, "value": 3}
			]
		}
	]
}`

var _ = Describe("Export dotplot", func() {
	It("should export svg", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		afero.WriteFile(fs.Instance, "data.json", []byte(jsonText), 0644)

		settings := Settings{
			DownsampleThreshold: 5,
			Format:              "svg",
		}

		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"137\" height=\"137\" viewBox=\"0 0 137 137\">\n" +
			"\t<g transform=\"translate(57)\">\n" +
			"\t\t<text y=\"55\" x=\"6\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 6, 55)\">bait1</text>\n" +
			"\t\t<text y=\"55\" x=\"26\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 26, 55)\">bait2</text>\n" +
			"\t\t<text y=\"55\" x=\"46\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 46, 55)\">bait3</text>\n" +
			"\t\t<text y=\"55\" x=\"66\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 66, 55)\">bait4</text>\n" +
			"\t</g>\n" +
			"\t<g transform=\"translate(0, 57)\">\n" +
			"\t\t<text y=\"15\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey1</text>\n" +
			"\t\t<text y=\"35\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey2</text>\n" +
			"\t\t<text y=\"55\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey3</text>\n" +
			"\t\t<text y=\"75\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey4</text>\n" +
			"\t</g>\n" +
			"\t<g id=\"minimap\" transform=\"translate(57, 57)\">\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"10\" cx=\"10\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"10\" cx=\"30\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"10\" cx=\"50\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"10\" cx=\"70\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"30\" cx=\"10\" r=\"5.700000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"30\" cx=\"30\" r=\"5.700000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"30\" cx=\"50\" r=\"5.700000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#002080\" cy=\"30\" cx=\"70\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"50\" cx=\"10\" r=\"2.130000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"50\" cx=\"30\" r=\"2.130000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"50\" cx=\"50\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"50\" cx=\"70\" r=\"2.130000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"70\" cx=\"10\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"70\" cx=\"30\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"70\" cx=\"50\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#002080\" cy=\"70\" cx=\"70\" r=\"6.380000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t</g>\n" +
			"\t<rect fill=\"none\" y=\"57\" x=\"57\" width=\"80\" height=\"80\" stroke=\"#000000\" stroke-width=\"0.5\" />\n" +
			"\t<g transform=\"translate(57, 57)\">\n" +
			"\t\t<rect y=\"40\" x=\"20\" width=\"40\" height=\"20\" stroke=\"#000000\" stroke-width=\"1\" fill=\"none\"/>\n" +
			"\t</g>\n" +
			"\t<g transform=\"translate(57, 57)\">\n" +
			"\t\t<text y=\"20\" x=\"40\" font-size=\"16\" text-anchor=\"middle\">a</text>\n" +
			"\t</g>\n" +
			"\t<text y=\"10\" x=\"97\" font-size=\"12\" text-anchor=\"middle\">Bait</text>\n" +
			"\t<text y=\"97\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 97)\">Prey</text>\n" +
			"</svg>\n"

		Export("data.json", settings)
		svgContent, _ := afero.ReadFile(fs.Instance, "svg/dotplot.svg")
		Expect(string(svgContent)).To(Equal(expected))
	})

	It("should export png via svg", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		afero.WriteFile(fs.Instance, "data.json", []byte(jsonText), 0644)

		oldConvert := convertSVG
		convertSVG = func(string, string, string) {}
		defer func() { convertSVG = oldConvert }()

		settings := Settings{
			DownsampleThreshold: 5,
			Format:              "png",
		}

		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"137\" height=\"137\" viewBox=\"0 0 137 137\">\n" +
			"\t<g transform=\"translate(57)\">\n" +
			"\t\t<text y=\"55\" x=\"6\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 6, 55)\">bait1</text>\n" +
			"\t\t<text y=\"55\" x=\"26\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 26, 55)\">bait2</text>\n" +
			"\t\t<text y=\"55\" x=\"46\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 46, 55)\">bait3</text>\n" +
			"\t\t<text y=\"55\" x=\"66\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 66, 55)\">bait4</text>\n" +
			"\t</g>\n" +
			"\t<g transform=\"translate(0, 57)\">\n" +
			"\t\t<text y=\"15\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey1</text>\n" +
			"\t\t<text y=\"35\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey2</text>\n" +
			"\t\t<text y=\"55\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey3</text>\n" +
			"\t\t<text y=\"75\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey4</text>\n" +
			"\t</g>\n" +
			"\t<g id=\"minimap\" transform=\"translate(57, 57)\">\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"10\" cx=\"10\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"10\" cx=\"30\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"10\" cx=\"50\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"10\" cx=\"70\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"30\" cx=\"10\" r=\"5.700000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"30\" cx=\"30\" r=\"5.700000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"30\" cx=\"50\" r=\"5.700000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#002080\" cy=\"30\" cx=\"70\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"50\" cx=\"10\" r=\"2.130000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"50\" cx=\"30\" r=\"2.130000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"50\" cx=\"50\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"50\" cx=\"70\" r=\"2.130000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"70\" cx=\"10\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"70\" cx=\"30\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"70\" cx=\"50\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#002080\" cy=\"70\" cx=\"70\" r=\"6.380000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t</g>\n" +
			"\t<rect fill=\"none\" y=\"57\" x=\"57\" width=\"80\" height=\"80\" stroke=\"#000000\" stroke-width=\"0.5\" />\n" +
			"\t<g transform=\"translate(57, 57)\">\n" +
			"\t\t<rect y=\"40\" x=\"20\" width=\"40\" height=\"20\" stroke=\"#000000\" stroke-width=\"1\" fill=\"none\"/>\n" +
			"\t</g>\n" +
			"\t<g transform=\"translate(57, 57)\">\n" +
			"\t\t<text y=\"20\" x=\"40\" font-size=\"16\" text-anchor=\"middle\">a</text>\n" +
			"\t</g>\n" +
			"\t<text y=\"10\" x=\"97\" font-size=\"12\" text-anchor=\"middle\">Bait</text>\n" +
			"\t<text y=\"97\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 97)\">Prey</text>\n" +
			"</svg>\n"

		Export("data.json", settings)
		svgContent, _ := afero.ReadFile(fs.Instance, "svg/dotplot.svg")
		Expect(string(svgContent)).To(Equal(expected))
	})
})
