package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaitprey(t *testing.T) {
	data := []map[string]string{
		{"bait": "bait1", "prey": "prey1"},
		{"bait": "bait2", "prey": "prey2"},
	}

	// TEST1: filter by a single bait and prey
	baits := []string{"bait2"}
  preys := []string{"prey2"}
	want := []map[string]string{
		{"bait": "bait2", "prey": "prey2"},
	}
	assert.Equal(
		t,
		want,
		Baitprey(data, baits, preys),
		"Single bait and prey filter is not returning correct slice map",
	)

  // TEST2: return empty slice map when not matches to bait and prey lists
	baits = []string{"bait2"}
  preys = []string{"prey1"}
	want = []map[string]string{}
	assert.Equal(
		t,
		want,
		Baitprey(data, baits, preys),
		"Returned value should have length zero",
	)

  // TEST3: filter by a multiple baits and preys
	baits = []string{"bait1", "bait2"}
  preys = []string{"prey1", "prey2"}
  want = data
	assert.Equal(
		t,
		want,
		Baitprey(data, baits, preys),
		"Multiple bait and prey filter is not returning correct slice map",
	)
}
