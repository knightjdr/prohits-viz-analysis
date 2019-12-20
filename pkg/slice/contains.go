package slice

// ContainsInt checks if an int is found in a slice of ints.
func ContainsInt(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// ContainsString checks if a string is found in a slice of strings.
func ContainsString(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
