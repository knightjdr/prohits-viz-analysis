package main

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/parse"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestFormatParams(t *testing.T) {
	// TEST1: returns parameter type from Data struct.
	data := parse.Data{
		AbundanceCap:    50,
		EdgeColor:       "blueBlack",
		FillColor:       "blueBlack",
		ImageType:       "dotplot",
		InvertColor:     false,
		MinAbundance:    5,
		PrimaryFilter:   0.01,
		SecondaryFilter: 0.05,
		ScoreType:       "lte",
	}
	wantedParams := typedef.Parameters{
		AbundanceCap:    50,
		EdgeColor:       "blueBlack",
		FillColor:       "blueBlack",
		InvertColor:     false,
		MinAbundance:    5,
		PrimaryFilter:   0.01,
		SecondaryFilter: 0.05,
		ScoreType:       "lte",
	}
	actualParams := FormatParams(&data)
	assert.EqualValues(t, wantedParams, actualParams)
}
