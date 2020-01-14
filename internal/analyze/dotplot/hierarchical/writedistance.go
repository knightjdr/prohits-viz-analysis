package hierarchical

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/minimap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/treeview"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func writeDistance(data *sortedData, clusteredData hclustData, settings types.Settings) {
	if settings.WriteDistance {
		createDistanceSVG(data.conditionDist, data.matrices.Conditions, settings, settings.Condition)
		createDistanceSVG(data.readoutDist, data.matrices.Readouts, settings, settings.Readout)

		createDistanceLegend(settings, settings.Condition)
		createDistanceLegend(settings, settings.Readout)

		createDistanceMinimap(data.conditionDist, settings, "condition")
		createDistanceMinimap(data.readoutDist, settings, "readout")

		createDistanceInteractive(data.conditionDist, data.matrices.Conditions, settings, "condition")
		createDistanceInteractive(data.readoutDist, data.matrices.Readouts, settings, "readout")

		createConditionDistanceTreeview(data, clusteredData, settings)
		createReadoutDistanceTreeview(data, clusteredData, settings)
	}
}

func createDistanceSVG(matrix [][]float64, labels []string, settings types.Settings, filehandle string) {
	dims := dimensions.Calculate(matrix, labels, labels, false)
	filename := fmt.Sprintf("svg/%[1]s-%[1]s.svg", filehandle)

	heatmap := svg.InitializeHeatmap()
	heatmap.AbundanceCap = 1
	heatmap.CellSize = dims.CellSize
	heatmap.Columns = labels
	heatmap.FillColor = settings.FillColor
	heatmap.FontSize = dims.FontSize
	heatmap.Invert = true
	heatmap.LeftMargin = dims.LeftMargin
	heatmap.Matrix = matrix
	heatmap.MinAbundance = 0
	heatmap.PlotHeight = dims.PlotHeight
	heatmap.PlotWidth = dims.PlotWidth
	heatmap.Rows = labels
	heatmap.SvgHeight = dims.SvgHeight
	heatmap.SvgWidth = dims.SvgWidth
	heatmap.TopMargin = dims.TopMargin
	heatmap.XLabel = settings.Condition
	heatmap.YLabel = settings.Condition

	heatmap.Draw(filename)
}

func createDistanceLegend(settings types.Settings, title string) {
	legendData := heatmap.Legend{
		Filename:  fmt.Sprintf("svg/%s-distance-legend.svg", title),
		NumColors: 101,
		Settings: types.Settings{
			AbundanceCap: 1,
			FillColor:    settings.FillColor,
			InvertColor:  true,
			MinAbundance: 0,
		},
		Title: fmt.Sprintf("Distance - %s", title),
	}
	heatmap.CreateLegend(legendData)
}

func createDistanceMinimap(matrix [][]float64, settings types.Settings, title string) {
	minimapData := &minimap.Data{
		DownsampleThreshold: 1000,
		Filename:            fmt.Sprintf("minimap/%s.png", title),
		ImageType:           "heatmap",
		Matrices: &types.Matrices{
			Abundance: matrix,
		},
		Settings: types.Settings{
			AbundanceCap: 1,
			FillColor:    settings.FillColor,
			InvertColor:  true,
			MinAbundance: 0,
		},
	}
	minimap.Create(minimapData)
}

func createDistanceInteractive(matrix [][]float64, labels []string, settings types.Settings, title string) {
	interactiveData := &interactive.HeatmapData{
		Filename:  fmt.Sprintf("interactive/%[1]s-%[1]s.json", title),
		ImageType: "heatmap",
		Matrices: &types.Matrices{
			Abundance:  matrix,
			Conditions: labels,
			Readouts:   labels,
		},
		Minimap:  fmt.Sprintf("minimap/%s.png", title),
		Settings: settings,
	}
	interactiveData.Settings.AbundanceCap = 1
	interactiveData.Settings.InvertColor = true
	interactiveData.Settings.MinAbundance = 0

	interactive.CreateHeatmap(interactiveData)
}

func createConditionDistanceTreeview(data *sortedData, clusteredData hclustData, settings types.Settings) {
	treeviewData := treeview.Data{
		Filename: fmt.Sprintf("treeview/%[1]s-%[1]s", settings.Condition),
		Matrix:   data.conditionDist,
		Names: treeview.Names{
			Columns:         clusteredData.tree["condition"].Order,
			Rows:            clusteredData.tree["condition"].Order,
			UnsortedColumns: clusteredData.unsortedNames["condition"],
			UnsortedRows:    clusteredData.unsortedNames["condition"],
		},
		Trees: treeview.Trees{
			Column: clusteredData.dendrogram["condition"],
			Row:    clusteredData.dendrogram["condition"],
		},
	}

	treeview.Export(treeviewData)
}

func createReadoutDistanceTreeview(data *sortedData, clusteredData hclustData, settings types.Settings) {
	treeviewData := treeview.Data{
		Filename: fmt.Sprintf("treeview/%[1]s-%[1]s", settings.Readout),
		Matrix:   data.readoutDist,
		Names: treeview.Names{
			Columns:         clusteredData.tree["readout"].Order,
			Rows:            clusteredData.tree["readout"].Order,
			UnsortedColumns: clusteredData.unsortedNames["readout"],
			UnsortedRows:    clusteredData.unsortedNames["readout"],
		},
		Trees: treeview.Trees{
			Column: clusteredData.dendrogram["readout"],
			Row:    clusteredData.dendrogram["readout"],
		},
	}

	treeview.Export(treeviewData)
}
