// Package parse has functions for parsing data columns.
package parse

import (
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat"
)

// PipeSeparatedStringToMean parses a string containing pipe-separated float64 and returns average.
func PipeSeparatedStringToMean(str string) float64 {
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

	return stat.Mean(parsedFloats, nil)
}

// PipeSeparatedStringToArray parses a string containing pipe-separated float64 and returns array of values.
func PipeSeparatedStringToArray(str string) []float64 {
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

	return parsedFloats
}
