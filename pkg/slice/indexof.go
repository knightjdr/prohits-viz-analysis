package slice

// IndexOfString returns the index of a value in a slice.
func IndexOfString(str string, s []string) int {
	for key, value := range s {
		if str == value {
			return key
		}
	}
	return -1
}
