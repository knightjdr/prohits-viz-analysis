package errorcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinData(t *testing.T) {
	// TEST: valid data returns no error.
	data := []map[string]string{
		{"condition": "a", "readout": "b", "score": "0.1"},
		{"condition": "c", "readout": "d", "score": "0.2"},
	}
	err := MinData(data)
	assert.Nil(t, err, "Valid data should not return an error")

	// TEST: invalid data returns an error.
	data = []map[string]string{}
	err = MinData(data)
	assert.NotNil(t, err, "Invalid data should return an error")
}
