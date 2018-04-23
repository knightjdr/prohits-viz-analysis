package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreys(t *testing.T) {
	data := []map[string]string{
		{"bait": "bait1", "prey": "prey1"},
		{"bait": "bait1", "prey": "prey2"},
	}

	// TEST1: filter by a single prey
	preys := []string{"prey2"}
	want := []map[string]string{
		{"bait": "bait1", "prey": "prey2"},
	}
	assert.Equal(
		t,
		want,
		Preys(data, preys),
		"Single prey filter is not correct",
	)

	// TEST1: filter by multiple preys
	preys = []string{"prey1", "prey2"}
	want = data
	assert.Equal(
		t,
		want,
		Preys(data, preys),
		"Single prey filter is not correct",
	)
}
