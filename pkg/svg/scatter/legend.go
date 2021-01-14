package scatter

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/spf13/afero"
)

// Legend settings.
type Legend struct {
	Filename string
	Points   []map[string]string
	Title    string
}

// CreateLegend creates a legend for a dotplot.
func CreateLegend(data Legend) {
	var svg strings.Builder

	svgHeight := calculateLegendHeight(data.Points)
	createLegendHeader(&svg, svgHeight)

	createLegendTitle(&svg, data.Title)
	createLegendPoints(&svg, data.Points)

	svg.WriteString("</svg>\n")
	afero.WriteFile(fs.Instance, data.Filename, []byte(svg.String()), 0644)
}

func calculateLegendHeight(points []map[string]string) int {
	return 40 + (len(points) * 30)
}

func createLegendHeader(svg *strings.Builder, height int) {
	svg.WriteString(fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"200\" height=\"%[1]d\" viewBox=\"0 0 200 %[1]d\">\n", height))
}

func createLegendTitle(svg *strings.Builder, title string) {
	svg.WriteString(fmt.Sprintf("\t<text y=\"20\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">%s</text>\n", title))
}

func createLegendPoints(svg *strings.Builder, points []map[string]string) {
	for i, point := range points {
		yCircle := (i * 30) + 50
		yText := (i * 30) + 50 + 4
		svg.WriteString(fmt.Sprintf("\t<circle cx=\"20\" cy=\"%d\" fill=\"%s\" r=\"6\" />\n", yCircle, point["color"]))
		svg.WriteString(fmt.Sprintf("\t<text font-size=\"12\" x=\"35\" y=\"%d\">%s</text>\n", yText, point["text"]))
	}
}
