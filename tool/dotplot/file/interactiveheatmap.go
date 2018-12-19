package file

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/image/interactive"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// InteractiveHeatmap creates an interactive file for a heatmap/dotplot.
func InteractiveHeatmap(
	imageType string,
	abundance, ratios, scores [][]float64,
	columns, rows []string,
	filename string,
	parameters typedef.Parameters,
) {
	matrices := typedef.Matrices{
		Abundance:  abundance,
		Conditions: columns,
		Ratio:      ratios,
		Readouts:   rows,
		Score:      scores,
	}
	parameters.XLabel = parameters.Condition
	parameters.YLabel = parameters.Readout
	interactiveData := interactive.Data{
		Filename:   fmt.Sprintf("interactive/%s.json", filename),
		ImageType:  imageType,
		Matrices:   &matrices,
		Minimap:    fmt.Sprintf("minimap/%s.png", filename),
		Parameters: parameters,
	}
	interactive.ParseHeatmap(&interactiveData)
}
