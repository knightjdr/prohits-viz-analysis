package float

import "strconv"

// RemoveTrailingZeros removes trailing zeros from a float64.
func RemoveTrailingZeros(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
