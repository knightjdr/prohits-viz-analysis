package correlation

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/minimap"
	heatmapPNG "github.com/knightjdr/prohits-viz-analysis/internal/pkg/png/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/treeview"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func createCorrelationImages(conditionData, readoutData *correlationData, settings types.Settings) {
	createCorrelationSVG(conditionData, settings, settings.Condition)
	createCorrelationSVG(readoutData, settings, settings.Readout)

	createCorrelationLegend(settings, settings.Condition)
	createCorrelationLegend(settings, settings.Readout)

	createPNG(conditionData.matrix, settings, settings.Condition)
	createPNG(readoutData.matrix, settings, settings.Readout)

	createCorrelationMinimap(conditionData, settings, settings.Condition)
	createCorrelationMinimap(readoutData, settings, settings.Readout)

	createCorrelationInteractive(conditionData, settings, settings.Condition)
	createCorrelationInteractive(readoutData, settings, settings.Readout)

	createCorrelationTreeview(conditionData, settings, settings.Condition)
	createCorrelationTreeview(readoutData, settings, settings.Readout)
}

func createCorrelationSVG(data *correlationData, settings types.Settings, label string) {
	dims := dimensions.Calculate(data.matrix, data.sortedLabels, data.sortedLabels, false)

	heatmap := svg.InitializeHeatmap()
	heatmap.AbundanceCap = 1
	heatmap.CellSize = dims.CellSize
	heatmap.Columns = data.sortedLabels
	heatmap.FillColor = settings.FillColor
	heatmap.FontSize = dims.FontSize
	heatmap.Invert = settings.InvertColor
	heatmap.LeftMargin = dims.LeftMargin
	heatmap.Matrix = data.matrix
	heatmap.MinAbundance = -1
	heatmap.PlotHeight = dims.PlotHeight
	heatmap.PlotWidth = dims.PlotWidth
	heatmap.Rows = data.sortedLabels
	heatmap.SvgHeight = dims.SvgHeight
	heatmap.SvgWidth = dims.SvgWidth
	heatmap.TopMargin = dims.TopMargin
	heatmap.XLabel = label
	heatmap.YLabel = label

	heatmap.Draw(fmt.Sprintf("svg/%[1]s-%[1]s.svg", label))
}

func createCorrelationLegend(settings types.Settings, label string) {
	filename := fmt.Sprintf("%[1]s-%[1]s-legend", label)

	legendData := heatmap.Legend{
		Filename:  fmt.Sprintf("svg/%s.svg", filename),
		NumColors: 101,
		Settings: types.Settings{
			AbundanceCap: 1,
			FillColor:    settings.FillColor,
			InvertColor:  settings.InvertColor,
			MinAbundance: -1,
		},
		Title: fmt.Sprintf("Correlation - %s", label),
	}
	heatmap.CreateLegend(legendData)

	if settings.Png {
		svg.ConvertToPNG(fmt.Sprintf("svg/%s.svg", filename), fmt.Sprintf("png/%s.png", filename), "white")
	}
}

func createPNG(matrix [][]float64, settings types.Settings, label string) {
	outfile := fmt.Sprintf("png/%[1]s-%[1]s.png", label)

	if downsample.Should(matrix, 0) {
		downsampled := downsample.Matrix(matrix, 0)
		dims := dimensions.Calculate(downsampled, []string{}, []string{}, false)

		heatmap := heatmapPNG.Initialize()
		heatmap.AbundanceCap = 1
		heatmap.CellSize = dims.CellSize
		heatmap.ColorSpace = settings.FillColor
		heatmap.Height = dims.PlotHeight
		heatmap.Invert = settings.InvertColor
		heatmap.MinAbundance = -1
		heatmap.Width = dims.PlotWidth

		heatmap.Draw(downsampled, outfile)
	} else {
		svg.ConvertToPNG(fmt.Sprintf("svg/%[1]s-%[1]s.svg", label), outfile, "white")
	}
}

func createCorrelationMinimap(data *correlationData, settings types.Settings, label string) {
	minimapData := &minimap.Data{
		DownsampleThreshold: 1000,
		Filename:            fmt.Sprintf("minimap/%[1]s-%[1]s.png", label),
		ImageType:           "heatmap",
		Matrices: &types.Matrices{
			Abundance: data.matrix,
		},
		Settings: types.Settings{
			AbundanceCap: 1,
			FillColor:    settings.FillColor,
			InvertColor:  settings.InvertColor,
			MinAbundance: -1,
		},
	}
	minimap.Create(minimapData)
}

func createCorrelationInteractive(data *correlationData, settings types.Settings, label string) {
	interactiveData := &interactive.HeatmapData{
		AnalysisType: "correlation",
		Filename:     fmt.Sprintf("interactive/%[1]s-%[1]s.json", label),
		Matrices: &types.Matrices{
			Abundance:  data.matrix,
			Conditions: data.sortedLabels,
			Readouts:   data.sortedLabels,
		},
		Minimap:    fmt.Sprintf("minimap/%[1]s-%[1]s.png", label),
		Parameters: settings,
		Settings: map[string]interface{}{
			"abundanceCap":  1,
			"fillColor":     settings.FillColor,
			"imageType":     "heatmap",
			"invertColor":   settings.InvertColor,
			"minAbundance":  -1,
			"primaryFilter": 0,
		},
	}
	interactiveData.Parameters.XLabel = label
	interactiveData.Parameters.YLabel = label

	interactive.CreateHeatmap(interactiveData)
}

func createCorrelationTreeview(data *correlationData, settings types.Settings, label string) {
	treeviewData := treeview.Data{
		Filename: fmt.Sprintf("treeview/%[1]s-%[1]s", label),
		Matrix:   data.matrix,
		Names: treeview.Names{
			Columns:         data.sortedLabels,
			Rows:            data.sortedLabels,
			UnsortedColumns: data.labels,
			UnsortedRows:    data.labels,
		},
		Trees: treeview.Trees{
			Column: data.dendrogram,
			Row:    data.dendrogram,
		},
	}

	treeview.Export(treeviewData)
}
