package csv

import (
	"os"
	"regexp"

	"bou.ke/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/spf13/afero"
)

var _ = Describe("Read csv file to map", func() {
	It("should read a two-column csv file to a map", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		csvText := "A\t1\n" +
			"B\t2\n" +
			"C\t3\n"
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/csv.txt", []byte(csvText), 0444)

		expected := map[string]string{
			"A": "1",
			"B": "2",
			"C": "3",
		}

		Expect(ReadToMap("test/csv.txt", '\t')).To(Equal(expected))
	})
})

var _ = Describe("Read csv file to a map of slices", func() {
	It("should read a two-column csv file to a map", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		csvText := "A\t1\n" +
			"B\t2\n" +
			"B\t3\n" +
			"C\t3\n"
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/csv.txt", []byte(csvText), 0444)

		expected := map[string][]string{
			"A": {"1"},
			"B": {"2", "3"},
			"C": {"3"},
		}

		Expect(ReadToSliceMap("test/csv.txt", '\t')).To(Equal(expected))
	})
})

var _ = Describe("Read csv file to slice via header", func() {
	It("should read a csv file using a header for mapping to keys", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		csvText := "field1\tfield2\tfield3\n" +
			"A\t1\tC\n" +
			"D\t2\tF\n" +
			"G\t3\tI\n"
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

		Expect(ReadToSliceViaHeader("test/csv.txt", '\t', headerMap)).To(Equal(expected))
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
