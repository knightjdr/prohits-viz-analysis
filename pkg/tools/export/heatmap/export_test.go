package heatmap_test

import (
	"encoding/base64"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	. "github.com/knightjdr/prohits-viz-analysis/pkg/tools/export/heatmap"
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
	"settings": {"abundanceCap":4,"abundanceType":"positive","fillColor":"blue","imageType":"heatmap","minAbundance":0,"xLabel":"Bait","yLabel":"Prey"},
	"columnDB": ["bait1", "bait2", "bait3", "bait4"],
	"columnOrder": [0, 1, 2, 3],
	"rowOrder": [0, 1, 2, 3],
	"rowDB": [
		{
			"name": "prey1",
			"data": [
				{"value": 1},
				{"value": 1},
				{"value": 2},
				{"value": 2}
			]
		},
		{
			"name": "prey2",
			"data": [
				{"value": 2},
				{"value": 2},
				{"value": 2},
				{"value": 3}
			]
		},
		{
			"name": "prey3",
			"data": [
				{"value": 1},
				{"value": 1},
				{"value": 4},
				{"value": 1}
			]
		},
		{
			"name": "prey4",
			"data": [
				{"value": 4},
				{"value": 2},
				{"value": 2},
				{"value": 3}
			]
		}
	]
}`

var _ = Describe("Export heatmap", func() {
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
			"\t\t<rect fill=\"#809fff\" y=\"0\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#809fff\" y=\"0\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"0\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"0\" x=\"60\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"20\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"20\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"20\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#002080\" y=\"20\" x=\"60\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#809fff\" y=\"40\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#809fff\" y=\"40\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#000000\" y=\"40\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#809fff\" y=\"40\" x=\"60\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#000000\" y=\"60\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"60\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"60\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#002080\" y=\"60\" x=\"60\" width=\"20\" height=\"20\" />\n" +
			"\t</g>\n" +
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
		svgContent, _ := afero.ReadFile(fs.Instance, "svg/heatmap.svg")
		Expect(string(svgContent)).To(Equal(expected))
	})

	It("should export png", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		afero.WriteFile(fs.Instance, "data.json", []byte(jsonText), 0644)

		settings := Settings{
			DownsampleThreshold: 5,
			Format:              "png",
		}

		expected := "iVBORw0KGgoAAAANSUhEUgAAAFAAAABQCAIAAAABc2X6AAAA80lEQVR4nOzawYmEMBhHcV0s" +
			"ZEuxk2U7yVayTCeWMp04R2VOQwLfH17eO3kJ4TdxIkq29n8uidojMu3ylZk2l2B6gukJpieYnmB6gukJpieY" +
			"nmB667JnvmkN9fzrHjrdCm+BOY/1ui6/vxLgO/Ko/k+lV7i8cvDbkpbjo5tWYqnLV3g/L+f9uqr0plW+S0/3" +
			"HBZMTzA9wfQE0xNMb2s//YPbb+zNtvt42ejLQ/fEqR9rultaMD3B9ATTE0xPMD3B9ATTE0xvOnDyuMVQ361v" +
			"3Ng3rdShNg+mfZ5geoLpCaYnmJ5geoLpCaYnmN504FcAAAD//0+DILFPvR1fAAAAAElFTkSuQmCC"

		Export("data.json", settings)
		pngContent, _ := afero.ReadFile(fs.Instance, "png/heatmap.png")
		Expect(base64.StdEncoding.EncodeToString((pngContent))).To(Equal(expected))
	})
})
