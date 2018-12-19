/*Package file in dotplot contains functions for creating image and log files */
package file

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/image/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Distance draws an svg for a distance heatmap.
func Distance(dist [][]float64, names []string, label, filename string, userParams typedef.Parameters) {
	matrices := typedef.Matrices{
		Abundance:  dist,
		Conditions: names,
		Readouts:   names,
	}
	parameters := userParams
	parameters.AbundanceCap = float64(1)
	parameters.XLabel = label
	parameters.InvertColor = true
	parameters.YLabel = label

	file := fmt.Sprintf("svg/%s.svg", filename)
	data := svg.Data{
		Filename:   file,
		ImageType:  "heatmap",
		Matrices:   &matrices,
		Parameters: parameters,
	}

	svg.Heatmap(&data)
	return
}
