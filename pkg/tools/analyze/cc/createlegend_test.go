package cc

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Draw condition-condition legend", func() {
	It("should draw legend to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("svg", 0755)

		settings := types.Settings{
			Condition:       "Bait",
			PrimaryFilter:   0.01,
			Score:           "FDR",
			ScoreType:       "lte",
			SecondaryFilter: 0.05,
		}

		expectedSVG := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"200\" height=\"100\" viewBox=\"0 0 200 100\">\n" +
			"\t<text y=\"20\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">Bait-Bait</text>\n" +
			"\t<circle cx=\"20\" cy=\"50\" fill=\"#0066cc\" r=\"6\" />\n" +
			"\t<text font-size=\"12\" x=\"35\" y=\"54\">FDR &#8804; 0.01</text>\n" +
			"\t<circle cx=\"20\" cy=\"80\" fill=\"#99ccff\" r=\"6\" />\n" +
			"\t<text font-size=\"12\" x=\"35\" y=\"84\">0.01 &lt; FDR &#8804; 0.05</text>\n" +
			"</svg>\n"

		createLegend(settings)
		actualSVG, _ := afero.ReadFile(fs.Instance, "svg/Bait-Bait-legend.svg")
		Expect(string(actualSVG)).To(Equal(expectedSVG))
	})
})

var _ = Describe("Create legend score symbols", func() {
	It("should create a slice of symbols to use when score type is \"gte\"", func() {
		expected := []string{"≥", ">", "≥"}

		Expect(createLegendScoreSymbol("gte", false)).To(Equal(expected))
	})

	It("should create a slice of symbols to use when score type is \"gte\" using html entities", func() {
		expected := []string{"&#8805;", "&gt;", "&#8805;"}

		Expect(createLegendScoreSymbol("gte", true)).To(Equal(expected))
	})

	It("should create a slice of symbols to use when score type is \"lte\"", func() {
		expected := []string{"≤", "<", "≤"}

		Expect(createLegendScoreSymbol("lte", false)).To(Equal(expected))
	})

	It("should create a slice of symbols to use when score type is \"lte\" using html entities", func() {
		expected := []string{"&#8804;", "&lt;", "&#8804;"}

		Expect(createLegendScoreSymbol("lte", true)).To(Equal(expected))
	})
})
