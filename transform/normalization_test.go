package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalization(t *testing.T) {
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
	// TEST1: no normalization.
	assert.Equal(
		t,
		data,
		Normalization(data, "none", ""),
		"Normalization not required should return original data",
	)

	// TEST2: readout normalization.
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
	assert.Equal(
		t,
		want,
		Normalization(data, "readout", "readout1"),
		"Readout normalization data transformation is not correct",
	)

	// TEST3: total sum normalization.
	data = []map[string]interface{}{
		{"condition": "condition1", "readout": "readout1", "abundance": "10"},
		{"condition": "condition1", "readout": "readout2", "abundance": "5"},
		{"condition": "condition2", "readout": "readout1", "abundance": "1"},
		{"condition": "condition2", "readout": "readout2", "abundance": "2"},
		{"condition": "condition3", "readout": "readout1", "abundance": "10|5"},
		{"condition": "condition3", "readout": "readout2", "abundance": "4|6"},
		{"condition": "condition4", "readout": "readout1", "abundance": "15|8|7"},
		{"condition": "condition5", "readout": "readout2", "abundance": "10"},
	}
	want = []map[string]interface{}{
		{"condition": "condition1", "readout": "readout1", "abundance": "10"},
		{"condition": "condition1", "readout": "readout2", "abundance": "5"},
		{"condition": "condition2", "readout": "readout1", "abundance": "5"},
		{"condition": "condition2", "readout": "readout2", "abundance": "10"},
		{"condition": "condition3", "readout": "readout1", "abundance": "6|3"},
		{"condition": "condition3", "readout": "readout2", "abundance": "2.4|3.6"},
		{"condition": "condition4", "readout": "readout1", "abundance": "7.5|4|3.5"},
		{"condition": "condition5", "readout": "readout2", "abundance": "15"},
	}
	assert.Equal(
		t,
		want,
		Normalization(data, "total", ""),
		"Total sum normalization data transformation is not correct",
	)
}
