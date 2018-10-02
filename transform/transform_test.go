package transform

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestReadouts(t *testing.T) {
	data := []map[string]interface{}{
		{"condition": "condition1", "readout": "readout1", "abundance": "7", "control": "4", "readoutLength": "3"},
		{"condition": "condition1", "readout": "readout2", "abundance": "22", "control": "2", "readoutLength": "5"},
		{"condition": "condition2", "readout": "readout1", "abundance": "16", "control": "4", "readoutLength": "3"},
		{"condition": "condition2", "readout": "readout2", "abundance": "7", "control": "2", "readoutLength": "5"},
	}

	// TEST1: typical dataset performing all transformations
	want := []map[string]interface{}{
		{"condition": "condition1", "readout": "readout1", "abundance": "2", "control": "4", "readoutLength": "3"},
		{"condition": "condition1", "readout": "readout2", "abundance": "4", "control": "2", "readoutLength": "5"},
		{"condition": "condition2", "readout": "readout1", "abundance": "4", "control": "4", "readoutLength": "3"},
		{"condition": "condition2", "readout": "readout2", "abundance": "2", "control": "2", "readoutLength": "5"},
	}
	dataset := typedef.Dataset{
		Data: data,
		Parameters: typedef.Parameters{
			Control:              "controlColumn",
			LogBase:              "2",
			Normalization:        "total",
			NormalizationReadout: "",
			ReadoutLength:        "readoutLengthColumn",
		},
	}
	assert.Equal(t, want, Readouts(dataset), "Dataset not transformed correctly")
}
