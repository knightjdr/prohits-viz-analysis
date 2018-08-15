package main

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestFormatParams(t *testing.T) {
	// TEST1: returns parameter type from Data struct.
	data := Data{
		EdgeColor:        "blueBlack",
		FillColor:        "blueBlack",
		ImageType:        "dotplot",
		Invert:           false,
		MaximumAbundance: 50,
		PrimaryFilter:    0.01,
		SecondaryFilter:  0.05,
		ScoreType:        "lte",
	}
	wantedParams := typedef.Parameters{
		EdgeColor:        "blueBlack",
		FillColor:        "blueBlack",
		Invert:           false,
		MaximumAbundance: 50,
		PrimaryFilter:    0.01,
		SecondaryFilter:  0.05,
		ScoreType:        "lte",
	}
	actualParams := FormatParams(&data)
	assert.EqualValues(t, wantedParams, actualParams)
}
