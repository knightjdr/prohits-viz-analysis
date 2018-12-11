package errorcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCondition(t *testing.T) {
	// TEST: required condition number does not return an error.
	data := []map[string]string{
		{"condition": "a"},
		{"condition": "c"},
	}
	err := MinCondition(data, "dotplot")
	assert.Nil(t, err, "Required condition number should not return an error")

	// TEST: invalid data returns an error.
	data = []map[string]string{
		{"condition": "a"},
	}
	err = MinCondition(data, "dotplot")
	assert.NotNil(t, err, "Less than the required condition number should return an error")
}
