package heatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/png"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func createPNG(data *heatmap, matrices *types.Matrices, settings Settings) {
	files.CreateFolders([]string{"png"})

	downsampleData(data, matrices, settings.DownsampleThreshold)

	dims := dimensions.Calculate(matrices.Abundance, []string{}, []string{}, false)

	image := png.InitializeHeatmap()
	image.AbundanceCap = data.Settings.AbundanceCap
	image.Annotations = data.Annotations
	image.ColorSpace = data.Settings.FillColor
	image.CellSize = dims.CellSize
	image.FontPath = settings.FontPath
	image.Height = dims.PlotHeight
	image.Invert = data.Settings.InvertColor
	image.Markers = data.Markers
	image.MinAbundance = data.Settings.MinAbundance
	image.Width = dims.PlotWidth

	image.Draw(matrices.Abundance, "png/heatmap.png")
}
