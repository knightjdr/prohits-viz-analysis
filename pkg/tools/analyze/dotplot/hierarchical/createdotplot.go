package hierarchical

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/pkg/minimap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/dotplot"
	"github.com/knightjdr/prohits-viz-analysis/pkg/treeview"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// CreateDotplot image.
func CreateDotplot(data *SortedData, clusteredData HclustData, settings types.Settings) {
	createDotplotSVG(data, settings)
	createDotplotLegend(data, settings)
	createDotplotMinimap(data, settings)
	createDotplotInteractive(data, settings)
	createDotplotTreeview(data, clusteredData, settings)
}

func createDotplotSVG(data *SortedData, settings types.Settings) {
	dims := dimensions.Calculate(data.Matrices.Abundance, data.Matrices.Conditions, data.Matrices.Readouts, false)

	dotplot := svg.InitializeDotplot()
	dotplot.CellSize = dims.CellSize
	dotplot.Columns = data.Matrices.Conditions
	dotplot.EdgeColor = settings.EdgeColor
	dotplot.FillColor = settings.FillColor
	dotplot.FillMax = settings.FillMax
	dotplot.FillMin = settings.FillMin
	dotplot.FontSize = dims.FontSize
	dotplot.Invert = settings.InvertColor
	dotplot.LeftMargin = dims.LeftMargin
	dotplot.Matrices = data.Matrices
	dotplot.PlotHeight = dims.PlotHeight
	dotplot.PlotWidth = dims.PlotWidth
	dotplot.PrimaryFilter = settings.PrimaryFilter
	dotplot.Ratio = dims.Ratio
	dotplot.Rows = data.Matrices.Readouts
	dotplot.ScoreType = settings.ScoreType
	dotplot.SecondaryFilter = settings.SecondaryFilter
	dotplot.SvgHeight = dims.SvgHeight
	dotplot.SvgWidth = dims.SvgWidth
	dotplot.TopMargin = dims.TopMargin
	dotplot.XLabel = settings.Condition
	dotplot.YLabel = settings.Readout

	dotplot.Draw("svg/dotplot.svg")
}

func createDotplotLegend(data *SortedData, settings types.Settings) {
	legendData := dotplot.Legend{
		Filename:  "svg/dotplot-legend.svg",
		NumColors: 101,
		Settings:  settings,
		Title:     fmt.Sprintf("Dotplot - %s", settings.Abundance),
	}
	dotplot.CreateLegend(legendData)
}

func createDotplotMinimap(data *SortedData, settings types.Settings) {
	minimapData := &minimap.Data{
		DownsampleThreshold: 1000,
		Filename:            "minimap/dotplot.png",
		ImageType:           "dotplot",
		Matrices:            data.Matrices,
		Settings:            settings,
	}
	minimap.Create(minimapData)
}

func createDotplotInteractive(data *SortedData, settings types.Settings) {
	interactiveData := &interactive.HeatmapData{
		AnalysisType: "dotplot",
		Filename:     "interactive/dotplot.json",
		Matrices:     data.Matrices,
		Minimap:      "minimap/dotplot.png",
		Parameters:   settings,
		Settings: map[string]interface{}{
			"abundanceCap":    settings.AbundanceCap,
			"abundanceType":   settings.AbundanceType,
			"edgeColor":       settings.EdgeColor,
			"fillColor":       settings.FillColor,
			"fillMax":         settings.FillMax,
			"fillMin":         settings.FillMin,
			"imageType":       "dotplot",
			"invertColor":     settings.InvertColor,
			"minAbundance":    settings.MinAbundance,
			"primaryFilter":   settings.PrimaryFilter,
			"secondaryFilter": settings.SecondaryFilter,
		},
	}
	interactiveData.Parameters.XLabel = settings.Condition
	interactiveData.Parameters.YLabel = settings.Readout

	interactive.CreateHeatmap(interactiveData)
}

func createDotplotTreeview(data *SortedData, clusteredData HclustData, settings types.Settings) {
	treeviewData := treeview.Data{
		Filename: fmt.Sprintf("treeview/%s-%s", settings.Condition, settings.Readout),
		Matrix:   data.Matrices.Abundance,
		Names: treeview.Names{
			Columns:         clusteredData.Tree["condition"].Order,
			Rows:            clusteredData.Tree["readout"].Order,
			UnsortedColumns: clusteredData.UnsortedNames["condition"],
			UnsortedRows:    clusteredData.UnsortedNames["readout"],
		},
		Trees: treeview.Trees{
			Column: clusteredData.Dendrogram["condition"],
			Row:    clusteredData.Dendrogram["readout"],
		},
	}

	treeview.Export(treeviewData)
}
