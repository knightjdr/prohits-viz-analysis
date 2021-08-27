// Package heatmap draws a png heatmap.
package heatmap

import (
	"image"
	"image/png"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Heatmap object.
type Heatmap struct {
	Annotations types.Annotations
	CellSize    int
	ColorSpace  string
	FillMax     float64
	FillMin     float64
	FontPath    string
	Height      int
	Invert      bool
	Markers     types.Markers
	NumColors   int
	Width       int
}

// Initialize a heatmap.
func Initialize() *Heatmap {
	return &Heatmap{
		ColorSpace: "blue",
		NumColors:  101,
	}
}

// Draw a heatmap in png format.
func (h *Heatmap) Draw(matrix [][]float64, filename string) {
	img := initializeImage(h)

	drawGrid(img, h, matrix)
	drawMarkers(img, h)
	drawAnnotations(img, h)

	file, _ := fs.Instance.Create(filename)
	png.Encode(file, img)
}

func initializeImage(h *Heatmap) *image.RGBA {
	topLeft := image.Point{0, 0}
	lowRight := image.Point{h.Width, h.Height}
	return image.NewRGBA(image.Rectangle{topLeft, lowRight})
}
