package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadouts(t *testing.T) {
	data := []map[string]string{
		{"condition": "condition1", "readout": "readout1"},
		{"condition": "condition1", "readout": "readout2"},
	}

	// TEST1: filter by a single readout.
	readouts := []string{"readout2"}
	want := []map[string]string{
		{"condition": "condition1", "readout": "readout2"},
	}
	assert.Equal(
		t,
		want,
		Readouts(data, readouts),
		"Single readout filter is not correct",
	)

	// TEST1: filter by multiple readouts.
	readouts = []string{"readout1", "readout2"}
	want = data
	assert.Equal(
		t,
		want,
		Readouts(data, readouts),
		"Single readout filter is not correct",
	)
}
