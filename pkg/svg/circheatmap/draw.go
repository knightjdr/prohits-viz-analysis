// Package circheatmap draws a svg circular heat map.
package circheatmap

import (
	"math"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

// Circular heat map properties
type CircHeatmapSVG struct {
	Dimensions CircHeatmapDimensions
	Legend     types.CircHeatmapLegend
	Plot       types.CircHeatmap
	ShowKnown  bool
}

// CircHeatmapDimensions contains dimensions for plotting
type CircHeatmapDimensions struct {
	Center    float64
	PlotSize  float64 // height/width in pixels
	Radius    float64 // radius of outermost drawing area
	Thickness float64 // thickness of each circle
}

// Initialize a circular heat map.
func Initialize() *CircHeatmapSVG {
	return &CircHeatmapSVG{
		Dimensions: CircHeatmapDimensions{
			Center:   375,
			PlotSize: 750,
			Radius:   360,
		},
	}
}

// SetThickness of each circle
func (c *CircHeatmapSVG) SetThickness() {
	numCircles := float64(len(c.Legend))
	radius := c.Dimensions.PlotSize / 2
	c.Dimensions.Thickness = math.Floor(radius / (1 + (1.25 * numCircles)))
}

// Draw a circular heatmap in svg format.
func (c *CircHeatmapSVG) Draw(filename string) {
	file, err := fs.Instance.Create(filename)
	log.CheckError(err, true)
	defer file.Close()

	writeString := WriteElement(file)

	c.SetThickness()

	writeHeader(c, writeString)
	writeBackground(c, writeString)
	writePlot(c, writeString)
	writeString("</svg>\n")
}

// WriteElement writes an svg element/tag to a file.
func WriteElement(file afero.File) func(str string) {
	return func(str string) {
		file.WriteString(str)
	}
}
