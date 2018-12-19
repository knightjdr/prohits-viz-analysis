package nocluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecList(t *testing.T) {
	fileData := []map[string]string{
		{"condition": "a"},
		{"condition": "b"},
		{"condition": "a"},
		{"condition": "c"},
	}

	// TEST: should return a list of names
	inputList := []string{"b", "a", "d"}
	expected := []string{"b", "a"}
	assert.Equal(t, expected, checkList(fileData, "condition", inputList), "Incorrect name list returned")
}
