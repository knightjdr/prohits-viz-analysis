package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDummy(t *testing.T) {
	// TEST1: returns slices or dummy names.
	wantColumns := []string{"column", "column", "column"}
	wantRows := []string{"row", "row", "row"}
	actualColumns, actualRows := Dummy(3, 3)
	assert.ElementsMatch(t, wantColumns, actualColumns)
	assert.ElementsMatch(t, wantRows, actualRows)
}
