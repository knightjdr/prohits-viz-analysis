package main

import (
	"github.com/knightjdr/prohits-viz-analysis/parse"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// FormatParams converts Data struct to Parameters type.
func FormatParams(data *parse.Data) (params typedef.Parameters) {
	params = typedef.Parameters{
		AnnotationFontSize: data.AnnotationFontSize,
		EdgeColor:          data.EdgeColor,
		FillColor:          data.FillColor,
		Invert:             data.Invert,
		MarkerColor:        data.MarkerColor,
		MaximumAbundance:   data.MaximumAbundance,
		PrimaryFilter:      data.PrimaryFilter,
		ScoreType:          data.ScoreType,
		SecondaryFilter:    data.SecondaryFilter,
	}
	return
}
