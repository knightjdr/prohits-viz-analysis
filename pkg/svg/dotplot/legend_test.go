package dotplot

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Draw dotplot legend", func() {
	It("should draw legend to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := Legend{
			Filename:  "legend.svg",
			NumColors: 6,
			Settings: types.Settings{
				AbundanceCap:    50,
				EdgeColor:       "blue",
				FillColor:       "blue",
				MinAbundance:    0,
				PrimaryFilter:   0.01,
				Score:           "BFDR",
				ScoreType:       "lte",
				SecondaryFilter: 0.05,
			},
			Title: "dotplot title",
		}

		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"200\" height=\"240\" viewBox=\"0 0 200 240\">\n" +
			"\t<rect width=\"100%\" height=\"100%\" fill=\"white\" />\n" +
			"\t<text y=\"20\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">dotplot title</text>\n" +
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
			"\t<g>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"100\" cx=\"60\" r=\"6\" />\n" +
			"\t\t<circle fill=\"#000000\" cy=\"100\" cx=\"135\" r=\"12\" />\n" +
			"\t\t<line fill=\"none\" stroke=\"#000000\" stroke-width=\"1\" x1=\"70\" y1=\"100\" x2=\"119\" y2=\"100\"/>\n" +
			"\t\t<polygon fill=\"#000000\" points=\"110,96 112,100 110,104 119,100\"/>\n" +
			"\t\t<text y=\"130\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">Relative abundance</text>\n" +
			"\t</g>\n" +
			"\t<g>\n" +
			"\t\t<text y=\"220\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">BFDR</text>\n" +
			"\t\t<circle fill=\"none\" cy=\"165\" cx=\"50\" r=\"12\" stroke=\"#000000\" stroke-width=\"2\" />\n" +
			"\t\t<text y=\"195\" x=\"50\" font-size=\"12\" text-anchor=\"middle\">&#8804; 0.01</text>\n" +
			"\t\t<circle fill=\"none\" cy=\"165\" cx=\"100\" r=\"12\" stroke=\"#3366ff\" stroke-width=\"2\" />\n" +
			"\t\t<text y=\"195\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">&#8804; 0.05</text>\n" +
			"\t\t<circle fill=\"none\" cy=\"165\" cx=\"150\" r=\"12\" stroke=\"#99b3ff\" stroke-width=\"2\" />\n" +
			"\t\t<text y=\"195\" x=\"150\" font-size=\"12\" text-anchor=\"middle\">&gt; 0.05</text>\n" +
			"\t</g>\n" +
			"</svg>\n"

		CreateLegend(data)
		actual, _ := afero.ReadFile(fs.Instance, "legend.svg")
		Expect(string(actual)).To(Equal(expected))
	})
})

var _ = Describe("Legend abundance graphic", func() {
	It("should create graphic", func() {
		expected := "\t<g>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"100\" cx=\"60\" r=\"6\" />\n" +
			"\t\t<circle fill=\"#000000\" cy=\"100\" cx=\"135\" r=\"12\" />\n" +
			"\t\t<line fill=\"none\" stroke=\"#000000\" stroke-width=\"1\" x1=\"70\" y1=\"100\" x2=\"119\" y2=\"100\"/>\n" +
			"\t\t<polygon fill=\"#000000\" points=\"110,96 112,100 110,104 119,100\"/>\n" +
			"\t\t<text y=\"130\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">Relative abundance</text>\n" +
			"\t</g>\n"

		var svg strings.Builder
		createAbundanceElement(&svg)
		Expect(svg.String()).To(Equal(expected))
	})
})

var _ = Describe("Legend score graphic", func() {
	It("should create graphic", func() {
		gradientData := &Dotplot{
			EdgeColor: "blue",
			FillColor: "blue",
			NumColors: 6,
		}
		_, edgeGradient := createGradients(gradientData)

		data := Legend{
			NumColors: 6,
			Settings: types.Settings{
				PrimaryFilter:   0.01,
				Score:           "BFDR",
				ScoreType:       "lte",
				SecondaryFilter: 0.05,
			},
		}

		expected := "\t<g>\n" +
			"\t\t<text y=\"220\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">BFDR</text>\n" +
			"\t\t<circle fill=\"none\" cy=\"165\" cx=\"50\" r=\"12\" stroke=\"#000000\" stroke-width=\"2\" />\n" +
			"\t\t<text y=\"195\" x=\"50\" font-size=\"12\" text-anchor=\"middle\">&#8804; 0.01</text>\n" +
			"\t\t<circle fill=\"none\" cy=\"165\" cx=\"100\" r=\"12\" stroke=\"#3366ff\" stroke-width=\"2\" />\n" +
			"\t\t<text y=\"195\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">&#8804; 0.05</text>\n" +
			"\t\t<circle fill=\"none\" cy=\"165\" cx=\"150\" r=\"12\" stroke=\"#99b3ff\" stroke-width=\"2\" />\n" +
			"\t\t<text y=\"195\" x=\"150\" font-size=\"12\" text-anchor=\"middle\">&gt; 0.05</text>\n" +
			"\t</g>\n"

		var svg strings.Builder
		createScoreElement(&svg, data, edgeGradient)
		Expect(svg.String()).To(Equal(expected))
	})
})

var _ = Describe("Create legend score symbols", func() {
	It("should create a slice of symbols to use when score type is \"gte\"", func() {
		expected := []string{"&#8805;", "&#8805;", "&lt;"}

		Expect(createLegendScoreSymbol("gte")).To(Equal(expected))
	})

	It("should create a slice of symbols to use when score type is \"lte\"", func() {
		expected := []string{"&#8804;", "&#8804;", "&gt;"}

		Expect(createLegendScoreSymbol("lte")).To(Equal(expected))
	})
})
