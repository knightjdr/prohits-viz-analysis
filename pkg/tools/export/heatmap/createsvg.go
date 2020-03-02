package heatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createSVG(data *Heatmap, matrices *types.Matrices) {
	files.CreateFolders([]string{"svg"})

	dims := dimensions.Calculate(matrices.Abundance, matrices.Conditions, matrices.Readouts, false)

	image := svg.InitializeHeatmap()
	image.AbundanceCap = data.Settings.AbundanceCap
	image.Annotations = data.Annotations
	image.CellSize = dims.CellSize
	image.Columns = matrices.Conditions
	image.FillColor = data.Settings.FillColor
	image.FontSize = dims.FontSize
	image.Invert = data.Settings.InvertColor
	image.LeftMargin = dims.LeftMargin
	image.Markers = data.Markers
	image.Matrix = matrices.Abundance
	image.MinAbundance = data.Settings.MinAbundance
	image.PlotHeight = dims.PlotHeight
	image.PlotWidth = dims.PlotWidth
	image.Rows = matrices.Readouts
	image.SvgHeight = dims.SvgHeight
	image.SvgWidth = dims.SvgWidth
	image.TopMargin = dims.TopMargin
	image.XLabel = data.Settings.XLabel
	image.YLabel = data.Settings.YLabel

	image.Draw("svg/heatmap.svg")
}
