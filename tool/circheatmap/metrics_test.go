package circheatmap

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestMetrics(t *testing.T) {
	// TEST: all possible metrics
	parameters := typedef.Parameters{
		Abundance:      "AvgSpec",
		OtherAbundance: []string{"FC", "Other"},
	}
	expected := map[string]string{
		"abundance": "AvgSpec",
		"FC":        "FC",
		"Other":     "Other",
	}
	result := metrics(parameters)
	assert.Equal(t, result, expected, "Readout metrics are not parsed correctly")

	// TEST: minumum possible metrics
	parameters = typedef.Parameters{
		Abundance: "AvgSpec",
	}
	expected = map[string]string{
		"abundance": "AvgSpec",
	}
	result = metrics(parameters)
	assert.Equal(t, result, expected, "Minimal readout metrics are not parsed correctly")
}
