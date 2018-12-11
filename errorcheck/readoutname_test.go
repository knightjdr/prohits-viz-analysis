package errorcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadoutName(t *testing.T) {
	// TEST1: valid readout names returns no error.
	data := []map[string]string{
		{"readout": "b"},
		{"readout": "d"},
	}
	err := ReadoutName(data)
	assert.Nil(t, err, "Valid data should not return an error")

	// TEST2: missing readout name returns an error.
	data = []map[string]string{
		{"readout": "b"},
		{"readout": ""},
	}
	err = ReadoutName(data)
	assert.NotNil(t, err, "Missing readout name should return an error")
}
