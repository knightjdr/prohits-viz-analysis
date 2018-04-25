package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalization(t *testing.T) {
	data := []map[string]interface{}{
		{"bait": "bait1", "prey": "prey1", "abundance": "10"},
		{"bait": "bait1", "prey": "prey2", "abundance": "5"},
		{"bait": "bait2", "prey": "prey1", "abundance": "1"},
		{"bait": "bait2", "prey": "prey2", "abundance": "2"},
		{"bait": "bait3", "prey": "prey1", "abundance": "10|5"},
		{"bait": "bait3", "prey": "prey2", "abundance": "4|6"},
		{"bait": "bait4", "prey": "prey1", "abundance": "12|8|5"},
		{"bait": "bait4", "prey": "prey2", "abundance": "8|2|15.5"},
		{"bait": "bait5", "prey": "prey2", "abundance": "10"},
	}
	// TEST1: no normalization.
	assert.Equal(
		t,
		data,
		Normalization(data, "none", ""),
		"Normalization not required should return original data",
	)

	// TEST2: prey normalization.
	want := []map[string]interface{}{
		{"bait": "bait1", "prey": "prey1", "abundance": "12.5"},
		{"bait": "bait1", "prey": "prey2", "abundance": "6.25"},
		{"bait": "bait2", "prey": "prey1", "abundance": "12.5"},
		{"bait": "bait2", "prey": "prey2", "abundance": "25"},
		{"bait": "bait3", "prey": "prey1", "abundance": "8.33|4.17"},
		{"bait": "bait3", "prey": "prey2", "abundance": "3.33|5"},
		{"bait": "bait4", "prey": "prey1", "abundance": "6|4|2.5"},
		{"bait": "bait4", "prey": "prey2", "abundance": "4|1|7.75"},
		{"bait": "bait5", "prey": "prey2", "abundance": "10"},
	}
	assert.Equal(
		t,
		want,
		Normalization(data, "prey", "prey1"),
		"Prey normalization data transformation is not correct",
	)

	// TEST3: total sum normalization.
	data = []map[string]interface{}{
		{"bait": "bait1", "prey": "prey1", "abundance": "10"},
		{"bait": "bait1", "prey": "prey2", "abundance": "5"},
		{"bait": "bait2", "prey": "prey1", "abundance": "1"},
		{"bait": "bait2", "prey": "prey2", "abundance": "2"},
		{"bait": "bait3", "prey": "prey1", "abundance": "10|5"},
		{"bait": "bait3", "prey": "prey2", "abundance": "4|6"},
		{"bait": "bait4", "prey": "prey1", "abundance": "15|8|7"},
		{"bait": "bait5", "prey": "prey2", "abundance": "10"},
	}
	want = []map[string]interface{}{
		{"bait": "bait1", "prey": "prey1", "abundance": "10"},
		{"bait": "bait1", "prey": "prey2", "abundance": "5"},
		{"bait": "bait2", "prey": "prey1", "abundance": "5"},
		{"bait": "bait2", "prey": "prey2", "abundance": "10"},
		{"bait": "bait3", "prey": "prey1", "abundance": "6|3"},
		{"bait": "bait3", "prey": "prey2", "abundance": "2.4|3.6"},
		{"bait": "bait4", "prey": "prey1", "abundance": "7.5|4|3.5"},
		{"bait": "bait5", "prey": "prey2", "abundance": "15"},
	}
	assert.Equal(
		t,
		want,
		Normalization(data, "total", ""),
		"Total sum normalization data transformation is not correct",
	)
}
