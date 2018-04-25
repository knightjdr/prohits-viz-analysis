package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogTransform(t *testing.T) {
	data := []map[string]interface{}{
		{"abundance": "2"},
		{"abundance": "1"},
		{"abundance": "8|4"},
		{"abundance": "16|1|4"},
	}

	// TEST1: Invalid or no log transformation required
	assert.Equal(
		t,
		data,
		LogTransform(data, "none"),
		"Log transformation not required should return original data",
	)

	// TEST2: log base 2.
	want := []map[string]interface{}{
		{"abundance": "1"},
		{"abundance": "0"},
		{"abundance": "3|2"},
		{"abundance": "4|0|2"},
	}
	assert.Equal(
		t,
		want,
		LogTransform(data, "2"),
		"Log base 2 data transformation is not correct",
	)

	// TEST3: log base 10.
	data = []map[string]interface{}{
		{"abundance": "10"},
		{"abundance": "1"},
		{"abundance": "100|0.01"},
		{"abundance": "3|10|1"},
	}
	want = []map[string]interface{}{
		{"abundance": "1"},
		{"abundance": "0"},
		{"abundance": "2|0"},
		{"abundance": "0.48|1|0"},
	}
	assert.Equal(
		t,
		want,
		LogTransform(data, "10"),
		"Log base 10 data transformation is not correct",
	)
}
