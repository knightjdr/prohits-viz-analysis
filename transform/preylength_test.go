package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreyLength(t *testing.T) {
	data := []map[string]interface{}{
		{"prey": "prey1", "abundance": "10", "preyLength": "2"},
		{"prey": "prey2", "abundance": "1", "preyLength": "5"},
		{"prey": "prey1", "abundance": "10|5", "preyLength": "2"},
		{"prey": "prey3", "abundance": "10|5|2.5", "preyLength": "10"},
	}
	want := []map[string]interface{}{
		{"prey": "prey1", "abundance": "25", "preyLength": "2"},
		{"prey": "prey2", "abundance": "1", "preyLength": "5"},
		{"prey": "prey1", "abundance": "25|12.5", "preyLength": "2"},
		{"prey": "prey3", "abundance": "5|2.5|1.25", "preyLength": "10"},
	}

	// TEST1: different abundance formats with prey length normalization.
	transformed := PreyLength(data, "preyLengthColumn")
	assert.Equal(t, want, transformed, "Prey length normalization is not correct")

	// TEST2: when prey length normalization is not requested, return original data.
	transformed = PreyLength(data, "")
	assert.Equal(t, data, data, "No prey length normalization should return input data")
}
