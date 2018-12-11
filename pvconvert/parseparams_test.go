package main

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestInvertColorToBool(t *testing.T) {
	// Test
	assert.True(
		t,
		invertColorToBool(1),
		"1 not converted to true",
	)

	// Test
	assert.False(
		t,
		invertColorToBool(0),
		"0 not converted to false",
	)
}

func TestScoreTypeToBool(t *testing.T) {
	// Test
	assert.Equal(
		t,
		"gte",
		scoreTypeToBool(1),
		"1 not converted to gte",
	)

	// Test
	assert.Equal(
		t,
		"lte",
		scoreTypeToBool(0),
		"0 not converted to lte",
	)
}

func TestInferSettings(t *testing.T) {
	// Test dotplot with non-negative values
	csv := []map[string]string{
		{"value": "5"}, {"value": "10"}, {"value": "100"}, {"value": "34"},
	}
	inferred := inferSettings(csv, "dotplot")
	assert.Equal(t, float64(50), inferred.AbundanceCap, "Inferred dotplot abundance cap does not match expected")
	assert.Equal(t, "blueBlack", inferred.EdgeColor, "Inferred dotplot edge color does not match expected")
	assert.Equal(t, "blueBlack", inferred.FillColor, "Inferred dotplot fill color does not match expected")
	assert.Equal(t, float64(0), inferred.MinAbundance, "Inferred dotplot min abundance does not match expected")

	// Test dotplot with negative values
	csv = []map[string]string{
		{"value": "-5.5"}, {"value": "10"}, {"value": "34.3"}, {"value": "-3"},
	}
	inferred = inferSettings(csv, "dotplot")
	assert.Equal(t, float64(35), inferred.AbundanceCap, "Inferred dotplot abundance cap does not match expected")
	assert.Equal(t, "blueBlack", inferred.EdgeColor, "Inferred dotplot edge color does not match expected")
	assert.Equal(t, "redBlue", inferred.FillColor, "Inferred dotplot fill color does not match expected")
	assert.Equal(t, float64(-6), inferred.MinAbundance, "Inferred dotplot min abundance does not match expected")

	// Test heatmap with non-negative values
	csv = []map[string]string{
		{"value": "5"}, {"value": "10"}, {"value": "100"}, {"value": "34"},
	}
	inferred = inferSettings(csv, "heatmap")
	assert.Equal(t, float64(100), inferred.AbundanceCap, "Inferred heatmap abundance cap does not match expected")
	assert.Equal(t, "blueBlack", inferred.FillColor, "Inferred heatmap fill color does not match expected")
	assert.Equal(t, float64(0), inferred.MinAbundance, "Inferred heatmap min abundance does not match expected")

	// Test dotplot with negative values
	csv = []map[string]string{
		{"value": "-5.5"}, {"value": "10"}, {"value": "34.3"}, {"value": "-3"},
	}
	inferred = inferSettings(csv, "heatmap")
	assert.Equal(t, float64(35), inferred.AbundanceCap, "Inferred dotplot abundance cap does not match expected")
	assert.Equal(t, "redBlue", inferred.FillColor, "Inferred dotplot fill color does not match expected")
	assert.Equal(t, float64(-6), inferred.MinAbundance, "Inferred dotplot min abundance does not match expected")
}

func TestParseParams(t *testing.T) {
	// Test JSON params with all present
	csv := []map[string]string{
		{"values": "5", "params": `{
				"type": "dotplot",
				"xAxis": "Bait",
				"yAxis": "Prey",
				"filterType": 0,
				"primary": 0.01,
				"secondary": 0.05,
				"score": "BFDR",
				"abundance": "AvgSpec",
				"invert": 1
			}`,
		},
		{"values": "25.5"},
		{"values": "100"},
	}
	imageType, parameters := parseParams(csv)
	assert.Equal(t, "dotplot", imageType, "Processed JSON image type does not match expected")
	expectedParameters := typedef.Parameters{
		AbundanceCap:    50,
		Abundance:       "AvgSpec",
		Condition:       "Bait",
		EdgeColor:       "blueBlack",
		FillColor:       "blueBlack",
		InvertColor:     true,
		MinAbundance:    0,
		PrimaryFilter:   0.01,
		Readout:         "Prey",
		Score:           "BFDR",
		ScoreType:       "lte",
		SecondaryFilter: 0.05,
	}
	assert.Equal(
		t,
		expectedParameters,
		parameters,
		"Processed JSON parameters do not match expected",
	)

	// Test non json params with all present
	csv = []map[string]string{
		{"params": "dotplot"},
		{"params": "1"},
		{"params": "0.01"},
		{"params": "0.05"},
		{"params": "BFDR"},
		{"params": "AvgSpec"},
		{"params": "0"},
	}
	imageType, parameters = parseParams(csv)
	assert.Equal(t, "dotplot", imageType, "Processed text image type does not match expected")
	expectedParameters = typedef.Parameters{
		AbundanceCap:    50,
		Abundance:       "AvgSpec",
		EdgeColor:       "blueBlack",
		FillColor:       "blueBlack",
		InvertColor:     false,
		MinAbundance:    0,
		PrimaryFilter:   0.01,
		Score:           "BFDR",
		ScoreType:       "gte",
		SecondaryFilter: 0.05,
	}
	assert.Equal(
		t,
		expectedParameters,
		parameters,
		"Processed text parameters do not match expected",
	)
}
