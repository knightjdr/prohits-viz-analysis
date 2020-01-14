package heatmap

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Draw heatmap legend", func() {
	It("should draw legend to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := Legend{
			Filename:  "legend.svg",
			NumColors: 6,
			Settings: types.Settings{
				AbundanceCap: 50,
				FillColor:    "blue",
				MinAbundance: 0,
			},
			Title: "heatmap title",
		}

		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"200\" height=\"240\" viewBox=\"0 0 200 240\">\n" +
			"\t<text y=\"20\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">heatmap title</text>\n" +
			"\t<g>\n" +
			"\t\t<rect fill=\"#ffffff\" y=\"30\" x=\"25\" width=\"25\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#99b3ff\" y=\"30\" x=\"50\" width=\"25\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#3366ff\" y=\"30\" x=\"75\" width=\"25\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0033cc\" y=\"30\" x=\"100\" width=\"25\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#001966\" y=\"30\" x=\"125\" width=\"25\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#000000\" y=\"30\" x=\"150\" width=\"25\" height=\"20\" />\n" +
			"\t</g>\n" +
			"\t<rect fill=\"none\" y=\"30\" x=\"25\" width=\"150\" height=\"20\" stroke=\"#000000\" stroke-width=\"1\"/>\n" +
			"\t<text y=\"65\" x=\"175\" font-size=\"12\" text-anchor=\"middle\">50</text>\n" +
			"\t<text y=\"65\" x=\"25\" font-size=\"12\" text-anchor=\"middle\">0</text>\n" +
			"</svg>\n"

		CreateLegend(data)
		actual, _ := afero.ReadFile(fs.Instance, "legend.svg")
		Expect(string(actual)).To(Equal(expected))
	})
})

var _ = Describe("Legend header", func() {
	It("should create header", func() {
		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"200\" height=\"240\" viewBox=\"0 0 200 240\">\n"

		var svg strings.Builder
		CreateLegendHeader(&svg)
		Expect(svg.String()).To(Equal(expected))
	})
})

var _ = Describe("Legend title", func() {
	It("should create title", func() {
		title := "dotplot title"

		expected := "\t<text y=\"20\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">dotplot title</text>\n"

		var svg strings.Builder
		CreateLegendTitle(&svg, title)
		Expect(svg.String()).To(Equal(expected))
	})
})

var _ = Describe("Legend fill gradient", func() {
	It("should create fill gradient", func() {
		gradientData := &Heatmap{
			FillColor: "blue",
			NumColors: 6,
		}
		fillGradient := createGradient(gradientData)

		data := Legend{
			NumColors: 6,
			Settings: types.Settings{
				AbundanceCap: 50,
				MinAbundance: 0,
			},
		}

		expected := "\t<g>\n" +
			"\t\t<rect fill=\"#ffffff\" y=\"30\" x=\"25\" width=\"25\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#99b3ff\" y=\"30\" x=\"50\" width=\"25\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#3366ff\" y=\"30\" x=\"75\" width=\"25\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0033cc\" y=\"30\" x=\"100\" width=\"25\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#001966\" y=\"30\" x=\"125\" width=\"25\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#000000\" y=\"30\" x=\"150\" width=\"25\" height=\"20\" />\n" +
			"\t</g>\n" +
			"\t<rect fill=\"none\" y=\"30\" x=\"25\" width=\"150\" height=\"20\" stroke=\"#000000\" stroke-width=\"1\"/>\n" +
			"\t<text y=\"65\" x=\"175\" font-size=\"12\" text-anchor=\"middle\">50</text>\n" +
			"\t<text y=\"65\" x=\"25\" font-size=\"12\" text-anchor=\"middle\">0</text>\n"

		var svg strings.Builder
		CreateFillGradient(&svg, data, fillGradient)
		Expect(svg.String()).To(Equal(expected))
	})
})
