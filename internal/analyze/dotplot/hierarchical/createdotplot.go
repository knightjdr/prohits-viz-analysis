package hierarchical

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/minimap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/dotplot"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/treeview"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func createDotplot(data *sortedData, clusteredData hclustData, settings types.Settings) {
	createDotplotSVG(data, settings)
	createLegend(data, settings)
	createMinimap(data, settings)
	createInteractive(data, settings)
	createTreeview(data, clusteredData)
}

func createDotplotSVG(data *sortedData, settings types.Settings) {
	dims := dimensions.Calculate(data.matrices.Abundance, data.matrices.Conditions, data.matrices.Readouts, false)

	dotplot := svg.InitializeDotplot()
	dotplot.AbundanceCap = settings.AbundanceCap
	dotplot.CellSize = dims.CellSize
	dotplot.Columns = data.matrices.Conditions
	dotplot.EdgeColor = settings.EdgeColor
	dotplot.FillColor = settings.FillColor
	dotplot.FontSize = dims.FontSize
	dotplot.Invert = settings.InvertColor
	dotplot.LeftMargin = dims.LeftMargin
	dotplot.Matrices = data.matrices
	dotplot.MinAbundance = settings.MinAbundance
	dotplot.PlotHeight = dims.PlotHeight
	dotplot.PlotWidth = dims.PlotWidth
	dotplot.PrimaryFilter = settings.PrimaryFilter
	dotplot.Ratio = dims.Ratio
	dotplot.Rows = data.matrices.Readouts
	dotplot.ScoreType = settings.ScoreType
	dotplot.SecondaryFilter = settings.SecondaryFilter
	dotplot.SvgHeight = dims.SvgHeight
	dotplot.SvgWidth = dims.SvgWidth
	dotplot.TopMargin = dims.TopMargin
	dotplot.XLabel = settings.Condition
	dotplot.YLabel = settings.Readout

	dotplot.Draw("svg/dotplot.svg")
}

func createMinimap(data *sortedData, settings types.Settings) {
	minimapData := &minimap.Data{
		DownsampleThreshold: 1000,
		Filename:            "minimap/minimap.png",
		ImageType:           "dotplot",
		Matrices:            data.matrices,
		Settings:            settings,
	}
	minimap.Create(minimapData)
}

func createInteractive(data *sortedData, settings types.Settings) {
	interactiveData := &interactive.HeatmapData{
		Filename:  "interactive/dotplot.json",
		ImageType: "dotplot",
		Matrices:  data.matrices,
		Minimap:   "minimap/minimap.png",
		Settings:  settings,
	}
	interactive.CreateHeatmap(interactiveData)
}

func createLegend(data *sortedData, settings types.Settings) {
	legendData := dotplot.Legend{
		Filename:  "svg/dotplot-legend.svg",
		NumColors: 101,
		Settings:  settings,
		Title:     fmt.Sprintf("Dotplot - %s", settings.Abundance),
	}
	dotplot.CreateLegend(legendData)
}

func createTreeview(data *sortedData, clusteredData hclustData) {
	treeviewData := treeview.Data{
		Filename: "treeview/bait-prey",
		Matrix:   data.matrices.Abundance,
		Names: treeview.Names{
			Columns:         clusteredData.tree["condition"].Order,
			Rows:            clusteredData.tree["readout"].Order,
			UnsortedColumns: clusteredData.unsortedNames["condition"],
			UnsortedRows:    clusteredData.unsortedNames["readout"],
		},
		Trees: treeview.Trees{
			Column: clusteredData.dendrogram["condition"],
			Row:    clusteredData.dendrogram["readout"],
		},
	}

	treeview.Export(treeviewData)
}
