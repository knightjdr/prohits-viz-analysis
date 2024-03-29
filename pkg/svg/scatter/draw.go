// Package scatter draws a svg scatter plot.
package scatter

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

// Scatter plot properties
type Scatter struct {
	Axes     Axes
	LogBase  string
	Plot     []types.ScatterPoint
	PlotSize float64 // Height/width in pixels
	Ticks    Ticks
	XLabel   string
	YLabel   string
}

type boundaries struct {
	x boundary
	y boundary
}

type boundary struct {
	max float64
	min float64
}

// Axes lines
type Axes struct {
	X Line
	Y Line
}

// Line for scatterplot
type Line struct {
	X1 float64
	X2 float64
	Y1 float64
	Y2 float64
}

// Ticks for axes
type Ticks struct {
	X      []float64
	XLabel []string
	Y      []float64
	YLabel []string
}

// Initialize a scatter plot.
func Initialize() *Scatter {
	return &Scatter{
		PlotSize: 750,
		XLabel:   "x",
		YLabel:   "y",
	}
}

// Draw a scatter in svg format.
func (s *Scatter) Draw(filename string) {
	file, err := fs.Instance.Create(filename)
	log.CheckError(err, true)
	defer file.Close()

	axisLength := s.PlotSize - 150
	formatData(s, axisLength)

	writeString := WriteElement(file)

	writeHeader(s, writeString)
	writeBackground(writeString)
	writePlot(s, axisLength, writeString)

	writeString("</svg>\n")
}

// WriteElement writes an svg element/tag to a file.
func WriteElement(file afero.File) func(str string) {
	return func(str string) {
		file.WriteString(str)
	}
}
