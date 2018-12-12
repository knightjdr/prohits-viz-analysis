package helper

// SliceContains checks if a string is found in a list of strings.
func SliceContains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
