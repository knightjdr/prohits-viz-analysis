package file

import (
	"github.com/knightjdr/prohits-viz-analysis/image/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Heatmap creates a condition readout heatmap.
func Heatmap(imageType string, abundance, ratios, scores [][]float64, columns, rows []string, userParams typedef.Parameters) {
	matrices := typedef.Matrices{
		Abundance:  abundance,
		Conditions: columns,
		Ratio:      ratios,
		Readouts:   rows,
		Score:      scores,
	}
	parameters := userParams
	parameters.XLabel = userParams.Condition
	parameters.YLabel = userParams.Readout

	data := svg.Data{
		Filename:   "svg/dotplot.svg",
		ImageType:  imageType,
		Matrices:   &matrices,
		Parameters: parameters,
	}
	svg.Heatmap(&data)
	return
}
