package file

import (
	"github.com/knightjdr/prohits-viz-analysis/image/minimap"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Minimap draws a minimap for the dotplot.
func Minimap(
	imageType string,
	abundance, ratios, scores [][]float64,
	sortedColumns, sortedRows []string,
	userParams typedef.Parameters,
) {
	matrices := typedef.Matrices{
		Abundance: abundance,
		Ratio:     ratios,
		Score:     scores,
	}
	parameters := typedef.Parameters{
		AbundanceCap:    userParams.AbundanceCap,
		EdgeColor:       userParams.EdgeColor,
		FillColor:       userParams.FillColor,
		InvertColor:     false,
		PrimaryFilter:   userParams.PrimaryFilter,
		SecondaryFilter: userParams.SecondaryFilter,
		ScoreType:       userParams.ScoreType,
	}
	data := minimap.Data{
		Filename:   "minimap/dotplot",
		ImageType:  "dotplot",
		Matrices:   &matrices,
		Parameters: parameters,
	}
	minimap.Write(&data)
	return
}
