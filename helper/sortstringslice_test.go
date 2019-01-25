package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortStringSlice(t *testing.T) {
	// TEST: sort ascending
	expected := []string{"a", "B", "c"}
	result := SortStringSlice([]string{"B", "c", "a"}, "asc")
	assert.Equal(t, expected, result, "Slice not sorted in ascending order")

	// TEST: sort ascending
	expected = []string{"c", "B", "a"}
	result = SortStringSlice([]string{"B", "c", "a"}, "des")
	assert.Equal(t, expected, result, "Slice not sorted in descending order")
}
