// Package slice contains functions for manipulating slices
package slice

// ConvertToMap converts a slice to a map of booleans.
func ConvertToMap(slice []string) (mapped map[string]bool) {
	mapped = make(map[string]bool)

	for _, element := range slice {
		mapped[element] = true
	}

	return
}
