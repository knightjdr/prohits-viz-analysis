package errorcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbundanceColumn(t *testing.T) {
	// TEST1: valid abundance column returns no error.
	validTests := [2]string{"1", "1|2.3|0.1"}
	for _, value := range validTests {
		data := []map[string]string{
			{"abundance": value},
		}
		err := AbundanceColumn(data)
		assert.Nil(t, err, "Valid abundance column should not return an error")
	}

	// TEST3: invalid abundance column returns error.
	invalidTests := [3]string{"a", "", "1|a|b|0.1"}
	for _, value := range invalidTests {
		data := []map[string]string{
			{"abundance": value},
		}
		err := AbundanceColumn(data)
		assert.NotNil(t, err, "Invalid abundance column should return an error")
	}
}
