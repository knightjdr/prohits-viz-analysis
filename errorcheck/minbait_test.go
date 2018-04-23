package errorcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinBait(t *testing.T) {
	// TEST1: required bait number does not return an error
	data := []map[string]interface{}{
		{"bait": "a", "prey": "b", "score": 0.1},
		{"bait": "c", "prey": "d", "score": 0.2},
	}
	err := MinBait(data, "dotplot")
	assert.Nil(t, err, "Required bait number should not return an error")

	// TEST2: invalid data returns an error
	data = []map[string]interface{}{
		{"bait": "a", "prey": "b", "score": 0.1},
	}
	err = MinBait(data, "dotplot")
	assert.NotNil(t, err, "Less than the required bait number should return an error")
}
