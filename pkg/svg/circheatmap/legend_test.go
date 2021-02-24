package circheatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Draw circheatmap legend", func() {
	It("should draw legend to file with knownness indicator", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := Legend{
			Elements: types.CircHeatmapLegend{
				{Attribute: "element1", Color: "blue", Max: 50, Min: 0},
				{Attribute: "element2", Color: "red", Max: 25, Min: 10},
			},
			Filename: "legend.svg",
			Known:    "interaction",
			Title:    "circheatmap title",
		}

		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"210\" height=\"180\" viewBox=\"0 0 210 180\">\n" +
			"\t<rect width=\"100%\" height=\"100%\" fill=\"white\" />\n" +
			"\t<text y=\"20\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">circheatmap title</text>\n" +
			"\t<g transform=\"translate(0 30)\">\n" +
			"\t\t<g transform=\"translate(0 0)\">\n" +
			"\t\t\t<defs>\n" +
			"\t\t\t\t<linearGradient id=\"element1-legendGradient\">\n" +
			"\t\t\t\t\t<stop offset=\"0%\" stop-color=\"#ffffff\" />\n" +
			"\t\t\t\t\t<stop offset=\"50%\" stop-color=\"#0040ff\" />\n" +
			"\t\t\t\t\t<stop offset=\"100%\" stop-color=\"#000000\" />\n" +
			"\t\t\t\t</linearGradient>\n" +
			"\t\t\t</defs>\n" +
			"\t\t\t<g>\n" +
			"\t\t\t\t<text x=\"100\" y=\"20\" text-anchor=\"middle\">element1</text>\n" +
			"\t\t\t\t<rect x=\"25\" y=\"30\" height=\"20\" width=\"150\" fill=\"url(#element1-legendGradient)\" />\n" +
			"\t\t\t\t<text x=\"20\" y=\"45\" text-anchor=\"end\">0</text>\n" +
			"\t\t\t\t<text x=\"180\" y=\"45\" text-anchor=\"start\">50</text>\n" +
			"\t\t\t</g>\n" +
			"\t\t</g>\n" +
			"\t\t<g transform=\"translate(0 50)\">\n" +
			"\t\t\t<defs>\n" +
			"\t\t\t\t<linearGradient id=\"element2-legendGradient\">\n" +
			"\t\t\t\t\t<stop offset=\"0%\" stop-color=\"#ffffff\" />\n" +
			"\t\t\t\t\t<stop offset=\"50%\" stop-color=\"#ff0000\" />\n" +
			"\t\t\t\t\t<stop offset=\"100%\" stop-color=\"#000000\" />\n" +
			"\t\t\t\t</linearGradient>\n" +
			"\t\t\t</defs>\n" +
			"\t\t\t<g>\n" +
			"\t\t\t\t<text x=\"100\" y=\"20\" text-anchor=\"middle\">element2</text>\n" +
			"\t\t\t\t<rect x=\"25\" y=\"30\" height=\"20\" width=\"150\" fill=\"url(#element2-legendGradient)\" />\n" +
			"\t\t\t\t<text x=\"20\" y=\"45\" text-anchor=\"end\">10</text>\n" +
			"\t\t\t\t<text x=\"180\" y=\"45\" text-anchor=\"start\">25</text>\n" +
			"\t\t\t</g>\n" +
			"\t\t</g>\n" +
			"\t\t<g transform=\"translate(0 120)\">\n" +
			"\t\t\t<text text-anchor=\"middle\" x=\"100\" y=\"0\">Known interaction</text>\n" +
			"\t\t\t<line stroke=\"black\" stroke-width=\"3\" x1=\"50\" x2=\"150\" y1=\"10\" y2=\"10\"/>\n" +
			"\t\t</g>\n" +
			"\t</g>\n" +
			"</svg>\n"

		CreateLegend(data)
		actual, _ := afero.ReadFile(fs.Instance, "legend.svg")
		Expect(string(actual)).To(Equal(expected))
	})

	It("should draw legend to file without knownness indicator", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := Legend{
			Elements: types.CircHeatmapLegend{
				{Attribute: "element1", Color: "blue", Max: 50, Min: 0},
				{Attribute: "element2", Color: "red", Max: 25, Min: 10},
			},
			Filename: "legend.svg",
			Known:    "",
			Title:    "circheatmap title",
		}

		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"210\" height=\"140\" viewBox=\"0 0 210 140\">\n" +
			"\t<rect width=\"100%\" height=\"100%\" fill=\"white\" />\n" +
			"\t<text y=\"20\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">circheatmap title</text>\n" +
			"\t<g transform=\"translate(0 30)\">\n" +
			"\t\t<g transform=\"translate(0 0)\">\n" +
			"\t\t\t<defs>\n" +
			"\t\t\t\t<linearGradient id=\"element1-legendGradient\">\n" +
			"\t\t\t\t\t<stop offset=\"0%\" stop-color=\"#ffffff\" />\n" +
			"\t\t\t\t\t<stop offset=\"50%\" stop-color=\"#0040ff\" />\n" +
			"\t\t\t\t\t<stop offset=\"100%\" stop-color=\"#000000\" />\n" +
			"\t\t\t\t</linearGradient>\n" +
			"\t\t\t</defs>\n" +
			"\t\t\t<g>\n" +
			"\t\t\t\t<text x=\"100\" y=\"20\" text-anchor=\"middle\">element1</text>\n" +
			"\t\t\t\t<rect x=\"25\" y=\"30\" height=\"20\" width=\"150\" fill=\"url(#element1-legendGradient)\" />\n" +
			"\t\t\t\t<text x=\"20\" y=\"45\" text-anchor=\"end\">0</text>\n" +
			"\t\t\t\t<text x=\"180\" y=\"45\" text-anchor=\"start\">50</text>\n" +
			"\t\t\t</g>\n" +
			"\t\t</g>\n" +
			"\t\t<g transform=\"translate(0 50)\">\n" +
			"\t\t\t<defs>\n" +
			"\t\t\t\t<linearGradient id=\"element2-legendGradient\">\n" +
			"\t\t\t\t\t<stop offset=\"0%\" stop-color=\"#ffffff\" />\n" +
			"\t\t\t\t\t<stop offset=\"50%\" stop-color=\"#ff0000\" />\n" +
			"\t\t\t\t\t<stop offset=\"100%\" stop-color=\"#000000\" />\n" +
			"\t\t\t\t</linearGradient>\n" +
			"\t\t\t</defs>\n" +
			"\t\t\t<g>\n" +
			"\t\t\t\t<text x=\"100\" y=\"20\" text-anchor=\"middle\">element2</text>\n" +
			"\t\t\t\t<rect x=\"25\" y=\"30\" height=\"20\" width=\"150\" fill=\"url(#element2-legendGradient)\" />\n" +
			"\t\t\t\t<text x=\"20\" y=\"45\" text-anchor=\"end\">10</text>\n" +
			"\t\t\t\t<text x=\"180\" y=\"45\" text-anchor=\"start\">25</text>\n" +
			"\t\t\t</g>\n" +
			"\t\t</g>\n" +
			"\t</g>\n" +
			"</svg>\n"

		CreateLegend(data)
		actual, _ := afero.ReadFile(fs.Instance, "legend.svg")
		Expect(string(actual)).To(Equal(expected))
	})
})
