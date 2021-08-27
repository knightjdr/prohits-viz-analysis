package minimap

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/pkg/png"
)

func createHeatmap(data *Data) {
	matrix := downsampleIfNeeded(data)
	dims := dimensions.Calculate(matrix, []string{}, []string{}, true)

	image := png.InitializeHeatmap()
	image.ColorSpace = data.Settings.FillColor
	image.CellSize = dims.CellSize
	image.FillMax = data.Settings.FillMax
	image.FillMin = data.Settings.FillMin
	image.Height = dims.PlotHeight
	image.Invert = data.Settings.InvertColor
	image.Width = dims.PlotWidth

	image.Draw(matrix, data.Filename)
}
