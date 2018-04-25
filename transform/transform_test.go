package transform

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/types"
	"github.com/stretchr/testify/assert"
)

func TestPreys(t *testing.T) {
	data := []map[string]interface{}{
		{"bait": "bait1", "prey": "prey1", "abundance": "7", "control": "4", "preyLength": "3"},
		{"bait": "bait1", "prey": "prey2", "abundance": "22", "control": "2", "preyLength": "5"},
		{"bait": "bait2", "prey": "prey1", "abundance": "16", "control": "4", "preyLength": "3"},
		{"bait": "bait2", "prey": "prey2", "abundance": "7", "control": "2", "preyLength": "5"},
	}

	// TEST1: typical dataset performing all transformations
	want := []map[string]interface{}{
		{"bait": "bait1", "prey": "prey1", "abundance": "2", "control": "4", "preyLength": "3"},
		{"bait": "bait1", "prey": "prey2", "abundance": "4", "control": "2", "preyLength": "5"},
		{"bait": "bait2", "prey": "prey1", "abundance": "4", "control": "4", "preyLength": "3"},
		{"bait": "bait2", "prey": "prey2", "abundance": "2", "control": "2", "preyLength": "5"},
	}
	dataset := types.Dataset{
		Data: data,
		Params: types.Parameters{
			Control:           "controlColumn",
			LogBase:           "2",
			Normalization:     "total",
			NormalizationPrey: "",
			PreyLength:        "preyLengthColumn",
		},
	}
	assert.Equal(t, want, Preys(dataset), "Dataset not transformed correctly")
}
