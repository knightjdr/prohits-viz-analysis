package heatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/pkg/png"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createPNG(data *Heatmap, matrices *types.Matrices, settings Settings) {
	files.CreateFolders([]string{"png"})

	DownsampleData(data, matrices, settings.DownsampleThreshold)

	dims := dimensions.Calculate(matrices.Abundance, []string{}, []string{}, false)

	image := png.InitializeHeatmap()
	image.Annotations = data.Annotations
	image.ColorSpace = data.Settings.FillColor
	image.CellSize = dims.CellSize
	image.FillMax = data.Settings.FillMax
	image.FillMin = data.Settings.FillMin
	image.FontPath = settings.FontPath
	image.Height = dims.PlotHeight
	image.Invert = data.Settings.InvertColor
	image.Markers = data.Markers
	image.Width = dims.PlotWidth

	image.Draw(matrices.Abundance, "png/heatmap.png")
}
