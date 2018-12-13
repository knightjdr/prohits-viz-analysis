package main

import (
	"github.com/knightjdr/prohits-viz-analysis/parse"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// FormatParams converts Data struct to Parameters type.
func FormatParams(data *parse.Data) (parameters typedef.Parameters) {
	parameters = typedef.Parameters{
		AbundanceCap:    data.AbundanceCap,
		Condition:       data.Condition,
		EdgeColor:       data.EdgeColor,
		FillColor:       data.FillColor,
		InvertColor:     data.InvertColor,
		MinAbundance:    data.MinAbundance,
		PrimaryFilter:   data.PrimaryFilter,
		Readout:         data.Readout,
		ScoreType:       data.ScoreType,
		SecondaryFilter: data.SecondaryFilter,
	}
	if parameters.Condition == "" {
		parameters.Condition = "Conditions"
	}
	if parameters.Readout == "" {
		parameters.Readout = "Readouts"
	}
	return
}
