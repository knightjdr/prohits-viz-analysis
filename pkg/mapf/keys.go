// Package mapf contains functions for manipulating maps.
package mapf

// KeysIntBool returns the keys of a map[int]float64.
func KeysIntBool(m map[int]bool) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

// KeysIntFloat returns the keys of a map[int]float64.
func KeysIntFloat(m map[int]float64) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

// KeysStringBool returns the keys of a map[string]bool.
func KeysStringBool(m map[string]bool) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

// KeysStringFloat returns the keys of a map[string]float64.
func KeysStringFloat(m map[string]float64) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

// KeysStringString returns the keys of a map[string]string.
func KeysStringString(m map[string]string) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}
