package errorcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreyLengthInt(t *testing.T) {
	// TEST1: valid prey length column returns no error
	data := []map[string]interface{}{
		{"bait": "a", "prey": "b", "preyLength": "10"},
	}
	err := PreyLengthInt(data, "preyLengthColumn")
	assert.Nil(t, err, "Valid prey length column should not return an error")

	// TEST2: missing prey length column returns no error
	data = []map[string]interface{}{
		{"bait": "a", "prey": "b", "preyLength": "10"},
	}
	err = PreyLengthInt(data, "")
	assert.Nil(t, err, "Missing prey length column should not return an error")

	// TEST3: invalid prey length column returns error
	tests := [3]interface{}{"a", "", "0.1"}
	for _, value := range tests {
		data = []map[string]interface{}{
			{"bait": "a", "prey": "b", "preyLength": value},
		}
		err = PreyLengthInt(data, "preyLengthColumn")
		assert.NotNil(t, err, "Invalid prey length column should return an error")
	}
}
