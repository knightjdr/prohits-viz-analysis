// Package dotplot draws a svg dotplot.
package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Dotplot object.
type Dotplot struct {
	AbundanceCap    float64
	Annotations     types.Annotations
	BoundingBox     bool
	CellSize        int
	Columns         []string
	EdgeColor       string
	FillColor       string
	FontSize        int
	Invert          bool
	LeftMargin      int
	Markers         types.Markers
	Matrices        *types.Matrices
	MinAbundance    float64
	NumColors       int
	PlotHeight      int
	PlotWidth       int
	PrimaryFilter   float64
	Ratio           float64
	Rows            []string
	ScoreType       string
	SecondaryFilter float64
	SvgHeight       int
	SvgWidth        int
	TopMargin       int
	XLabel          string
	YLabel          string
}

// Initialize a heatmap.
func Initialize() *Dotplot {
	return &Dotplot{
		BoundingBox: true,
		EdgeColor:   "blue",
		FillColor:   "blue",
		FontSize:    12,
		NumColors:   101,
		ScoreType:   "lte",
	}
}

// Draw a dotplot in svg format.
func (d *Dotplot) Draw(filename string) {
	file, err := fs.Instance.Create(filename)
	log.CheckError(err, true)
	defer file.Close()

	writeString := heatmap.WriteElement(file)

	writeHeader(d, writeString)
	writeLabels(d, writeString)
	writeDots(d, writeString)
	writeBoundingBox(d, writeString)
	writeMarkup(d, writeString)

	writeString("</svg>\n")
}
