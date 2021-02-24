package scatter

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Draw scatter legend", func() {
	It("should draw legend to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := Legend{
			Filename: "legend.svg",
			Points: []map[string]string{
				{"color": "#0066cc", "text": "point1"},
				{"color": "#99ccff", "text": "point2"},
			},
			Title: "scatter title",
		}

		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"200\" height=\"100\" viewBox=\"0 0 200 100\">\n" +
			"\t<rect width=\"100%\" height=\"100%\" fill=\"white\" />\n" +
			"\t<text y=\"20\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">scatter title</text>\n" +
			"\t<circle cx=\"20\" cy=\"50\" fill=\"#0066cc\" r=\"6\" />\n" +
			"\t<text font-size=\"12\" x=\"35\" y=\"54\">point1</text>\n" +
			"\t<circle cx=\"20\" cy=\"80\" fill=\"#99ccff\" r=\"6\" />\n" +
			"\t<text font-size=\"12\" x=\"35\" y=\"84\">point2</text>\n" +
			"</svg>\n"

		CreateLegend(data)
		actual, _ := afero.ReadFile(fs.Instance, "legend.svg")
		Expect(string(actual)).To(Equal(expected))
	})
})
