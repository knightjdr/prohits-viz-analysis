// Package heatmap draws a png heatmap.
package heatmap

import (
	"image"
	"image/color"
	"image/png"
	"math"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
)

// Heatmap object.
type Heatmap struct {
	AbundanceCap float64
	CellSize     int
	ColorSpace   string
	Height       int
	Invert       bool
	MinAbundance float64
	NumColors    int
	Width        int
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
	fillPixels := getCellFiller(img, h)

	for i, row := range matrix {
		for j, value := range row {
			fillPixels(i, j, value)
		}
	}

	file, _ := fs.Instance.Create(filename)
	png.Encode(file, img)
}

func initializeImage(h *Heatmap) *image.RGBA {
	topLeft := image.Point{0, 0}
	lowRight := image.Point{h.Width, h.Height}
	return image.NewRGBA(image.Rectangle{topLeft, lowRight})
}

func getCellFiller(img *image.RGBA, h *Heatmap) func(int, int, float64) {
	gradient := createGradient(h)
	findGradientIndex := getGradientIndex(h)

	return func(rowIndex, columnIndex int, value float64) {
		startX := columnIndex * h.CellSize
		startY := rowIndex * h.CellSize
		gradientColor := gradient[findGradientIndex(value)]
		cellColor := color.RGBA{uint8(gradientColor.RGB[0]), uint8(gradientColor.RGB[1]), uint8(gradientColor.RGB[2]), 255}
		for x := 0; x < h.CellSize; x++ {
			for y := 0; y < h.CellSize; y++ {
				img.Set(x+startX, y+startY, cellColor)
			}
		}
	}
}

func getGradientIndex(h *Heatmap) func(float64) int {
	mapToRange := float.GetRange(h.MinAbundance, h.AbundanceCap, 0, float64(h.NumColors-1))
	return func(value float64) int {
		outputValue := mapToRange(value)
		return int(math.Round(outputValue))
	}
}
