package scatter

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Draw", func() {
	It("should draw a scatter plot to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		scatter := Initialize()
		scatter.LogBase = "none"
		scatter.Plot = []types.ScatterPoint{
			{Color: "#ff0000", Label: "point1", X: 10, Y: 30},
			{Color: "#00ff00", Label: "point2", X: 17, Y: 10},
		}
		scatter.PlotSize = 250
		scatter.XLabel = "x-axis"
		scatter.YLabel = "y-axis"

		scatter.Draw("test/scatter.svg")

		actual, _ := afero.ReadFile(fs.Instance, "test/scatter.svg")
		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"" +
			" xml:space=\"preserve\" width=\"250\" height=\"250\" viewBox=\"0 0 250 250\">\n" +
			"\t<rect width=\"100%\" height=\"100%\" fill=\"white\" />\n" +
			"\t<g transform=\"translate(0 50)\">\n" +
			"\t\t<g transform=\"translate(100 0)\">\n" +
			"\t\t\t<circle cx=\"50\" cy=\"0\" fill=\"#ff0000\" r=\"4\"><title>point1</title></circle>\n" +
			"\t\t\t<circle cx=\"85\" cy=\"66.67\" fill=\"#00ff00\" r=\"4\"><title>point2</title></circle>\n" +
			"\t\t</g>\n" +
			"\t\t<g>\n" +
			"\t\t\t<g transform=\"translate(100 0)\">\n" +
			"\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"0\" x2=\"100\" y1=\"100\" y2=\"100\"/>\n" +
			"\t\t\t\t<g>\n" +
			"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"50\" x2=\"50\" y1=\"100\" y2=\"110\" />\n" +
			"\t\t\t\t\t<text font-size=\"12\" text-anchor=\"middle\" x=\"50\" y=\"128\">10</text>\n" +
			"\t\t\t\t</g>\n" +
			"\t\t\t\t<g>\n" +
			"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"100\" x2=\"100\" y1=\"100\" y2=\"110\" />\n" +
			"\t\t\t\t\t<text font-size=\"12\" text-anchor=\"middle\" x=\"100\" y=\"128\">20</text>\n" +
			"\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"translate(0 100)\">\n" +
			"\t\t\t\t\t<text text-anchor=\"middle\" x=\"50\" y=\"70\">x-axis</text>\n" +
			"\t\t\t\t</g>\n" +
			"\t\t\t</g>\n" +
			"\t\t\t<g transform=\"translate(100 0)\">\n" +
			"\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"0\" x2=\"0\" y1=\"100\" y2=\"0\"/>\n" +
			"\t\t\t\t<g>\n" +
			"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"0\" x2=\"-10\" y1=\"66.67\" y2=\"66.67\" />\n" +
			"\t\t\t\t\t<text dy=\"4\" font-size=\"12\" text-anchor=\"end\" x=\"-13\" y=\"66.67\">10</text>\n" +
			"\t\t\t\t</g>\n" +
			"\t\t\t\t<g>\n" +
			"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"0\" x2=\"-10\" y1=\"33.33\" y2=\"33.33\" />\n" +
			"\t\t\t\t\t<text dy=\"4\" font-size=\"12\" text-anchor=\"end\" x=\"-13\" y=\"33.33\">20</text>\n" +
			"\t\t\t\t</g>\n" +
			"\t\t\t\t<g>\n" +
			"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"0\" x2=\"-10\" y1=\"0\" y2=\"0\" />\n" +
			"\t\t\t\t\t<text dy=\"4\" font-size=\"12\" text-anchor=\"end\" x=\"-13\" y=\"0\">30</text>\n" +
			"\t\t\t\t</g>\n" +
			"\t\t\t</g>\n" +
			"\t\t\t<text dy=\"15\" text-anchor=\"middle\" transform=\"rotate(-90, 0, 50)\" x=\"0\" y=\"50\">y-axis</text>\n" +
			"\t\t</g>\n" +
			"\t</g>\n" +
			"</svg>\n"
		Expect(string(actual)).To(Equal(expected))
	})
})
