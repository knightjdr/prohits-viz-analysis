package dotplot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConditionReadoutMatrix(t *testing.T) {
	dataset := []map[string]interface{}{
		{"condition": "acondition", "readout": "xreadout", "abundance": "5", "score": 0.01},
		{"condition": "acondition", "readout": "zreadout", "abundance": "10", "score": 0.02},
		{"condition": "acondition", "readout": "yreadout", "abundance": "23", "score": float64(0)},
		{"condition": "ccondition", "readout": "zreadout", "abundance": "7", "score": 0.01},
		{"condition": "ccondition", "readout": "xreadout", "abundance": "14.3", "score": 0.08},
		{"condition": "bcondition", "readout": "yreadout", "abundance": "17.8", "score": 0.01},
		{"condition": "bcondition", "readout": "xreadout", "abundance": "2", "score": 0.01},
	}

	// TEST1: dataset converted to matrix with smaller scores better.
	wantConditionList := []string{"acondition", "bcondition", "ccondition"}
	wantAbundance := [][]float64{
		{5, 2, 14.3},
		{23, 17.8, 0},
		{10, 0, 7},
	}
	wantReadoutList := []string{"xreadout", "yreadout", "zreadout"}
	wantScore := [][]float64{
		{0.01, 0.01, 0.08},
		{0, 0.01, 0.08},
		{0.02, 0.08, 0.01},
	}
	data := ConditionReadoutMatrix(dataset, "lte")
	assert.Equal(t, wantAbundance, data.Abundance, "Data not converted to condition readout abundance matrix")
	assert.Equal(t, wantConditionList, data.Conditions, "Condition list not correct")
	assert.Equal(t, wantReadoutList, data.Readouts, "Readout list not correct")
	assert.Equal(t, wantScore, data.Score, "Data not converted to condition readout score matrix")

	// TEST2: dataset converted to matrix with larger scores better.
	wantScore = [][]float64{
		{0.01, 0.01, 0.08},
		{0, 0.01, 0},
		{0.02, 0, 0.01},
	}
	data = ConditionReadoutMatrix(dataset, "gte")
	assert.Equal(t, wantScore, data.Score, "Data not converted to condition readout score matrix")
}
