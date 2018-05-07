package svg

import (
	"fmt"
	"testing"
)

func TestHeatmap(t *testing.T) {
	columns := []string{"bait1", "bait2", "bait3"}
	matrix := [][]float64{
		{25, 5, 50.2},
		{100, 30, 7},
		{5, 2.3, 8},
	}
	options := map[string]interface{}{
		"colorSpace":       "blueBlack",
		"maximumAbundance": float64(50),
	}
	rows := []string{"prey1", "prey2", "prey3"}

	// TEST1: create svg.
	svg := Heatmap(matrix, columns, rows, options)
	fmt.Println(svg)
}
