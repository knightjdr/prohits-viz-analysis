package specificity

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Draw specificity legend", func() {
	It("should draw legend to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("svg", 0755)

		settings := types.Settings{}

		expectedSVG := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"200\" height=\"100\" viewBox=\"0 0 200 100\">\n" +
			"\t<rect width=\"100%\" height=\"100%\" fill=\"white\" />\n" +
			"\t<text y=\"20\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">specificity</text>\n" +
			"\t<circle cx=\"20\" cy=\"50\" fill=\"#6e97ff\" r=\"6\" />\n" +
			"\t<text font-size=\"12\" x=\"35\" y=\"54\">Infinite specificity</text>\n" +
			"\t<circle cx=\"20\" cy=\"80\" fill=\"#dfcd06\" r=\"6\" />\n" +
			"\t<text font-size=\"12\" x=\"35\" y=\"84\">Finite specificity</text>\n" +
			"</svg>\n"

		createLegend(settings)
		actualSVG, _ := afero.ReadFile(fs.Instance, "svg/specificity-legend.svg")
		Expect(string(actualSVG)).To(Equal(expectedSVG))
	})
})
