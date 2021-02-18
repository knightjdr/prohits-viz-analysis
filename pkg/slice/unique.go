package slice

import "sort"

// UniqueInts returns only the unique values in a slice of ints.
func UniqueInts(slice []int) []int {
	dict := DictInt(slice)

	unique := make([]int, len(dict))
	i := 0
	for key := range dict {
		unique[i] = key
		i++
	}
	sort.Ints(unique)

	return unique
}

// UniqueStrings returns only the unique values in a slice of strings.
func UniqueStrings(slice []string) []string {
	dict := Dict(slice)

	unique := make([]string, len(dict))
	i := 0
	for key := range dict {
		unique[i] = key
		i++
	}
	sort.Strings(unique)

	return unique
}
