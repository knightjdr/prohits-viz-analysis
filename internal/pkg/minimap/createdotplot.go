package minimap

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/convert"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/dotplot"
)

var convertSVG = convert.RSVG

func createDotplot(data *Data, dims *dimensions.Heatmap) {
	image := dotplot.Initialize()
	image.AbundanceCap = data.Settings.AbundanceCap
	image.BoundingBox = false
	image.CellSize = dims.CellSize
	image.EdgeColor = data.Settings.EdgeColor
	image.FillColor = data.Settings.FillColor
	image.Invert = data.Settings.InvertColor
	image.Matrices = data.Matrices
	image.MinAbundance = data.Settings.MinAbundance
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
	convertSVG(svgFileName, data.Filename)
}
