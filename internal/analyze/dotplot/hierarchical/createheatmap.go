package hierarchical

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func createHeatmap(data *sortedData, clusteredData hclustData, settings types.Settings) {
	if settings.WriteHeatmap {
		createHeatmapSVG(data, settings)
		createHeatmapLegend(settings)
	}
}

func createHeatmapSVG(data *sortedData, settings types.Settings) {
	dims := dimensions.Calculate(data.matrices.Abundance, data.matrices.Conditions, data.matrices.Readouts, false)

	heatmap := svg.InitializeHeatmap()
	heatmap.AbundanceCap = settings.AbundanceCap
	heatmap.CellSize = dims.CellSize
	heatmap.Columns = data.matrices.Conditions
	heatmap.FillColor = settings.FillColor
	heatmap.FontSize = dims.FontSize
	heatmap.Invert = settings.InvertColor
	heatmap.LeftMargin = dims.LeftMargin
	heatmap.Matrix = data.matrices.Abundance
	heatmap.MinAbundance = settings.MinAbundance
	heatmap.PlotHeight = dims.PlotHeight
	heatmap.PlotWidth = dims.PlotWidth
	heatmap.Rows = data.matrices.Readouts
	heatmap.SvgHeight = dims.SvgHeight
	heatmap.SvgWidth = dims.SvgWidth
	heatmap.TopMargin = dims.TopMargin
	heatmap.XLabel = settings.Condition
	heatmap.YLabel = settings.Readout

	heatmap.Draw("svg/heatmap.svg")
}

func createHeatmapLegend(settings types.Settings) {
	legendData := heatmap.Legend{
		Filename:  "svg/heatmap-legend.svg",
		NumColors: 101,
		Settings:  settings,
		Title:     fmt.Sprintf("Heatmap - %s", settings.Abundance),
	}
	heatmap.CreateLegend(legendData)
}
