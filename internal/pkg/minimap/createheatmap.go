package minimap

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/png"
)

func createHeatmap(data *Data, dims *dimensions.Heatmap) {
	image := png.InitializeHeatmap()
	image.AbundanceCap = data.Settings.AbundanceCap
	image.ColorSpace = data.Settings.FillColor
	image.CellSize = dims.CellSize
	image.Height = dims.PlotHeight
	image.Invert = data.Settings.InvertColor
	image.MinAbundance = data.Settings.MinAbundance
	image.Width = dims.PlotWidth

	image.Draw(data.Matrices.Abundance, data.Filename)
}
