package circheatmap

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestSegmentSettings(t *testing.T) {
	// TEST
	parameters := typedef.Parameters{
		AbundanceCap: 50,
		FillColor:    "blueBlack",
		MinAbundance: 0,
	}
	readoutMetrics := map[string]string{
		"abundance": "Abd",
		"FC":        "FC",
		"HeLa":      "HeLa expression",
		"HEK 293":   "HEK 293 expression",
	}
	wantOrder := []string{"abundance", "FC", "HEK 293", "HeLa"}
	wantSettings := []typedef.CircHeatmapSetttings{
		{AbundanceCap: 50, Color: "blueBlack", MinAbundance: 0, Name: "Abd"},
		{AbundanceCap: 50, Color: "blueBlack", MinAbundance: 0, Name: "FC"},
		{AbundanceCap: 50, Color: "blueBlack", MinAbundance: 0, Name: "HEK 293 expression"},
		{AbundanceCap: 50, Color: "blueBlack", MinAbundance: 0, Name: "HeLa expression"},
	}
	order, settings := segmentSettings(readoutMetrics, parameters)
	assert.Equal(t, wantOrder, order, "Segments not formatted correctly")
	assert.Equal(t, wantSettings, settings, "Segments not formatted correctly")
}
