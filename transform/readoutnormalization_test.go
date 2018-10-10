package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadoutNormalization(t *testing.T) {
	data := []map[string]interface{}{
		{"condition": "condition1", "readout": "readout1", "abundance": "10"},
		{"condition": "condition1", "readout": "readout2", "abundance": "5"},
		{"condition": "condition2", "readout": "readout1", "abundance": "1"},
		{"condition": "condition2", "readout": "readout2", "abundance": "2"},
		{"condition": "condition3", "readout": "readout1", "abundance": "10|5"},
		{"condition": "condition3", "readout": "readout2", "abundance": "4|6"},
		{"condition": "condition4", "readout": "readout1", "abundance": "12|8|5"},
		{"condition": "condition4", "readout": "readout2", "abundance": "8|2|15.5"},
		{"condition": "condition5", "readout": "readout2", "abundance": "10"},
	}
	want := []map[string]interface{}{
		{"condition": "condition1", "readout": "readout1", "abundance": "12.5"},
		{"condition": "condition1", "readout": "readout2", "abundance": "6.25"},
		{"condition": "condition2", "readout": "readout1", "abundance": "12.5"},
		{"condition": "condition2", "readout": "readout2", "abundance": "25"},
		{"condition": "condition3", "readout": "readout1", "abundance": "8.33|4.17"},
		{"condition": "condition3", "readout": "readout2", "abundance": "3.33|5"},
		{"condition": "condition4", "readout": "readout1", "abundance": "6|4|2.5"},
		{"condition": "condition4", "readout": "readout2", "abundance": "4|1|7.75"},
		{"condition": "condition5", "readout": "readout2", "abundance": "10"},
	}

	// TEST1: readout length normalization
	assert.Equal(
		t,
		want,
		ReadoutNormalization(data, "readout1"),
		"Readouts are not being normalized correctly by a specific readout",
	)
}
