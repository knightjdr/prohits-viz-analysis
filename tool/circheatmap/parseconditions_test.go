package circheatmap

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestParseConditions(t *testing.T) {
	fileData := []map[string]string{
		{"abundance": "10.5", "condition": "a", "Other": "4", "readout": "readout1", "score": "0"},
		{"abundance": "170.5", "condition": "a", "Other": "7", "readout": "readout2", "score": "0"},
		{"abundance": "0.8", "condition": "b", "Other": "23.9", "readout": "readout3", "score": "0.01"},
		{"abundance": "14.2", "condition": "b", "Other": "74.7", "readout": "readout4", "score": "0.2"},
	}
	parameters := typedef.Parameters{
		PrimaryFilter: 0.01,
		ScoreType:     "lte",
	}
	readoutMetrics := map[string]string{
		"abundance": "abundance",
		"Other":     "Other",
	}

	// TEST
	conditionNames, readoutMap, conditionData := parseConditions(fileData, parameters, readoutMetrics)
	expectedMap := map[string]bool{
		"readout1": true,
		"readout2": true,
		"readout3": true,
		"readout4": true,
	}
	expectedNames := []string{"a", "b"}
	assert.Equal(t, expectedNames, conditionNames, "Condition names not parsed from data")
	assert.Equal(t, expectedMap, readoutMap, "Readout map not parsed from data")
	expectedReadoutMetrics := map[string]map[string]map[string]float64{
		"a": {
			"readout1": {
				"abundance": 10.5,
				"Other":     4,
			},
			"readout2": {
				"abundance": 170.5,
				"Other":     7,
			},
		},
		"b": {
			"readout3": {
				"abundance": 0.8,
				"Other":     23.9,
			},
		},
	}
	assert.Equal(t, expectedReadoutMetrics, conditionData, "Readout metrics not parsed from data")
}
