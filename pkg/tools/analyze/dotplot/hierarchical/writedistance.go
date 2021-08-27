package hierarchical

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/pkg/minimap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/treeview"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// WriteDistance images.
func WriteDistance(data *SortedData, clusteredData HclustData, settings types.Settings) {
	if settings.WriteDistance {
		if len(data.ConditionDist) > 0 {
			createDistanceSVG(data.ConditionDist, data.Matrices.Conditions, settings, settings.Condition)
			createDistanceLegend(settings, settings.Condition)
			createDistanceMinimap(data.ConditionDist, settings, settings.Condition)
			createDistanceInteractive(data.ConditionDist, data.Matrices.Conditions, settings, settings.Condition)
			createConditionDistanceTreeview(data, clusteredData, settings)
		}
		if len(data.ReadoutDist) > 0 {
			createDistanceSVG(data.ReadoutDist, data.Matrices.Readouts, settings, settings.Readout)
			createDistanceLegend(settings, settings.Readout)
			createDistanceMinimap(data.ReadoutDist, settings, settings.Readout)
			createDistanceInteractive(data.ReadoutDist, data.Matrices.Readouts, settings, settings.Readout)
			createReadoutDistanceTreeview(data, clusteredData, settings)
		}
	}
}

func createDistanceSVG(matrix [][]float64, labels []string, settings types.Settings, filehandle string) {
	dims := dimensions.Calculate(matrix, labels, labels, false)
	filename := fmt.Sprintf("svg/%[1]s-%[1]s.svg", filehandle)

	heatmap := svg.InitializeHeatmap()
	heatmap.CellSize = dims.CellSize
	heatmap.Columns = labels
	heatmap.FillColor = "blue"
	heatmap.FillMax = 1
	heatmap.FillMin = 0
	heatmap.FontSize = dims.FontSize
	heatmap.Invert = true
	heatmap.LeftMargin = dims.LeftMargin
	heatmap.Matrix = matrix
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
			FillColor:   "blue",
			FillMax:     1,
			FillMin:     0,
			InvertColor: true,
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
			FillColor:   "blue",
			FillMax:     1,
			FillMin:     0,
			InvertColor: true,
		},
	}
	minimap.Create(minimapData)
}

func createDistanceInteractive(matrix [][]float64, labels []string, settings types.Settings, title string) {
	interactiveData := &interactive.HeatmapData{
		AnalysisType: "heatmap",
		Filename:     fmt.Sprintf("interactive/%[1]s-%[1]s.json", title),
		Matrices: &types.Matrices{
			Abundance:  matrix,
			Conditions: labels,
			Readouts:   labels,
		},
		Minimap:    fmt.Sprintf("minimap/%s.png", title),
		Parameters: settings,
		Settings: map[string]interface{}{
			"abundanceCap":  1,
			"abundanceType": "positive",
			"fillColor":     "blue",
			"fillMax":       1,
			"fillMin":       0,
			"imageType":     "heatmap",
			"invertColor":   true,
			"minAbundance":  0,
			"primaryFilter": 0,
		},
	}
	interactiveData.Parameters.XLabel = title
	interactiveData.Parameters.YLabel = title

	interactive.CreateHeatmap(interactiveData)
}

func createConditionDistanceTreeview(data *SortedData, clusteredData HclustData, settings types.Settings) {
	treeviewData := treeview.Data{
		Filename: fmt.Sprintf("treeview/%[1]s-%[1]s", settings.Condition),
		Matrix:   data.ConditionDist,
		Names: treeview.Names{
			Columns:         clusteredData.Tree["condition"].Order,
			Rows:            clusteredData.Tree["condition"].Order,
			UnsortedColumns: clusteredData.UnsortedNames["condition"],
			UnsortedRows:    clusteredData.UnsortedNames["condition"],
		},
		Trees: treeview.Trees{
			Column: clusteredData.Dendrogram["condition"],
			Row:    clusteredData.Dendrogram["condition"],
		},
	}

	treeview.Export(treeviewData)
}

func createReadoutDistanceTreeview(data *SortedData, clusteredData HclustData, settings types.Settings) {
	treeviewData := treeview.Data{
		Filename: fmt.Sprintf("treeview/%[1]s-%[1]s", settings.Readout),
		Matrix:   data.ReadoutDist,
		Names: treeview.Names{
			Columns:         clusteredData.Tree["readout"].Order,
			Rows:            clusteredData.Tree["readout"].Order,
			UnsortedColumns: clusteredData.UnsortedNames["readout"],
			UnsortedRows:    clusteredData.UnsortedNames["readout"],
		},
		Trees: treeview.Trees{
			Column: clusteredData.Dendrogram["readout"],
			Row:    clusteredData.Dendrogram["readout"],
		},
	}

	treeview.Export(treeviewData)
}
