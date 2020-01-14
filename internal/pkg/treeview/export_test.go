package treeview_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/treeview"
)

var _ = Describe("Export treeview files", func() {
	It("should write files", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := Data{
			Filename: "test",
			Matrix: [][]float64{
				{1, 2, 3, 1, 2, 3},
				{1, 2, 3, 1, 2, 3},
				{4, 5, 6, 4, 5, 6},
				{4, 5, 6, 4, 5, 6},
				{7, 8, 9, 7, 8, 9},
				{7, 8, 9, 7, 8, 9},
			},
			Names: Names{
				Columns:         []string{"columnA", "columnB", "columnC", "columnD", "columnE", "columnF"},
				Rows:            []string{"rowA", "rowB", "rowC", "rowD", "rowE", "rowF"},
				UnsortedColumns: []string{"columnA", "columnC", "columnB", "columnD", "columnE", "columnF"},
				UnsortedRows:    []string{"rowA", "rowC", "rowB", "rowD", "rowE", "rowF"},
			},
			Trees: Trees{
				Column: hclust.Dendrogram{
					{Leafa: 0, Leafb: 3, Lengtha: 0.05, Lengthb: 0.05, Node: 6},
					{Leafa: 2, Leafb: 5, Lengtha: 0.075, Lengthb: 0.075, Node: 7},
					{Leafa: 1, Leafb: 6, Lengtha: 0.1, Lengthb: 0.05, Node: 8},
					{Leafa: 4, Leafb: 8, Lengtha: 0.2, Lengthb: 0.1, Node: 9},
					{Leafa: 7, Leafb: 9, Lengtha: 0.225, Lengthb: 0.1, Node: 10},
				},
				Row: hclust.Dendrogram{
					{Leafa: 0, Leafb: 3, Lengtha: 0.05, Lengthb: 0.05, Node: 6},
					{Leafa: 2, Leafb: 5, Lengtha: 0.075, Lengthb: 0.075, Node: 7},
					{Leafa: 1, Leafb: 6, Lengtha: 0.1, Lengthb: 0.05, Node: 8},
					{Leafa: 4, Leafb: 8, Lengtha: 0.2, Lengthb: 0.1, Node: 9},
					{Leafa: 7, Leafb: 9, Lengtha: 0.225, Lengthb: 0.1, Node: 10},
				},
			},
		}

		Export(data)

		expectedCDT := "GID\tUNIQID\tNAME\tGWEIGHT\tcolumnA\tcolumnB\tcolumnC\tcolumnD\tcolumnE\tcolumnF\n" +
			"AID\t\t\t\tARRY0X\tARRY1X\tARRY2X\tARRY3X\tARRY4X\tARRY5X\n" +
			"EWEIGHT\t\t\t\t1\t1\t1\t1\t1\t1\n" +
			"GENE0X\trowA\trowA\t1\t1.00000\t2.00000\t3.00000\t1.00000\t2.00000\t3.00000\n" +
			"GENE1X\trowB\trowB\t1\t1.00000\t2.00000\t3.00000\t1.00000\t2.00000\t3.00000\n" +
			"GENE2X\trowC\trowC\t1\t4.00000\t5.00000\t6.00000\t4.00000\t5.00000\t6.00000\n" +
			"GENE3X\trowD\trowD\t1\t4.00000\t5.00000\t6.00000\t4.00000\t5.00000\t6.00000\n" +
			"GENE4X\trowE\trowE\t1\t7.00000\t8.00000\t9.00000\t7.00000\t8.00000\t9.00000\n" +
			"GENE5X\trowF\trowF\t1\t7.00000\t8.00000\t9.00000\t7.00000\t8.00000\t9.00000\n"
		actualCDT, _ := afero.ReadFile(fs.Instance, "test.cdt")
		Expect(string(actualCDT)).To(Equal(expectedCDT))

		expectedATR := "NODE1X\tARRY0X\tARRY3X\t1.00000\n" +
			"NODE2X\tARRY1X\tARRY5X\t0.95652\n" +
			"NODE3X\tARRY2X\tNODE1X\t0.88000\n" +
			"NODE4X\tARRY4X\tNODE3X\t0.70968\n" +
			"NODE5X\tNODE2X\tNODE4X\t0.54321\n"
		actualATR, _ := afero.ReadFile(fs.Instance, "test.atr")
		Expect(string(actualATR)).To(Equal(expectedATR))

		expectedGTR := "NODE1X\tGENE0X\tGENE3X\t1.00000\n" +
			"NODE2X\tGENE1X\tGENE5X\t0.95652\n" +
			"NODE3X\tGENE2X\tNODE1X\t0.88000\n" +
			"NODE4X\tGENE4X\tNODE3X\t0.70968\n" +
			"NODE5X\tNODE2X\tNODE4X\t0.54321\n"
		actualGTR, _ := afero.ReadFile(fs.Instance, "test.gtr")
		Expect(string(actualGTR)).To(Equal(expectedGTR))
	})
})
