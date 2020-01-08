package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/export/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func createSVG(data *heatmap.Heatmap, matrices *types.Matrices) {
	files.CreateFolders([]string{"svg"})

	dims := dimensions.Calculate(matrices.Abundance, matrices.Conditions, matrices.Readouts, false)

	image := svg.InitializeDotplot()
	image.AbundanceCap = data.Settings.AbundanceCap
	image.Annotations = data.Annotations
	image.CellSize = dims.CellSize
	image.Columns = matrices.Conditions
	image.EdgeColor = data.Settings.EdgeColor
	image.FillColor = data.Settings.FillColor
	image.FontSize = dims.FontSize
	image.Invert = data.Settings.InvertColor
	image.LeftMargin = dims.LeftMargin
	image.Markers = data.Markers
	image.Matrices = matrices
	image.MinAbundance = data.Settings.MinAbundance
	image.PlotHeight = dims.PlotHeight
	image.PlotWidth = dims.PlotWidth
	image.PrimaryFilter = data.Settings.PrimaryFilter
	image.Ratio = dims.Ratio
	image.Rows = matrices.Readouts
	image.ScoreType = data.Settings.ScoreType
	image.SecondaryFilter = data.Settings.SecondaryFilter
	image.SvgHeight = dims.SvgHeight
	image.SvgWidth = dims.SvgWidth
	image.TopMargin = dims.TopMargin
	image.XLabel = data.Settings.XLabel
	image.YLabel = data.Settings.YLabel

	image.Draw("svg/dotplot.svg")
}