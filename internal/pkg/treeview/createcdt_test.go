package treeview

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
)

var _ = Describe("Write CDT file", func() {
	It("should write file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := Data{
			Filename: "test",
			Matrix: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			Names: Names{
				Columns: []string{"columnA", "columnB", "columnC"},
				Rows:    []string{"row1", "row2", "row3"},
			},
		}

		expected := "GID\tUNIQID\tNAME\tGWEIGHT\tcolumnA\tcolumnB\tcolumnC\n" +
			"AID\t\t\t\tARRY0X\tARRY1X\tARRY2X\n" +
			"EWEIGHT\t\t\t\t1\t1\t1\n" +
			"GENE0X\trow1\trow1\t1\t1.00000\t2.00000\t3.00000\n" +
			"GENE1X\trow2\trow2\t1\t4.00000\t5.00000\t6.00000\n" +
			"GENE2X\trow3\trow3\t1\t7.00000\t8.00000\t9.00000\n"

		createCDT(data)
		actual, _ := afero.ReadFile(fs.Instance, "test.cdt")
		Expect(string(actual)).To(Equal(expected))
	})
})

var _ = Describe("Write CDT header", func() {
	It("should write header to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		columns := []string{"columnA", "columnB", "columnC"}
		file, _ := fs.Instance.Create("test.cdt")

		expected := "GID\tUNIQID\tNAME\tGWEIGHT\tcolumnA\tcolumnB\tcolumnC\n"

		writeCDTHeader(file, columns)
		actual, _ := afero.ReadFile(fs.Instance, "test.cdt")
		Expect(string(actual)).To(Equal(expected))
	})
})

var _ = Describe("Write array IDs", func() {
	It("should write ARRY row to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		file, _ := fs.Instance.Create("test.cdt")

		expected := "AID\t\t\t\tARRY0X\tARRY1X\tARRY2X\n"

		writeArrayIDs(file, 3)
		actual, _ := afero.ReadFile(fs.Instance, "test.cdt")
		Expect(string(actual)).To(Equal(expected))
	})
})

var _ = Describe("Write CDT EWEIGHT", func() {
	It("should write EWEIGHT row to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		file, _ := fs.Instance.Create("test.cdt")

		expected := "EWEIGHT\t\t\t\t1\t1\t1\n"

		writeCDTEweight(file, 3)
		actual, _ := afero.ReadFile(fs.Instance, "test.cdt")
		Expect(string(actual)).To(Equal(expected))
	})
})

var _ = Describe("Write CDT matrix", func() {
	It("should write matrix to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := Data{
			Matrix: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			Names: Names{
				Rows: []string{"row1", "row2", "row3"},
			},
		}
		file, _ := fs.Instance.Create("test.cdt")

		expected := "GENE0X\trow1\trow1\t1\t1.00000\t2.00000\t3.00000\n" +
			"GENE1X\trow2\trow2\t1\t4.00000\t5.00000\t6.00000\n" +
			"GENE2X\trow3\trow3\t1\t7.00000\t8.00000\t9.00000\n"

		writeCDTMatrix(file, data)
		actual, _ := afero.ReadFile(fs.Instance, "test.cdt")
		Expect(string(actual)).To(Equal(expected))
	})
})
