package errorcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreyName(t *testing.T) {
	// TEST1: valid prey names returns no error
	data := []map[string]interface{}{
		{"bait": "a", "prey": "b", "score": 0.1},
		{"bait": "c", "prey": "d", "score": 0.2},
	}
	err := PreyName(data)
	assert.Nil(t, err, "Valid data should not return an error")

	// TEST2: missing prey name returns an error
	data = []map[string]interface{}{
		{"bait": "a", "prey": "b", "score": 0.1},
		{"bait": "c", "prey": "", "score": 0.2},
	}
	err = PreyName(data)
	assert.NotNil(t, err, "Missing prey name should return an error")
}
