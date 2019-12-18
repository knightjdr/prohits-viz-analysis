package convert

import "sort"

// sortLabels sorts a map of labels either alphabetically or by the index
// stored as each entry's value.
// input = []map{
//    {"C label": 2}
//		{"B label": 0},
//		{"A label": 1}
// output alphabetically: []string{"A label", "B label", "C label"}
// output by index: []string{"B label", "A label", "C label"}
func sortLabels(labels map[string]int, alphabetically bool) []string {
	sortedLabels := make([]string, len(labels))
	if alphabetically {
		index := 0
		for label := range labels {
			sortedLabels[index] = label
			index++
		}
		sort.Strings(sortedLabels)
	} else {
		for label, value := range labels {
			sortedLabels[value] = label
		}
	}
	return sortedLabels
}
