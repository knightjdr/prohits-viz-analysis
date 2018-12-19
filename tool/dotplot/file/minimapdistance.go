package file

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/image/minimap"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// MinimapDistance draws a distance heatmap.
func MinimapDistance(
	abundance [][]float64,
	sorted []string,
	fillColor string,
	imageName string,
) {
	matrices := typedef.Matrices{
		Abundance: abundance,
	}
	// Define dotplot parameters.
	parameters := typedef.Parameters{
		AbundanceCap: float64(1),
		FillColor:    fillColor,
		InvertColor:  true,
	}

	data := minimap.Data{
		Filename:   fmt.Sprintf("minimap/%s", imageName),
		ImageType:  "heatmap",
		Matrices:   &matrices,
		Parameters: parameters,
	}
	minimap.Write(&data)
	return
}
