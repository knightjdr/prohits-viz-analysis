package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummedNormalization(t *testing.T) {
	data := []map[string]string{
		{"condition": "condition1", "readout": "readout1", "abundance": "10"},
		{"condition": "condition1", "readout": "readout2", "abundance": "5"},
		{"condition": "condition2", "readout": "readout1", "abundance": "1"},
		{"condition": "condition2", "readout": "readout2", "abundance": "2"},
		{"condition": "condition3", "readout": "readout1", "abundance": "10|5"},
		{"condition": "condition3", "readout": "readout2", "abundance": "4|6"},
		{"condition": "condition4", "readout": "readout1", "abundance": "15|8|7"},
		{"condition": "condition5", "readout": "readout2", "abundance": "10"},
		{"condition": "condition6", "readout": "readout2", "abundance": "0"},
		{"condition": "condition7", "readout": "readout2", "abundance": "30"},
	}
	want := []map[string]string{
		{"condition": "condition1", "readout": "readout1", "abundance": "10"},
		{"condition": "condition1", "readout": "readout2", "abundance": "5"},
		{"condition": "condition2", "readout": "readout1", "abundance": "5"},
		{"condition": "condition2", "readout": "readout2", "abundance": "10"},
		{"condition": "condition3", "readout": "readout1", "abundance": "6|3"},
		{"condition": "condition3", "readout": "readout2", "abundance": "2.4|3.6"},
		{"condition": "condition4", "readout": "readout1", "abundance": "7.5|4|3.5"},
		{"condition": "condition5", "readout": "readout2", "abundance": "15"},
		{"condition": "condition6", "readout": "readout2", "abundance": "0"},
		{"condition": "condition7", "readout": "readout2", "abundance": "15"},
	}

	// TEST1: summed normalization
	assert.Equal(
		t,
		want,
		SummedNormalization(data),
		"Readouts are not being normalized correctly by total sum",
	)
}
