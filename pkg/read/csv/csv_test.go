package csv

import (
	"os"
	"regexp"

	"github.com/bouk/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/spf13/afero"
)

var csvText = `field1	field2	field3
A	1	C
D	2	F
G	3	I
`

var _ = Describe("CSV reader", func() {
	It("should read a csv file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/csv.txt", []byte(csvText), 0444)

		headerMap := map[string]string{
			"field1": "field-1",
			"field3": "field-3",
		}

		expected := []map[string]string{
			{"field-1": "A", "field2": "1", "field-3": "C"},
			{"field-1": "D", "field2": "2", "field-3": "F"},
			{"field-1": "G", "field2": "3", "field-3": "I"},
		}

		Expect(Read("test/csv.txt", '\t', headerMap)).To(Equal(expected))
	})
})

var _ = Describe("Open file", func() {
	It("should exit and log error when file cannot be opened", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		fakeExit := func(int) {
			panic("os.Exit called")
		}
		exitPatch := monkey.Patch(os.Exit, fakeExit)
		defer exitPatch.Unpatch()

		Expect(func() { openFile("test/missing.txt") }).To(Panic(), "should exit when missing file to open")

		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		matched, _ := regexp.MatchString("file does not exist", string(logfile))
		Expect(matched).To(BeTrue(), "should write error when missing file to open")
	})
})
