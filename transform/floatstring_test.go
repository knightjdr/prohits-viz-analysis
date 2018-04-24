package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatToString(t *testing.T) {
	type test struct {
		float float64
		want  string
	}

	// TEST1: conversion.
	tests := [4]test{
		test{10.032453, "10.032453"},
		test{2.183, "2.183"},
		test{-156.789235, "-156.789235"},
		test{10, "10"},
	}
	for _, test := range tests {
		assert.Equal(
			t,
			test.want,
			FloatToString(test.float),
			"Floats are not being converted to strings correctly",
		)
	}
}
