package minimap

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
)

var convertSVG = svg.ConvertToPNG

func createDotplot(data *Data) {
	dims := dimensions.Calculate(data.Matrices.Abundance, []string{}, []string{}, true)

	image := svg.InitializeDotplot()
	image.BoundingBox = false
	image.CellSize = dims.CellSize
	image.EdgeColor = data.Settings.EdgeColor
	image.FillColor = data.Settings.FillColor
	image.FillMax = data.Settings.FillMax
	image.FillMin = data.Settings.FillMin
	image.Invert = data.Settings.InvertColor
	image.Matrices = data.Matrices
	image.PlotHeight = dims.PlotHeight
	image.PlotWidth = dims.PlotWidth
	image.PrimaryFilter = data.Settings.PrimaryFilter
	image.Ratio = dims.Ratio
	image.ScoreType = data.Settings.ScoreType
	image.SecondaryFilter = data.Settings.SecondaryFilter
	image.SvgHeight = dims.SvgHeight
	image.SvgWidth = dims.SvgWidth

	filename := strings.TrimSuffix(data.Filename, filepath.Ext(data.Filename))
	svgFileName := fmt.Sprintf("%s.svg", filename)

	image.Draw(svgFileName)
	convertSVG(svgFileName, data.Filename, "none")
}
