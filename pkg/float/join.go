// Package float defines functions for float transformations.
package float

import (
	"strconv"
	"strings"
)

// Join a slice of float64 to a string.
func Join(s []float64, sep string) string {
	str := make([]string, len(s))

	for i, value := range s {
		str[i] = strconv.FormatFloat(value, 'f', -1, 64)
	}

	return strings.Join(str, sep)
}
