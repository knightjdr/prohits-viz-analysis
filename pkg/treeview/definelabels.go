package treeview

import "fmt"

import "github.com/knightjdr/prohits-viz-analysis/pkg/slice"

type id struct {
	sorted   string
	uniqueID string
	unsorted string
}

func getTreeLabeler(sorted, unsorted []string, prefix string) func(int) string {
	noLeafs := len(sorted)

	return func(nodeIndex int) string {
		if nodeIndex < noLeafs {
			sortedIndex := slice.IndexOfString(unsorted[nodeIndex], sorted)
			return fmt.Sprintf("%s%dX", prefix, sortedIndex)
		}
		return fmt.Sprintf("NODE%dX", nodeIndex-noLeafs+1)
	}
}
