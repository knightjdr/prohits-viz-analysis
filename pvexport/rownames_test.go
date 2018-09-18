package main

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/parse"
	"github.com/stretchr/testify/assert"
)

func TestRowNames(t *testing.T) {
	// TEST1: returns parameter type from Data struct.
	rows := []parse.Row{
		{Data: []parse.Column{}, Name: "a"},
		{Data: []parse.Column{}, Name: "b"},
		{Data: []parse.Column{}, Name: "c"},
	}
	expected := []string{"a", "b", "c"}
	assert.EqualValues(t, expected, RowNames(rows))
}
