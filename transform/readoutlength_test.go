package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadoutLength(t *testing.T) {
	data := []map[string]interface{}{
		{"readout": "readout1", "abundance": "10", "readoutLength": "2"},
		{"readout": "readout2", "abundance": "1", "readoutLength": "5"},
		{"readout": "readout1", "abundance": "10|5", "readoutLength": "2"},
		{"readout": "readout3", "abundance": "10|5|2.5", "readoutLength": "10"},
	}
	want := []map[string]interface{}{
		{"readout": "readout1", "abundance": "25", "readoutLength": "2"},
		{"readout": "readout2", "abundance": "1", "readoutLength": "5"},
		{"readout": "readout1", "abundance": "25|12.5", "readoutLength": "2"},
		{"readout": "readout3", "abundance": "5|2.5|1.25", "readoutLength": "10"},
	}

	// TEST1: different abundance formats with readout length normalization.
	transformed := ReadoutLength(data, "readoutLengthColumn")
	assert.Equal(t, want, transformed, "Readout length normalization is not correct")

	// TEST2: when readout length normalization is not requested, return original data.
	transformed = ReadoutLength(data, "")
	assert.Equal(t, data, transformed, "No readout length normalization should return input data")
}
