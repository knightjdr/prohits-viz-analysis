package helper

import (
	"sort"
	"strings"
)

type less func()

// SortStringSlice sorts a slice of strings according to direction,
// either "asc" for ascending or "des" for descending
func SortStringSlice(data []string, direction string) []string {
	if direction == "asc" {
		sort.Slice(data, func(i, j int) bool { return strings.ToLower(data[i]) < strings.ToLower(data[j]) })
	} else {
		sort.Slice(data, func(i, j int) bool { return strings.ToLower(data[i]) > strings.ToLower(data[j]) })
	}
	return data
}
