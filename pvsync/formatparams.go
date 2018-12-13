package main

import (
	"github.com/knightjdr/prohits-viz-analysis/parse"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// FormatParams converts Data struct to Parameters type.
func FormatParams(data *parse.Data) (parameters typedef.Parameters) {
	parameters = typedef.Parameters{
		AbundanceCap:    data.AbundanceCap,
		EdgeColor:       data.EdgeColor,
		FillColor:       data.FillColor,
		InvertColor:     data.InvertColor,
		MinAbundance:    data.MinAbundance,
		PrimaryFilter:   data.PrimaryFilter,
		ScoreType:       data.ScoreType,
		SecondaryFilter: data.SecondaryFilter,
	}
	return
}
