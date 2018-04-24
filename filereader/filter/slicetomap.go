package filter

// SliceToMap converts a slice to a map of booleans.
func SliceToMap(slice []string) (mapped map[string]bool) {
	mapped = make(map[string]bool)
	for _, element := range slice {
		mapped[element] = true
	}
	return
}
