package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/color"
)

func createGradients(d *Dotplot) ([]color.Space, []color.Space) {
	fillGradient := color.InitializeGradient()
	fillGradient.ColorSpace = d.FillColor
	fillGradient.Invert = d.Invert
	fillGradient.NumColors = d.NumColors

	edgeGradient := color.InitializeGradient()
	edgeGradient.ColorSpace = d.EdgeColor
	edgeGradient.NumColors = d.NumColors

	return fillGradient.CreateColorGradient(), edgeGradient.CreateColorGradient()
}
