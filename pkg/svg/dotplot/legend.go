package dotplot

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/color"
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/spf13/afero"
)

// Legend settings.
type Legend = heatmap.Legend

// CreateLegend creates a legend for a dotplot.
func CreateLegend(data Legend) {
	var svg strings.Builder

	heatmap.CreateLegendHeader(&svg)
	heatmap.CreateLegendTitle(&svg, data.Title)

	gradientData := &Dotplot{
		EdgeColor: data.Settings.EdgeColor,
		FillColor: data.Settings.FillColor,
		Invert:    data.Settings.InvertColor,
		NumColors: data.NumColors,
	}
	fillGradient, edgeGradient := createGradients(gradientData)
	heatmap.CreateFillGradient(&svg, data, fillGradient)

	createAbundanceElement(&svg)
	createScoreElement(&svg, data, edgeGradient)

	svg.WriteString("</svg>\n")
	afero.WriteFile(fs.Instance, data.Filename, []byte(svg.String()), 0644)
}

func createAbundanceElement(svg *strings.Builder) {
	svg.WriteString("\t<g>\n" +
		"\t\t<circle fill=\"#000000\" cy=\"100\" cx=\"60\" r=\"6\" />\n" +
		"\t\t<circle fill=\"#000000\" cy=\"100\" cx=\"135\" r=\"12\" />\n" +
		"\t\t<line fill=\"none\" stroke=\"#000000\" stroke-width=\"1\" x1=\"70\" y1=\"100\" x2=\"119\" y2=\"100\"/>\n" +
		"\t\t<polygon fill=\"#000000\" points=\"110,96 112,100 110,104 119,100\"/>\n" +
		"\t\t<text y=\"130\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">Relative abundance</text>\n" +
		"\t</g>\n",
	)
}

func createScoreElement(svg *strings.Builder, data Legend, gradient []color.Space) {
	scoreSymbols := createLegendScoreSymbol(data.Settings.ScoreType)

	svg.WriteString(
		fmt.Sprintf(
			"\t<g>\n"+
				"\t\t<text y=\"220\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">%s</text>\n"+
				"\t\t<circle fill=\"none\" cy=\"165\" cx=\"50\" r=\"12\" stroke=\"%s\" stroke-width=\"2\" />\n"+
				"\t\t<text y=\"195\" x=\"50\" font-size=\"12\" text-anchor=\"middle\">%s %s</text>\n"+
				"\t\t<circle fill=\"none\" cy=\"165\" cx=\"100\" r=\"12\" stroke=\"%s\" stroke-width=\"2\" />\n"+
				"\t\t<text y=\"195\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">%s %s</text>\n"+
				"\t\t<circle fill=\"none\" cy=\"165\" cx=\"150\" r=\"12\" stroke=\"%s\" stroke-width=\"2\" />\n"+
				"\t\t<text y=\"195\" x=\"150\" font-size=\"12\" text-anchor=\"middle\">%s %s</text>\n"+
				"\t</g>\n",
			data.Settings.Score,
			gradient[data.NumColors-1].Hex,
			scoreSymbols[0],
			float.RemoveTrailingZeros(data.Settings.PrimaryFilter),
			gradient[(data.NumColors-1)/2].Hex,
			scoreSymbols[1],
			float.RemoveTrailingZeros(data.Settings.SecondaryFilter),
			gradient[(data.NumColors-1)/4].Hex,
			scoreSymbols[2],
			float.RemoveTrailingZeros(data.Settings.SecondaryFilter),
		),
	)
}

func createLegendScoreSymbol(scoreType string) []string {
	if scoreType == "gte" {
		return []string{"≥", "≥", "<"}
	}
	return []string{"≤", "≤", ">"}
}
