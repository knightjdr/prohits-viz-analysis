package hierarchical

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/png/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/spf13/afero"
)

// CreatePNGs as output format.
func CreatePNGs(data *SortedData, clusteredData HclustData, settings types.Settings) {
	if settings.Png {
		createDotplotPNG(data.Matrices.Abundance, settings)
		createHeatmapPNG(data.Matrices.Abundance, settings)
		createDistancePNG(data, settings)
		convertLegends(settings)
	}
}

func createDotplotPNG(matrix [][]float64, settings types.Settings) {
	writeHeatmapPNG(matrix, settings, "dotplot")
}

func createHeatmapPNG(matrix [][]float64, settings types.Settings) {
	if settings.WriteHeatmap {
		writeHeatmapPNG(matrix, settings, "heatmap")
	}
}

func createDistancePNG(data *SortedData, settings types.Settings) {
	if settings.WriteDistance {
		distanceSettings := types.Settings{
			AbundanceCap: 1,
			FillColor:    settings.FillColor,
			InvertColor:  true,
			MinAbundance: 0,
		}
		writeHeatmapPNG(data.ConditionDist, distanceSettings, fmt.Sprintf("%[1]s-%[1]s", settings.Condition))
		writeHeatmapPNG(data.ReadoutDist, distanceSettings, fmt.Sprintf("%[1]s-%[1]s", settings.Readout))
	}
}

func writeHeatmapPNG(matrix [][]float64, settings types.Settings, filehandle string) {
	if len(matrix) == 0 {
		return
	}

	if downsample.Should(matrix, 0) {
		downsampled := downsample.Matrix(matrix, 0)
		dims := dimensions.Calculate(downsampled, []string{}, []string{}, false)

		heatmap := heatmap.Initialize()
		heatmap.AbundanceCap = settings.AbundanceCap
		heatmap.CellSize = dims.CellSize
		heatmap.ColorSpace = settings.FillColor
		heatmap.Height = dims.PlotHeight
		heatmap.Invert = settings.InvertColor
		heatmap.MinAbundance = settings.MinAbundance
		heatmap.Width = dims.PlotWidth

		filename := fmt.Sprintf("png/%s.png", filehandle)
		heatmap.Draw(downsampled, filename)
	} else {
		svg.ConvertToPNG(fmt.Sprintf("svg/%s.svg", filehandle), fmt.Sprintf("png/%s.png", filehandle), "white")
	}
}

func convertLegends(settings types.Settings) {
	svg.ConvertToPNG("svg/dotplot-legend.svg", "png/dotplot-legend.png", "white")

	if settings.WriteHeatmap {
		svg.ConvertToPNG("svg/heatmap-legend.svg", "png/heatmap-legend.png", "white")
	}

	if settings.WriteDistance {
		conditionFileName := fmt.Sprintf("%s-distance-legend", settings.Condition)
		conditionFilePath := fmt.Sprintf("svg/%s.svg", conditionFileName)
		exists, _ := afero.Exists(fs.Instance, conditionFilePath)
		if exists {
			svg.ConvertToPNG(conditionFilePath, fmt.Sprintf("png/%s.png", conditionFileName), "white")
		}

		readoutFileName := fmt.Sprintf("%s-distance-legend", settings.Readout)
		readoutFilePath := fmt.Sprintf("svg/%s.svg", readoutFileName)
		exists, _ = afero.Exists(fs.Instance, readoutFilePath)
		if exists {
			svg.ConvertToPNG(readoutFilePath, fmt.Sprintf("png/%s.png", readoutFileName), "white")
		}
	}
}
