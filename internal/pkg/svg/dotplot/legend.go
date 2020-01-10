package dotplot

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/color"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/math"
	"github.com/spf13/afero"
)

// Legend settings.
type Legend struct {
	Filename  string
	NumColors int
	Settings  types.Settings
	Title     string
}

// CreateLegend creates a legend for a dotplot.
func CreateLegend(data Legend) {
	var svg strings.Builder

	createLegendHeader(&svg)
	createLegendTitle(&svg, data.Title)

	gradientData := &Dotplot{
		EdgeColor: data.Settings.EdgeColor,
		FillColor: data.Settings.FillColor,
		Invert:    data.Settings.InvertColor,
		NumColors: data.NumColors,
	}
	fillGradient, edgeGradient := createGradients(gradientData)
	createFillGradient(&svg, data, fillGradient)

	createAbundanceElement(&svg)
	createScoreElement(&svg, data, edgeGradient)

	svg.WriteString("</svg>\n")
	afero.WriteFile(fs.Instance, data.Filename, []byte(svg.String()), 0644)
}

func createLegendHeader(svg *strings.Builder) {
	svg.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"200\" height=\"240\" viewBox=\"0 0 200 240\">\n")
}

func createLegendTitle(svg *strings.Builder, title string) {
	svg.WriteString(fmt.Sprintf("\t<text y=\"20\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">%s</text>\n", title))
}

func createFillGradient(svg *strings.Builder, data Legend, gradient []color.Space) {
	gradientCellWidth := math.Round(float64(150)/float64(data.NumColors), 0.01)
	gradientCellWidthString := float.RemoveTrailingZeros(gradientCellWidth)

	svg.WriteString("\t<g>\n")
	for i, color := range gradient {
		x := (float64(i) * gradientCellWidth) + 25
		svg.WriteString(
			fmt.Sprintf(
				"\t\t<rect fill=\"%s\" y=\"30\" x=\"%s\" width=\"%s\" height=\"20\" />\n",
				color.Hex,
				float.RemoveTrailingZeros(x),
				gradientCellWidthString,
			),
		)
	}
	svg.WriteString("\t</g>\n")

	// Draw border around gradient.
	svg.WriteString("\t<rect fill=\"none\" y=\"30\" x=\"25\" width=\"150\" height=\"20\" stroke=\"#000000\" stroke-width=\"1\"/>\n")

	// Add min and max labels.
	svg.WriteString(
		fmt.Sprintf(
			"\t<text y=\"65\" x=\"175\" font-size=\"12\" text-anchor=\"middle\">%s</text>\n",
			float.RemoveTrailingZeros(data.Settings.AbundanceCap),
		),
	)
	svg.WriteString(
		fmt.Sprintf(
			"\t<text y=\"65\" x=\"25\" font-size=\"12\" text-anchor=\"middle\">%s</text>\n",
			float.RemoveTrailingZeros(data.Settings.MinAbundance),
		),
	)
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
