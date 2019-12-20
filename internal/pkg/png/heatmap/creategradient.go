package heatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/color"
)

func createGradient(h *Heatmap) []color.Space {
	gradient := color.InitializeGradient()
	gradient.ColorSpace = h.ColorSpace
	gradient.ColorType = "RGB"
	gradient.Invert = h.Invert
	gradient.NumColors = h.NumColors
	return gradient.CreateColorGradient()
}
