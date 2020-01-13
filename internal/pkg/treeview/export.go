// Package treeview exports a matrix and tree to java treeview format.
package treeview

import (
	"github.com/knightjdr/hclust"
)

// Data for treeview conversion.
type Data struct {
	ColumnLabeler func(int) string
	Filename      string
	Matrix        [][]float64
	Names         Names
	RowLabeler    func(int) string
	Trees         Trees
}

// Names contains row and column names.
type Names struct {
	Columns         []string
	Rows            []string
	UnsortedColumns []string
	UnsortedRows    []string
}

// Trees contains row and column dendorgrams.
type Trees struct {
	Column []hclust.SubCluster
	Row    []hclust.SubCluster
}

// Export a matrix to treeview format.
func Export(data Data) {
	createCDT(data)
	createATR(data)
	createGTR(data)
}
