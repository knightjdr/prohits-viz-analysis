package heatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/color"
)

func createGradient(h *Heatmap) []color.Space {
	gradient := color.InitializeGradient()
	gradient.ColorSpace = h.FillColor
	gradient.Invert = h.Invert
	gradient.NumColors = h.NumColors
	return gradient.CreateColorGradient()
}
