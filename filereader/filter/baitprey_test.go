package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConditionReadout(t *testing.T) {
	data := []map[string]string{
		{"condition": "condition1", "readout": "readout1"},
		{"condition": "condition2", "readout": "readout2"},
	}

	// TEST1: filter by a single condition and readout.
	conditions := []string{"condition2"}
	readouts := []string{"readout2"}
	want := []map[string]string{
		{"condition": "condition2", "readout": "readout2"},
	}
	assert.Equal(
		t,
		want,
		ConditionReadout(data, conditions, readouts),
		"Single condition and readout filter is not returning correct slice map",
	)

	// TEST2: return empty slice map when not matches to condition and readout lists.
	conditions = []string{"condition2"}
	readouts = []string{"readout1"}
	want = []map[string]string{}
	assert.Equal(
		t,
		want,
		ConditionReadout(data, conditions, readouts),
		"Returned value should have length zero",
	)

	// TEST3: filter by a multiple conditions and readouts.
	conditions = []string{"condition1", "condition2"}
	readouts = []string{"readout1", "readout2"}
	want = data
	assert.Equal(
		t,
		want,
		ConditionReadout(data, conditions, readouts),
		"Multiple condition and readout filter is not returning correct slice map",
	)
}
