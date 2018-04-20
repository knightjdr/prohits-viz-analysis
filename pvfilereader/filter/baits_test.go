package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaits(t *testing.T) {
	data := []map[string]string{
		{"bait": "bait1", "prey": "prey1"},
		{"bait": "bait2", "prey": "prey2"},
	}

	// TEST1: filter by a single bait
	baits := []string{"bait2"}
	want := []map[string]string{
		{"bait": "bait2", "prey": "prey2"},
	}
	assert.Equal(
		t,
		want,
		Baits(data, baits),
		"Single bait filter is not correct",
	)

	// TEST1: filter by multiple baits
	baits = []string{"bait1", "bait2"}
	want = data
	assert.Equal(
		t,
		want,
		Baits(data, baits),
		"Single bait filter is not correct",
	)
}
