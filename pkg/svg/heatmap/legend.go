package heatmap

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/color"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/math"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
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

	CreateLegendHeader(&svg)
	CreateLegendTitle(&svg, data.Title)

	gradientData := &Heatmap{
		FillColor: data.Settings.FillColor,
		Invert:    data.Settings.InvertColor,
		NumColors: data.NumColors,
	}
	fillGradient := createGradient(gradientData)
	CreateFillGradient(&svg, data, fillGradient)

	svg.WriteString("</svg>\n")
	afero.WriteFile(fs.Instance, data.Filename, []byte(svg.String()), 0644)
}

// CreateLegendHeader writes the opening svg tag.
func CreateLegendHeader(svg *strings.Builder) {
	svg.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"200\" height=\"240\" viewBox=\"0 0 200 240\">\n")
	svg.WriteString("\t<rect width=\"100%\" height=\"100%\" fill=\"white\" />\n")
}

// CreateLegendTitle writes the legend title.
func CreateLegendTitle(svg *strings.Builder, title string) {
	svg.WriteString(fmt.Sprintf("\t<text y=\"20\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">%s</text>\n", title))
}

// CreateFillGradient writes a gradient with the fill color.
func CreateFillGradient(svg *strings.Builder, data Legend, gradient []color.Space) {
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
