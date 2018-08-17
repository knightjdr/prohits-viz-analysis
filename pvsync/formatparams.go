package main

import (
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// FormatParams converts Data struct to Parameters type.
func FormatParams(data *Data) (params typedef.Parameters) {
	params = typedef.Parameters{
		EdgeColor:        data.EdgeColor,
		FillColor:        data.FillColor,
		Invert:           data.Invert,
		MaximumAbundance: data.MaximumAbundance,
		PrimaryFilter:    data.PrimaryFilter,
		ScoreType:        data.ScoreType,
		SecondaryFilter:  data.SecondaryFilter,
	}
	return
}
