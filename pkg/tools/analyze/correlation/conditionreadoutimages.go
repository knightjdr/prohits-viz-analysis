package correlation

import (
	"fmt"
	"math"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/pkg/data/filter"
	"github.com/knightjdr/prohits-viz-analysis/pkg/downsample"
	heatmapColor "github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/color"
	"github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/pkg/matrix/convert"
	matrixMath "github.com/knightjdr/prohits-viz-analysis/pkg/matrix/math"
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

	settings := adjustConditionReadoutSettings(analysis.Settings, matrix)

	createConditionReadoutSVG(matrix, conditionData.sortedLabels, readoutData.sortedLabels, settings)
	createConditionReadoutLegend(settings, matrix)
	createConditionReadoutPNG(matrix, settings)
	createConditionReadoutMinimap(matrix, settings)
	createConditionReadoutInteractive(matrix, conditionData.sortedLabels, readoutData.sortedLabels, settings)
	createConditionReadoutTreeview(matrix, conditionData, readoutData, settings)
}

func adjustConditionReadoutSettings(settings types.Settings, matrix [][]float64) types.Settings {
	adjusted := settings

	min, max := matrixMath.MinMax(matrix)
	if math.Abs(min) > max {
		max = math.Abs(min)
	}
	abundanceCap := math.Ceil(max)

	defaultCap := float64(50)
	if abundanceCap > defaultCap {
		abundanceCap = defaultCap
	}

	adjusted.AbundanceType = matrixMath.DefineValues(matrix)
	adjusted.AbundanceCap = abundanceCap
	adjusted.AutomaticallySetFill = true
	adjusted.MinAbundance = settings.ReadoutAbundanceFilter
	adjusted.PrimaryFilter = settings.ReadoutScoreFilter

	heatmapColor.SetFillLimits(&adjusted)
	heatmapColor.AdjustFillColor(&adjusted)

	return adjusted
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
	heatmap.CellSize = dims.CellSize
	heatmap.Columns = conditions
	heatmap.FillColor = settings.FillColor
	heatmap.FillMax = settings.FillMax
	heatmap.FillMin = settings.FillMin
	heatmap.FontSize = dims.FontSize
	heatmap.Invert = settings.InvertColor
	heatmap.LeftMargin = dims.LeftMargin
	heatmap.Matrix = matrix
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

func createConditionReadoutLegend(settings types.Settings, matrix [][]float64) {
	filename := fmt.Sprintf("%s-%s-legend", settings.Condition, settings.Readout)

	legendData := heatmap.Legend{
		Filename:  fmt.Sprintf("svg/%s.svg", filename),
		NumColors: 101,
		Settings: types.Settings{
			FillColor:   settings.FillColor,
			FillMax:     settings.FillMax,
			FillMin:     settings.FillMin,
			InvertColor: settings.InvertColor,
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
			heatmap.CellSize = dims.CellSize
			heatmap.ColorSpace = settings.FillColor
			heatmap.FillMax = settings.FillMax
			heatmap.FillMin = settings.FillMin
			heatmap.Invert = settings.InvertColor
			heatmap.Height = dims.PlotHeight
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
			FillColor:   settings.FillColor,
			FillMax:     settings.FillMax,
			FillMin:     settings.FillMin,
			InvertColor: settings.InvertColor,
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
			"abundanceType": settings.AbundanceType,
			"fillColor":     settings.FillColor,
			"imageType":     "heatmap",
			"invertColor":   settings.InvertColor,
			"minAbundance":  settings.MinAbundance,
			"primaryFilter": settings.PrimaryFilter,
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
