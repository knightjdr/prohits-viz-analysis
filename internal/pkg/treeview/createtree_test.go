package treeview

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
)

var _ = Describe("Write ATR file", func() {
	It("should write file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := Data{
			Filename: "test",
			Names: Names{
				Columns:         []string{"columnA", "columnB", "columnC", "columnD", "columnE", "columnF"},
				UnsortedColumns: []string{"columnA", "columnC", "columnB", "columnD", "columnE", "columnF"},
			},
			Trees: Trees{
				Column: hclust.Dendrogram{
					{Leafa: 0, Leafb: 3, Lengtha: 0.05, Lengthb: 0.05, Node: 6},
					{Leafa: 2, Leafb: 5, Lengtha: 0.075, Lengthb: 0.075, Node: 7},
					{Leafa: 1, Leafb: 6, Lengtha: 0.1, Lengthb: 0.05, Node: 8},
					{Leafa: 4, Leafb: 8, Lengtha: 0.2, Lengthb: 0.1, Node: 9},
					{Leafa: 7, Leafb: 9, Lengtha: 0.225, Lengthb: 0.1, Node: 10},
				},
			},
		}

		expected := "NODE1X\tARRY0X\tARRY3X\t1.00000\n" +
			"NODE2X\tARRY1X\tARRY5X\t0.95652\n" +
			"NODE3X\tARRY2X\tNODE1X\t0.88000\n" +
			"NODE4X\tARRY4X\tNODE3X\t0.70968\n" +
			"NODE5X\tNODE2X\tNODE4X\t0.54321\n"

		createATR(data)
		actual, _ := afero.ReadFile(fs.Instance, "test.atr")
		Expect(string(actual)).To(Equal(expected))
	})
})

var _ = Describe("Write GTR file", func() {
	It("should write file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := Data{
			Filename: "test",
			Names: Names{
				Rows:         []string{"rowA", "rowB", "rowC", "rowD", "rowE", "rowF"},
				UnsortedRows: []string{"rowA", "rowC", "rowB", "rowD", "rowE", "rowF"},
			},
			Trees: Trees{
				Row: hclust.Dendrogram{
					{Leafa: 0, Leafb: 3, Lengtha: 0.05, Lengthb: 0.05, Node: 6},
					{Leafa: 2, Leafb: 5, Lengtha: 0.075, Lengthb: 0.075, Node: 7},
					{Leafa: 1, Leafb: 6, Lengtha: 0.1, Lengthb: 0.05, Node: 8},
					{Leafa: 4, Leafb: 8, Lengtha: 0.2, Lengthb: 0.1, Node: 9},
					{Leafa: 7, Leafb: 9, Lengtha: 0.225, Lengthb: 0.1, Node: 10},
				},
			},
		}

		expected := "NODE1X\tGENE0X\tGENE3X\t1.00000\n" +
			"NODE2X\tGENE1X\tGENE5X\t0.95652\n" +
			"NODE3X\tGENE2X\tNODE1X\t0.88000\n" +
			"NODE4X\tGENE4X\tNODE3X\t0.70968\n" +
			"NODE5X\tNODE2X\tNODE4X\t0.54321\n"

		createGTR(data)
		actual, _ := afero.ReadFile(fs.Instance, "test.gtr")
		Expect(string(actual)).To(Equal(expected))
	})
})

var _ = Describe("Write tree nodes", func() {
	It("should write nodes to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		file, _ := fs.Instance.Create("test.atr")

		expected := "NODE1X\tlabel0\tlabel3\t1.00000\n" +
			"NODE2X\tlabel2\tlabel5\t0.95652\n" +
			"NODE3X\tlabel1\tlabel6\t0.88000\n" +
			"NODE4X\tlabel4\tlabel8\t0.70968\n" +
			"NODE5X\tlabel7\tlabel9\t0.54321\n"

		dendrogram := hclust.Dendrogram{
			{Leafa: 0, Leafb: 3, Lengtha: 0.05, Lengthb: 0.05, Node: 6},
			{Leafa: 2, Leafb: 5, Lengtha: 0.075, Lengthb: 0.075, Node: 7},
			{Leafa: 1, Leafb: 6, Lengtha: 0.1, Lengthb: 0.05, Node: 8},
			{Leafa: 4, Leafb: 8, Lengtha: 0.2, Lengthb: 0.1, Node: 9},
			{Leafa: 7, Leafb: 9, Lengtha: 0.225, Lengthb: 0.1, Node: 10},
		}

		label := func(index int) string {
			return fmt.Sprintf("label%d", index)
		}

		writeTreeNodes(file, dendrogram, label)
		actual, _ := afero.ReadFile(fs.Instance, "test.atr")
		Expect(string(actual)).To(Equal(expected))
	})
})

var _ = Describe("Convert distance to correlation", func() {
	It("should convert tree branch lengths to node correlation value", func() {
		dendrogram := hclust.Dendrogram{
			{Leafa: 0, Leafb: 3, Lengtha: 0.05, Lengthb: 0.05, Node: 6},
			{Leafa: 2, Leafb: 5, Lengtha: 0.075, Lengthb: 0.075, Node: 7},
			{Leafa: 1, Leafb: 6, Lengtha: 0.1, Lengthb: 0.05, Node: 8},
			{Leafa: 4, Leafb: 8, Lengtha: 0.2, Lengthb: 0.1, Node: 9},
			{Leafa: 7, Leafb: 9, Lengtha: 0.225, Lengthb: 0.1, Node: 10},
		}

		expected := []float64{1, 0.957, 0.88, 0.710, 0.543}

		actual := convertDistanceToCorrelation(dendrogram)
		for i, value := range actual {
			Expect(value).To(BeNumerically("~", expected[i], 0.001))
		}
	})
})
