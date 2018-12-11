package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreFunc(t *testing.T) {
	// TEST: larger scores are better
	scoreFunction := scoreFunc("gte")
	testValues := []float64{1.4, 5, 8, 20, 27}
	threshold := float64(15)
	expectedValues := []float64{1.4, 5, 8, 15, 15}
	for i := range testValues {
		assert.Equal(
			t,
			expectedValues[i],
			scoreFunction(testValues[i], threshold),
			"gte score function is not returning correct thresholded value",
		)
	}

	// TEST: smaller scores are better
	scoreFunction = scoreFunc("lte")
	testValues = []float64{1.4, 5, 8, 20, 27}
	threshold = float64(8.5)
	expectedValues = []float64{8.5, 8.5, 8.5, 20, 27}
	for i := range testValues {
		assert.Equal(
			t,
			expectedValues[i],
			scoreFunction(testValues[i], threshold),
			"lte score function is not returning correct thresholded value",
		)
	}
}

func TestSortLabels(t *testing.T) {
	labels := map[string]int{
		"B": 2,
		"C": 0,
		"A": 4,
		"E": 1,
		"D": 3,
	}

	// TEST: sort labels by index stored as value.
	expected := []string{"C", "E", "B", "D", "A"}
	assert.Equal(t, expected, sortLabels(labels, false), "Labels not sorted by index")

	// TEST: sort labels alphabetically.
	expected = []string{"A", "B", "C", "D", "E"}
	assert.Equal(t, expected, sortLabels(labels, true), "Labels not sorted alphabetically")
}

func TestConditionReadoutMatrix(t *testing.T) {
	dataset := []map[string]string{
		{"condition": "acondition", "readout": "xreadout", "abundance": "5", "score": "0.01"},
		{"condition": "acondition", "readout": "zreadout", "abundance": "10", "score": "0.02"},
		{"condition": "acondition", "readout": "yreadout", "abundance": "23", "score": "0"},
		{"condition": "ccondition", "readout": "zreadout", "abundance": "7", "score": "0.01"},
		{"condition": "ccondition", "readout": "xreadout", "abundance": "14.3", "score": "0.08"},
		{"condition": "bcondition", "readout": "yreadout", "abundance": "17.8", "score": "0.01"},
		{"condition": "bcondition", "readout": "xreadout", "abundance": "2", "score": "0.01"},
	}

	// TEST: dataset converted to matrix with smaller scores better.
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
	data := ConditionReadoutMatrix(dataset, "lte", true)
	assert.Equal(t, wantAbundance, data.Abundance, "Data not converted to condition readout abundance matrix")
	assert.Equal(t, wantConditionList, data.Conditions, "Condition list not correct")
	assert.Equal(t, wantReadoutList, data.Readouts, "Readout list not correct")
	assert.Equal(t, wantScore, data.Score, "Data not converted to condition readout score matrix")

	// TEST: dataset converted to matrix with larger scores better.
	wantScore = [][]float64{
		{0.01, 0.01, 0.08},
		{0, 0.01, 0},
		{0.02, 0, 0.01},
	}
	data = ConditionReadoutMatrix(dataset, "gte", true)
	assert.Equal(t, wantScore, data.Score, "Data not converted to condition readout score matrix")
}
