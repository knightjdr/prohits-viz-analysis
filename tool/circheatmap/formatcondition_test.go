package circheatmap

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestFormatCondition(t *testing.T) {
	data := map[string]map[string]float64{
		"readout1": map[string]float64{
			"abundance": 50,
			"HEK 293":   10.4,
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
			{Name: "RNA expression HEK 293", Values: []float64{10.4, 5}},
			{Name: "RNA expression HeLa", Values: []float64{6.7, 8.1}},
		},
	}
	result := formatCondition("conditionA", data, true, metricOrder, readoutMetrics)
	assert.Equal(t, want, result, "Plot not formatted correctly")
}
