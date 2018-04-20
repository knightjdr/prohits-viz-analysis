package filter

// Slicetomap converts a slice to a map of booleans
func Slicetomap(slice []string) map[string]bool {
	mapped := make(map[string]bool)
	for _, element := range slice {
		mapped[element] = true
	}
	return mapped
}
