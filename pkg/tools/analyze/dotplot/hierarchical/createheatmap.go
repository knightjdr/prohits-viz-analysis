package hierarchical

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// CreateHeatmap image.
func CreateHeatmap(data *SortedData, clusteredData HclustData, settings types.Settings) {
	if settings.WriteHeatmap {
		createHeatmapSVG(data, settings)
		createHeatmapLegend(settings)
	}
}

func createHeatmapSVG(data *SortedData, settings types.Settings) {
	dims := dimensions.Calculate(data.Matrices.Abundance, data.Matrices.Conditions, data.Matrices.Readouts, false)

	heatmap := svg.InitializeHeatmap()
	heatmap.CellSize = dims.CellSize
	heatmap.Columns = data.Matrices.Conditions
	heatmap.FillColor = settings.FillColor
	heatmap.FillMax = settings.FillMax
	heatmap.FillMin = settings.FillMin
	heatmap.FontSize = dims.FontSize
	heatmap.Invert = settings.InvertColor
	heatmap.LeftMargin = dims.LeftMargin
	heatmap.Matrix = data.Matrices.Abundance
	heatmap.PlotHeight = dims.PlotHeight
	heatmap.PlotWidth = dims.PlotWidth
	heatmap.Rows = data.Matrices.Readouts
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
