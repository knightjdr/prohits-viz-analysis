// Package slice contains functions for manipulating slices
package slice

// ConvertToBoolMap converts a slice to a map of booleans.
func ConvertToBoolMap(slice []string) (mapped map[string]bool) {
	mapped = make(map[string]bool)

	for _, element := range slice {
		mapped[element] = true
	}

	return
}

// ConvertToIntMap converts a slice to a map of indices.
func ConvertToIntMap(slice []string) (mapped map[string]int) {
	mapped = make(map[string]int)

	for i, element := range slice {
		mapped[element] = i
	}

	return
}
