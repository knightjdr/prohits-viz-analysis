package helper

// MapKeysFloat64 returns the keys from a map that has float64 values
func MapKeysFloat64(hash map[string]float64) []string {
	keys := make([]string, len(hash))

	i := 0
	for key := range hash {
		keys[i] = key
		i++
	}
	return keys
}
