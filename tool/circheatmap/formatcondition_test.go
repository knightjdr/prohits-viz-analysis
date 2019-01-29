package circheatmap

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestFormatReadout(t *testing.T) {
	// TEST: when readouts have known key
	testFunc := formatReadout(true)
	readoutData := []map[string]float64{
		{"known": 1},
		{"known": 0},
	}
	want := []map[string]interface{}{
		{"name": "readoutX", "known": true},
		{"name": "readoutX", "known": false},
	}
	for index, readoutDatum := range readoutData {
		assert.Equal(t, want[index], testFunc("readoutX", readoutDatum), "Readout should be formatted correctly")
	}

	// TEST: when readouts should not have known key
	testFunc = formatReadout(false)
	want = []map[string]interface{}{
		{"name": "readoutX"},
		{"name": "readoutX"},
	}
	for index, readoutDatum := range readoutData {
		assert.Equal(t, want[index], testFunc("readoutX", readoutDatum), "Readout should be formatted correctly")
	}
}

func TestReadoutKeys(t *testing.T) {
	// TEST: should return keys from map
	hash := map[string]map[string]float64{
		"readoutX": {"abundance": 50, "other": 25},
		"readoutY": {"abundance": 25, "other": 4},
		"readoutZ": {"abundance": 5, "other": 17},
	}
	want := []string{"readoutX", "readoutY", "readoutZ"}
	assert.Equal(t, want, readoutKeys(hash), "Readout keys should be returned from hash")
}

func TestFormatCondition(t *testing.T) {
	data := map[string]map[string]float64{
		"readout1": map[string]float64{
			"abundance": 50,
			"HEK 293":   10.4250345,
			"HeLa":      6.7,
			"known":     1,
		},
		"readout3": map[string]float64{
			"abundance": 10,
			"HEK 293":   5,
			"HeLa":      8.1,
			"known":     0,
		},
	}
	metricOrder := []string{"abundance", "HEK 293", "HeLa"}
	readoutMetrics := map[string]string{
		"abundance": "AvgSpec",
		"HEK 293":   "RNA expression HEK 293",
		"HeLa":      "RNA expression HeLa",
	}

	// TEST

	want := typedef.CircHeatmapPlot{
		Name: "conditionA",
		Readouts: []map[string]interface{}{
			{"known": true, "name": "readout1"},
			{"known": false, "name": "readout3"},
		},
		Segments: []typedef.CircHeatmapSegments{
			{Name: "AvgSpec", Values: []float64{50, 10}},
			{Name: "RNA expression HEK 293", Values: []float64{10.43, 5}},
			{Name: "RNA expression HeLa", Values: []float64{6.7, 8.1}},
		},
	}
	result := formatCondition("conditionA", data, true, metricOrder, readoutMetrics)
	assert.Equal(t, want, result, "Plot not formatted correctly")
}
