package correlation

import (
	"fmt"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/pkg/data/filter"
	"github.com/knightjdr/prohits-viz-analysis/pkg/downsample"
	"github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/pkg/matrix/convert"
	"github.com/knightjdr/prohits-viz-analysis/pkg/minimap"
	heatmapPNG "github.com/knightjdr/prohits-viz-analysis/pkg/png/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/treeview"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createBaitPreyImages(analysis *types.Analysis, conditionData, readoutData *correlationData) {
	matrices := getConditionReadoutMatrices(analysis)

	matrix, _ := hclust.Sort(matrices.Abundance, matrices.Conditions, conditionData.sortedLabels, "column")
	matrix, _ = hclust.Sort(matrix, matrices.Readouts, readoutData.sortedLabels, "row")

	createConditionReadoutSVG(matrix, conditionData.sortedLabels, readoutData.sortedLabels, analysis.Settings)
	createConditionReadoutLegend(analysis.Settings)
	createConditionReadoutPNG(matrix, analysis.Settings)
	createConditionReadoutMinimap(matrix, analysis.Settings)
	createConditionReadoutInteractive(matrix, conditionData.sortedLabels, readoutData.sortedLabels, analysis.Settings)
	createConditionReadoutTreeview(matrix, conditionData, readoutData, analysis.Settings)
}

func getConditionReadoutMatrices(analysis *types.Analysis) *types.Matrices {
	conditionReadoutAnalysis := &types.Analysis{
		Data: analysis.Data,
		Settings: types.Settings{
			MinAbundance:  analysis.Settings.ReadoutAbundanceFilter,
			PrimaryFilter: analysis.Settings.ReadoutScoreFilter,
			ScoreType:     analysis.Settings.ScoreType,
		},
	}
	filter.Process(conditionReadoutAnalysis)

	matrixSettings := convert.ConversionSettings{
		CalculateRatios: false,
		ScoreType:       analysis.Settings.ScoreType,
	}
	matrices := convert.FromTable(&analysis.Data, matrixSettings)

	return matrices
}

func createConditionReadoutSVG(matrix [][]float64, conditions, readouts []string, settings types.Settings) {
	dims := dimensions.Calculate(matrix, conditions, readouts, false)

	heatmap := svg.InitializeHeatmap()
	heatmap.AbundanceCap = settings.AbundanceCap
	heatmap.CellSize = dims.CellSize
	heatmap.Columns = conditions
	heatmap.FillColor = "blue"
	heatmap.FontSize = dims.FontSize
	heatmap.LeftMargin = dims.LeftMargin
	heatmap.Matrix = matrix
	heatmap.MinAbundance = settings.ReadoutAbundanceFilter
	heatmap.PlotHeight = dims.PlotHeight
	heatmap.PlotWidth = dims.PlotWidth
	heatmap.Rows = readouts
	heatmap.SvgHeight = dims.SvgHeight
	heatmap.SvgWidth = dims.SvgWidth
	heatmap.TopMargin = dims.TopMargin
	heatmap.XLabel = settings.Condition
	heatmap.YLabel = settings.Readout

	heatmap.Draw(fmt.Sprintf("svg/%s-%s.svg", settings.Condition, settings.Readout))
}

func createConditionReadoutLegend(settings types.Settings) {
	filename := fmt.Sprintf("%s-%s-legend", settings.Condition, settings.Readout)

	legendData := heatmap.Legend{
		Filename:  fmt.Sprintf("svg/%s.svg", filename),
		NumColors: 101,
		Settings: types.Settings{
			AbundanceCap: settings.AbundanceCap,
			FillColor:    settings.FillColor,
			InvertColor:  settings.InvertColor,
			MinAbundance: settings.ReadoutAbundanceFilter,
		},
		Title: settings.Abundance,
	}
	heatmap.CreateLegend(legendData)

	if settings.Png {
		svg.ConvertToPNG(fmt.Sprintf("svg/%s.svg", filename), fmt.Sprintf("png/%s.png", filename), "white")
	}
}

func createConditionReadoutPNG(matrix [][]float64, settings types.Settings) {
	if settings.Png {
		outfile := fmt.Sprintf("png/%s-%s.png", settings.Condition, settings.Readout)

		if downsample.Should(matrix, 0) {
			downsampled := downsample.Matrix(matrix, 0)
			dims := dimensions.Calculate(downsampled, []string{}, []string{}, false)

			heatmap := heatmapPNG.Initialize()
			heatmap.AbundanceCap = settings.AbundanceCap
			heatmap.CellSize = dims.CellSize
			heatmap.ColorSpace = "blue"
			heatmap.Height = dims.PlotHeight
			heatmap.MinAbundance = settings.ReadoutAbundanceFilter
			heatmap.Width = dims.PlotWidth

			heatmap.Draw(downsampled, outfile)
		} else {
			svg.ConvertToPNG(fmt.Sprintf("svg/%s-%s.svg", settings.Condition, settings.Readout), outfile, "white")
		}
	}
}

func createConditionReadoutMinimap(matrix [][]float64, settings types.Settings) {
	minimapData := &minimap.Data{
		DownsampleThreshold: 1000,
		Filename:            fmt.Sprintf("minimap/%s-%s.png", settings.Condition, settings.Readout),
		ImageType:           "heatmap",
		Matrices: &types.Matrices{
			Abundance: matrix,
		},
		Settings: types.Settings{
			AbundanceCap: settings.AbundanceCap,
			FillColor:    "blue",
			MinAbundance: settings.ReadoutAbundanceFilter,
		},
	}
	minimap.Create(minimapData)
}

func createConditionReadoutInteractive(matrix [][]float64, conditions, readouts []string, settings types.Settings) {
	filehandle := fmt.Sprintf("%s-%s", settings.Condition, settings.Readout)
	interactiveData := &interactive.HeatmapData{
		AnalysisType: "heatmap",
		Filename:     fmt.Sprintf("interactive/%s.json", filehandle),
		Matrices: &types.Matrices{
			Abundance:  matrix,
			Conditions: conditions,
			Readouts:   readouts,
		},
		Minimap:    fmt.Sprintf("minimap/%s.png", filehandle),
		Parameters: settings,
		Settings: map[string]interface{}{
			"abundanceCap":  settings.AbundanceCap,
			"fillColor":     "blue",
			"imageType":     "heatmap",
			"minAbundance":  settings.ReadoutAbundanceFilter,
			"primaryFilter": 0,
		},
	}
	interactiveData.Parameters.XLabel = settings.Condition
	interactiveData.Parameters.YLabel = settings.Readout

	interactive.CreateHeatmap(interactiveData)
}

func createConditionReadoutTreeview(matrix [][]float64, conditionData, readoutData *correlationData, settings types.Settings) {
	treeviewData := treeview.Data{
		Filename: fmt.Sprintf("treeview/%s-%s", settings.Condition, settings.Readout),
		Matrix:   matrix,
		Names: treeview.Names{
			Columns:         conditionData.sortedLabels,
			Rows:            readoutData.sortedLabels,
			UnsortedColumns: conditionData.labels,
			UnsortedRows:    readoutData.labels,
		},
		Trees: treeview.Trees{
			Column: conditionData.dendrogram,
			Row:    readoutData.dendrogram,
		},
	}

	treeview.Export(treeviewData)
}
