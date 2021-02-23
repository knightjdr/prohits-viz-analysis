package circheatmap

import (
	"fmt"
	"math"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/color"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/spf13/afero"
)

// LegendElement contains settings for each metric in the circheatmap.
type LegendElement struct {
	Color     string
	Max       float64
	Min       float64
	Attribute string
}

// Legend settings.
type Legend struct {
	Elements []LegendElement
	Filename string
	Known    string
	Title    string
}

// CreateLegend creates a legend for a circheatmap.
func CreateLegend(data Legend) {
	var svg strings.Builder

	svgHeight := calculateLegendHeight(len(data.Elements), data.Known)
	createLegendHeader(&svg, svgHeight)

	createLegendTitle(&svg, data.Title)
	svg.WriteString("\t<g transform=\"translate(0 30)\">\n")
	createLegendGradients(&svg, data.Elements)
	createLegendKnownBar(&svg, data.Known, svgHeight)
	svg.WriteString("\t</g>\n")

	svg.WriteString("</svg>\n")
	afero.WriteFile(fs.Instance, data.Filename, []byte(svg.String()), 0644)
}

func calculateLegendHeight(noElements int, known string) int {
	gradientHeight := (noElements * 50) + 40
	if known != "" {
		gradientHeight += 40
	}
	return gradientHeight
}

func createLegendHeader(svg *strings.Builder, height int) {
	svg.WriteString(fmt.Sprintf("<svg  xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"210\" height=\"%[1]d\" viewBox=\"0 0 210 %[1]d\">\n", height))
	svg.WriteString("\t<rect width=\"100%\" height=\"100%\" fill=\"white\" />\n")
}

func createLegendTitle(svg *strings.Builder, title string) {
	svg.WriteString(fmt.Sprintf("\t<text y=\"20\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">%s</text>\n", title))
}

func createLegendGradients(svg *strings.Builder, elements []LegendElement) {
	numColors := 101
	halfColorIndex := int(math.Floor(float64(numColors) / 2))

	for i, element := range elements {
		colorSettings := color.Gradient{
			ColorSpace: element.Color,
			NumColors:  numColors,
		}
		gradientFill := colorSettings.CreateColorGradient()

		svg.WriteString(fmt.Sprintf("\t\t<g transform=\"translate(0 %d)\">\n", i*50))
		svg.WriteString(fmt.Sprintf(
			"\t\t\t<defs>\n"+
				"\t\t\t\t<linearGradient id=\"%s-legendGradient\">\n"+
				"\t\t\t\t\t<stop offset=\"0%%\" stop-color=\"%s\" />\n"+
				"\t\t\t\t\t<stop offset=\"50%%\" stop-color=\"%s\" />\n"+
				"\t\t\t\t\t<stop offset=\"100%%\" stop-color=\"%s\" />\n"+
				"\t\t\t\t</linearGradient>\n"+
				"\t\t\t</defs>\n",
			element.Attribute,
			gradientFill[0].Hex,
			gradientFill[halfColorIndex].Hex,
			gradientFill[numColors-1].Hex,
		))
		svg.WriteString(fmt.Sprintf(
			"\t\t\t<g>\n"+
				"\t\t\t\t<text x=\"100\" y=\"20\" text-anchor=\"middle\">%[1]s</text>\n"+
				"\t\t\t\t<rect x=\"25\" y=\"30\" height=\"20\" width=\"150\" fill=\"url(#%[1]s-legendGradient)\" />\n"+
				"\t\t\t\t<text x=\"20\" y=\"45\" text-anchor=\"end\">%s</text>\n"+
				"\t\t\t\t<text x=\"180\" y=\"45\" text-anchor=\"start\">%s</text>\n"+
				"\t\t\t</g>\n",
			element.Attribute,
			float.RemoveTrailingZeros(element.Min),
			float.RemoveTrailingZeros(element.Max),
		))
		svg.WriteString("\t\t</g>\n")
	}
}

func createLegendKnownBar(svg *strings.Builder, known string, height int) {
	/* <g transform={`translate(0 ${gradientHeight + 70})`}>
		<text
			textAnchor="middle"
			x="100"
			y="0"
		>
			Known
		</text>
		<line
			stroke="black"
			strokeWidth="3"
			x1="50"
			x2="150"
			y1="10"
			y2="10"
		/>
	</g> */
	if known != "" {
		svg.WriteString(fmt.Sprintf(
			"\t\t<g transform=\"translate(0 %d)\">\n"+
				"\t\t\t<text text-anchor=\"middle\" x=\"100\" y=\"0\">Known %s</text>\n"+
				"\t\t\t<line stroke=\"black\" stroke-width=\"3\" x1=\"50\" x2=\"150\" y1=\"10\" y2=\"10\"/>\n"+
				"\t\t</g>\n",
			height-60,
			known,
		))
	}
}
