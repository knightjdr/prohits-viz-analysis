package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConditions(t *testing.T) {
	data := []map[string]string{
		{"condition": "condition1", "readout": "readout1"},
		{"condition": "condition2", "readout": "readout2"},
	}

	// TEST1: filter by a single condition.
	conditions := []string{"condition2"}
	want := []map[string]string{
		{"condition": "condition2", "readout": "readout2"},
	}
	assert.Equal(
		t,
		want,
		Conditions(data, conditions),
		"Single condition filter is not correct",
	)

	// TEST1: filter by multiple conditions.
	conditions = []string{"condition1", "condition2"}
	want = data
	assert.Equal(
		t,
		want,
		Conditions(data, conditions),
		"Single condition filter is not correct",
	)
}
