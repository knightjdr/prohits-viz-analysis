// Package parse has functions for parsing data columns.
package parse

import (
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/stats"
)

// PipeSeparatedFloat parses a string containing pipe-separated float64.
func PipeSeparatedFloat(str string) float64 {
	arr := strings.Split(str, "|")

	parsedFloats := make([]float64, 0)
	for _, str := range arr {
		value, err := strconv.ParseFloat(str, 64)
		if err == nil {
			parsedFloats = append(parsedFloats, value)
		} else {
			parsedFloats = append(parsedFloats, 0)
		}
	}

	return stats.MeanFloat(parsedFloats)
}
