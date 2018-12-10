package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertInvertColor(t *testing.T) {
	// Test
	assert.True(
		t,
		convertInvertColor(1),
		"1 not converted to true",
	)

	// Test
	assert.False(
		t,
		convertInvertColor(0),
		"0 not converted to false",
	)
}

func TestConvertScoreType(t *testing.T) {
	// Test
	assert.Equal(
		t,
		"gte",
		convertScoreType(1),
		"1 not converted to gte",
	)

	// Test
	assert.Equal(
		t,
		"lte",
		convertScoreType(0),
		"0 not converted to lte",
	)
}

func TestParseParams(t *testing.T) {
	// Test JSON params with all present
	csv := []map[string]string{
		{"params": `{
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
	}
	params := parseParams(csv)
	expected := parameters{
		abundanceColumn: "AvgSpec",
		conditionColumn: "Bait",
		imageType:       "dotplot",
		invertColor:     true,
		primaryFilter:   0.01,
		readoutColumn:   "Prey",
		secondaryFilter: 0.05,
		scoreColumn:     "BFDR",
		scoreType:       "lte",
	}
	assert.Equal(
		t,
		expected,
		params,
		"Processed JSON params do not match expected",
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
	params = parseParams(csv)
	expected = parameters{
		abundanceColumn: "AvgSpec",
		imageType:       "dotplot",
		invertColor:     false,
		primaryFilter:   0.01,
		secondaryFilter: 0.05,
		scoreColumn:     "BFDR",
		scoreType:       "gte",
	}
	assert.Equal(
		t,
		expected,
		params,
		"Processed text params do not match expected",
	)
}
