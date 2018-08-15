package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatData(t *testing.T) {
	rows := [][]Row{
		{
			{Ratio: 0.2, Score: 0.05, Value: 5},
			{Ratio: 0.4, Score: 0.01, Value: 10},
			{Ratio: 1, Score: 0.01, Value: 40},
		},
		{
			{Ratio: 0.1, Score: 0.08, Value: 8},
			{Ratio: 1, Score: 0.01, Value: 60},
			{Ratio: 0.2, Score: 0.03, Value: 15},
		},
		{
			{Ratio: 0.6, Score: 0.01, Value: 17},
			{Ratio: 0.2, Score: 0.05, Value: 5},
			{Ratio: 1, Score: 0.01, Value: 30},
		},
	}

	// TEST1: returns matrices.
	data := Data{
		Rows: rows,
	}
	wantAbundance := [][]float64{
		{5, 10, 40},
		{8, 60, 15},
		{17, 5, 30},
	}
	wantRatio := [][]float64{
		{0.2, 0.4, 1},
		{0.1, 1, 0.2},
		{0.6, 0.2, 1},
	}
	wantScore := [][]float64{
		{0.05, 0.01, 0.01},
		{0.08, 0.01, 0.03},
		{0.01, 0.05, 0.01},
	}
	actualAbundance, actualRatio, actualScore := FormatData(&data)
	assert.ElementsMatch(t, wantAbundance, actualAbundance)
	assert.ElementsMatch(t, wantRatio, actualRatio)
	assert.ElementsMatch(t, wantScore, actualScore)
}
