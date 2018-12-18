package interactive

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestRowData(t *testing.T) {
	// Starting with some hypothetical data for a dotplot, generate a data matrix
	// with formatted rows for use in the heatmap.
	matrices := new(typedef.Matrices)
	matrices.Abundance = [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matrices.Conditions = []string{"col1", "col2", "col3"}
	matrices.Ratio = [][]float64{
		{0.2, 0.5, 1},
		{0.7, 0.8, 1},
		{1, 0.2, 0.5},
	}
	matrices.Readouts = []string{"row1", "row2", "row3"}
	matrices.Score = [][]float64{
		{0.01, 0.05, 0.08},
		{1, 0.07, 0.5},
		{0.2, 0.7, 0.01},
	}

	// TEST: data conversion to row dotplot matrix.
	expected := []map[string]interface{}{
		{
			"name": "row1",
			"data": []map[string]float64{
				{"ratio": 0.2, "score": 0.01, "value": 1},
				{"ratio": 0.5, "score": 0.05, "value": 2},
				{"ratio": 1, "score": 0.08, "value": 3},
			},
		},
		{
			"name": "row2",
			"data": []map[string]float64{
				{"ratio": 0.7, "score": 1, "value": 4},
				{"ratio": 0.8, "score": 0.07, "value": 5},
				{"ratio": 1, "score": 0.5, "value": 6},
			},
		},
		{
			"name": "row3",
			"data": []map[string]float64{
				{"ratio": 1, "score": 0.2, "value": 7},
				{"ratio": 0.2, "score": 0.7, "value": 8},
				{"ratio": 0.5, "score": 0.01, "value": 9},
			},
		},
	}
	data := rowData("dotplot", matrices)
	assert.Equal(
		t,
		expected,
		data,
		"Dotplot row data not generated correctly",
	)

	// TEST: data conversion to row heatmap matrix.
	expected = []map[string]interface{}{
		{
			"name": "row1",
			"data": []map[string]float64{
				{"value": 1},
				{"value": 2},
				{"value": 3},
			},
		},
		{
			"name": "row2",
			"data": []map[string]float64{
				{"value": 4},
				{"value": 5},
				{"value": 6},
			},
		},
		{
			"name": "row3",
			"data": []map[string]float64{
				{"value": 7},
				{"value": 8},
				{"value": 9},
			},
		},
	}
	data = rowData("heatmap", matrices)
	assert.Equal(
		t,
		expected,
		data,
		"Heatmap row data not generated correctly",
	)
}
