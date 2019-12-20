package minimap

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/png/heatmap"
)

func createImage(imageType string, data *Data) {
	dimensions := dimensions.Calculate(data.Matrices.Abundance, []string{}, []string{}, true)
	if imageType == "dotplot" {

	} else {
		image := heatmap.Initialize()
		image.AbundanceCap = data.Settings.AbundanceCap
		image.ColorSpace = data.Settings.FillColor
		image.CellSize = dimensions.CellSize
		image.Height = dimensions.PlotHeight
		image.Invert = data.Settings.InvertColor
		image.MinAbundance = data.Settings.MinAbundance
		image.Width = dimensions.PlotWidth
		image.Draw(data.Matrices.Abundance, data.Filename)
	}

}
