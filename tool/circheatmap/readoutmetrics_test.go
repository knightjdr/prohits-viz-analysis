package circheatmap

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestReadoutMetrics(t *testing.T) {
	// TEST: all possible metrics
	parameters := typedef.Parameters{
		Abundance:      "AvgSpec",
		OtherAbundance: []string{"FC", "Other"},
	}
	want := map[string]string{
		"abundance": "AvgSpec",
		"FC":        "FC",
		"Other":     "Other",
	}
	result := readoutMetrics(parameters)
	assert.Equal(t, want, result, "Readout metrics are not parsed correctly")

	// TEST: minumum possible metrics
	parameters = typedef.Parameters{
		Abundance: "AvgSpec",
	}
	want = map[string]string{
		"abundance": "AvgSpec",
	}
	result = readoutMetrics(parameters)
	assert.Equal(t, want, result, "Minimal readout metrics are not parsed correctly")
}
