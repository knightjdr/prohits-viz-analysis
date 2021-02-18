// Package slice contains methods for slice conversion and modification
package slice

// Dict converts a slice into a dictionary/hash.
func Dict(slice []string) map[string]bool {
	dict := make(map[string]bool, len(slice))
	for _, val := range slice {
		dict[val] = true
	}
	return dict
}

// DictInt converts a slice into a dictionary/hash.
func DictInt(slice []int) map[int]bool {
	dict := make(map[int]bool, len(slice))
	for _, val := range slice {
		dict[val] = true
	}
	return dict
}
