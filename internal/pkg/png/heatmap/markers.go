package heatmap

import (
	"image"
	"image/color"

	customColor "github.com/knightjdr/prohits-viz-analysis/internal/pkg/color"
)

type box struct {
	color  color.RGBA
	height int
	width  int
	x      int
	y      int
}

type line struct {
	x0 int
	x1 int
	y0 int
	y1 int
}

type point struct {
	x int
	y int
}

func drawMarkers(img *image.RGBA, h *Heatmap) {
	if len(h.Markers.List) > 0 {
		color, _ := customColor.ConvertHexToRGB(h.Markers.Color)
		for _, marker := range h.Markers.List {
			b := box{
				color:  color,
				height: marker.Height * h.CellSize,
				width:  marker.Width * h.CellSize,
				x:      int(marker.X * float64(h.Width)),
				y:      int(marker.Y * float64(h.Height)),
			}
			drawBox(img, b)
		}
	}
}

func drawBox(img *image.RGBA, b box) {
	line1 := line{x0: b.x, x1: b.x + b.width, y0: b.y}
	drawLine(img, b.color, getLineCoordinates("x", line1))

	line2 := line{x0: b.x, y0: b.y, y1: b.y + b.height}
	drawLine(img, b.color, getLineCoordinates("y", line2))

	line3 := line{x0: b.x, x1: b.x + b.width, y0: b.y + b.height}
	drawLine(img, b.color, getLineCoordinates("x", line3))

	line4 := line{x0: b.x + b.width, y0: b.y, y1: b.y + b.height}
	drawLine(img, b.color, getLineCoordinates("y", line4))
}

func getLineCoordinates(direction string, l line) []point {
	coordinates := make([]point, 0)
	if direction == "x" {
		for x := l.x0; x <= l.x1; x++ {
			coordinates = append(coordinates, point{x: x, y: l.y0})
		}
	} else {
		for y := l.y0; y <= l.y1; y++ {
			coordinates = append(coordinates, point{x: l.x0, y: y})
		}
	}

	return coordinates
}

func drawLine(img *image.RGBA, color color.RGBA, coordinates []point) {
	for _, position := range coordinates {
		img.Set(position.x, position.y, color)
	}
}
