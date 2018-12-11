package errorcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadoutLengthInt(t *testing.T) {
	// TEST1: valid readout length column returns no error.
	data := []map[string]string{
		{"readoutLength": "10"},
	}
	err := ReadoutLengthInt(data, "readoutLengthColumn")
	assert.Nil(t, err, "Valid readout length column should not return an error")

	// TEST2: readout length column no requested returns no error.
	data = []map[string]string{
		{"readoutLength": "10"},
	}
	err = ReadoutLengthInt(data, "")
	assert.Nil(t, err, "Missing readout length column should not return an error")

	// TEST3: invalid readout length column returns error.
	tests := [3]string{"a", "", "0.1"}
	for _, value := range tests {
		data = []map[string]string{
			{"readoutLength": value},
		}
		err = ReadoutLengthInt(data, "readoutLengthColumn")
		assert.NotNil(t, err, "Invalid readout length column should return an error")
	}
}
