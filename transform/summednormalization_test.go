package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummedNormalization(t *testing.T) {
	data := []map[string]interface{}{
		{"bait": "bait1", "prey": "prey1", "abundance": "10"},
		{"bait": "bait1", "prey": "prey2", "abundance": "5"},
		{"bait": "bait2", "prey": "prey1", "abundance": "1"},
		{"bait": "bait2", "prey": "prey2", "abundance": "2"},
		{"bait": "bait3", "prey": "prey1", "abundance": "10|5"},
		{"bait": "bait3", "prey": "prey2", "abundance": "4|6"},
		{"bait": "bait4", "prey": "prey1", "abundance": "15|8|7"},
		{"bait": "bait5", "prey": "prey2", "abundance": "10"},
	}
	want := []map[string]interface{}{
		{"bait": "bait1", "prey": "prey1", "abundance": "10"},
		{"bait": "bait1", "prey": "prey2", "abundance": "5"},
		{"bait": "bait2", "prey": "prey1", "abundance": "5"},
		{"bait": "bait2", "prey": "prey2", "abundance": "10"},
		{"bait": "bait3", "prey": "prey1", "abundance": "6|3"},
		{"bait": "bait3", "prey": "prey2", "abundance": "2.4|3.6"},
		{"bait": "bait4", "prey": "prey1", "abundance": "7.5|4|3.5"},
		{"bait": "bait5", "prey": "prey2", "abundance": "15"},
	}

	// TEST1: summed normalization
	assert.Equal(
		t,
		want,
		SummedNormalization(data),
		"Preys are not being normalized correctly by total sum",
	)
}
