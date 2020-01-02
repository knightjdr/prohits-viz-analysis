// Package heatmap draws a svg heatmap.
package heatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/spf13/afero"
)

// Heatmap object.
type Heatmap struct {
	AbundanceCap float64
	Annotations  types.Annotations
	CellSize     int
	Columns      []string
	FillColor    string
	FontSize     int
	Invert       bool
	LeftMargin   int
	Markers      types.Markers
	Matrix       [][]float64
	MinAbundance float64
	NumColors    int
	PlotHeight   int
	PlotWidth    int
	Rows         []string
	SvgHeight    int
	SvgWidth     int
	TopMargin    int
	XLabel       string
	YLabel       string
}

// Initialize a heatmap.
func Initialize() *Heatmap {
	return &Heatmap{
		FillColor: "blue",
		FontSize:  12,
		NumColors: 101,
	}
}

// Draw a heatmap in svg format.
func (h *Heatmap) Draw(filename string) {
	file, err := fs.Instance.Create(filename)
	log.CheckError(err, true)
	defer file.Close()

	writeString := WriteElement(file)

	WriteHeader(h, writeString)
	WriteLabels(h, writeString)
	writeCells(h, writeString)
	WriteMarkup(h, writeString)

	writeString("</svg>\n")
}

// WriteElement writes an svg element/tag to a file.
func WriteElement(file afero.File) func(str string) {
	return func(str string) {
		file.WriteString(str)
	}
}
