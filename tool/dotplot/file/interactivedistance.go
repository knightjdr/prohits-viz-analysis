package file

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/image/interactive"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// InteractiveDistance creates and interactive file for a distance matrix.
func InteractiveDistance(dist [][]float64, names []string, label, filename string, userParams typedef.Parameters) {
	matrices := typedef.Matrices{
		Abundance:  dist,
		Conditions: names,
		Readouts:   names,
	}
	parameters := userParams
	parameters.AbundanceCap = 1
	parameters.InvertColor = true
	parameters.MinAbundance = 0
	parameters.XLabel = label
	parameters.YLabel = label
	interactiveData := interactive.Data{
		Filename:   fmt.Sprintf("interactive/%s.json", filename),
		ImageType:  "heatmap",
		Matrices:   &matrices,
		Minimap:    fmt.Sprintf("minimap/%s.png", filename),
		Parameters: parameters,
	}
	interactive.ParseHeatmap(&interactiveData)
}
