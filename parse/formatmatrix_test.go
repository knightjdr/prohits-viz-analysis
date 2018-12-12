package parse

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestFormatMatrix(t *testing.T) {
	rows := []Row{
		{
			Data: []Column{
				{Ratio: 0.2, Score: 0.05, Value: 5},
				{Ratio: 0.4, Score: 0.01, Value: 10},
				{Ratio: 1, Score: 0.01, Value: 40},
			},
			Name: "a",
		},
		{
			Data: []Column{
				{Ratio: 0.1, Score: 0.08, Value: 8},
				{Ratio: 1, Score: 0.01, Value: 60},
				{Ratio: 0.2, Score: 0.03, Value: 15},
			},
			Name: "b",
		},
		{
			Data: []Column{
				{Ratio: 0.6, Score: 0.01, Value: 17},
				{Ratio: 0.2, Score: 0.05, Value: 5},
				{Ratio: 1, Score: 0.01, Value: 30},
			},
			Name: "c",
		},
	}

	// TEST1: returns matrices.
	data := Data{
		Rows: rows,
	}
	expected := typedef.Matrices{
		Abundance: [][]float64{
			{5, 10, 40},
			{8, 60, 15},
			{17, 5, 30},
		},
		Ratio: [][]float64{
			{0.2, 0.4, 1},
			{0.1, 1, 0.2},
			{0.6, 0.2, 1},
		},
		Score: [][]float64{
			{0.05, 0.01, 0.01},
			{0.08, 0.01, 0.03},
			{0.01, 0.05, 0.01},
		},
	}
	result := FormatMatrix(&data)
	assert.ElementsMatch(t, expected.Abundance, result.Abundance)
	assert.ElementsMatch(t, expected.Ratio, result.Ratio)
	assert.ElementsMatch(t, expected.Score, result.Score)
}
