package errorcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestControlColumn(t *testing.T) {
	// TEST1: valid control column returns no error.
	validTests := [2]interface{}{"1", "1|2.3|0.1"}
	for _, value := range validTests {
		data := []map[string]interface{}{
			{"control": value},
		}
		err := ControlColumn(data, "controlColumn")
		assert.Nil(t, err, "Valid control column should not return an error")
	}

	// TEST2: missing control column returns no error.
	data := []map[string]interface{}{
		{"control": 10},
	}
	err := ControlColumn(data, "")
	assert.Nil(t, err, "Missing control column should not return an error")

	// TEST3: invalid control column returns error.
	invalidTests := [3]interface{}{"a", "", "1|a|b|0.1"}
	for _, value := range invalidTests {
		data = []map[string]interface{}{
			{"control": value},
		}
		err = ControlColumn(data, "controlColumn")
		assert.NotNil(t, err, "Invalid control column should return an error")
	}
}
