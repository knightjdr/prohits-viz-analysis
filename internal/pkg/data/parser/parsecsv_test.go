package parser

import (
	"errors"
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/spf13/afero"
)

var _ = Describe("Parse CSV", func() {
	It("should parse a single file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/testfile1.txt",
			[]byte("column1\tcolumn2\tcolumn3\na\tb\tc\n"),
			0444,
		)

		columnMap := map[string]string{
			"key1": "column1",
			"key2": "column3",
		}
		files := []string{"test/testfile1.txt"}
		filetype := []string{"text/plain"}

		expected := []map[string]string{
			{"key1": "a", "key2": "c"},
		}
		Expect(parseCSV(files, filetype, columnMap, false)).To(Equal(expected))
	})

	It("should parse two files", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/testfile1.txt",
			[]byte("column1\tcolumn2\tcolumn3\na\tb\tc\n"),
			0444,
		)
		afero.WriteFile(
			fs.Instance,
			"test/testfile2.txt",
			[]byte("column1,column2,column3\nd,e,f\n"),
			0444,
		)

		columnMap := map[string]string{
			"key1": "column1",
			"key2": "column3",
		}
		files := []string{"test/testfile1.txt", "test/testfile2.txt"}
		filetype := []string{"text/plain", "text/csv"}

		expected := []map[string]string{
			{"key1": "a", "key2": "c"},
			{"key1": "d", "key2": "f"},
		}
		Expect(parseCSV(files, filetype, columnMap, false)).To(Equal(expected))
	})

	It("should parse a file with an extra column", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/testfile1.txt",
			[]byte("column1\tcolumn2\tcolumn3\tcolumn4\na\tb\tc\n"),
			0444,
		)

		columnMap := map[string]string{
			"key1": "column1",
			"key2": "column3",
		}
		files := []string{"test/testfile1.txt"}
		filetype := []string{"text/plain"}

		expected := []map[string]string{
			{"key1": "a", "key2": "c"},
		}
		Expect(parseCSV(files, filetype, columnMap, false)).To(Equal(expected))
	})

	It("should parse a file with an missing column", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/testfile1.txt",
			[]byte("column1\tcolumn2\tcolumn3\na\tb\tc\n"),
			0444,
		)
		afero.WriteFile(
			fs.Instance,
			"test/testfile2.txt",
			[]byte("column1\tcolumn2\ng\th\n"),
			0444,
		)

		columnMap := map[string]string{
			"key1": "column1",
			"key2": "column3",
		}
		files := []string{"test/testfile2.txt", "test/testfile1.txt"}
		filetype := []string{"text/plain", "text/plain"}

		expected := []map[string]string{
			{"key1": "a", "key2": "c"},
		}
		Expect(parseCSV(files, filetype, columnMap, false)).To(Equal(expected))
	})

	It("should log error message when missing file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		columnMap := map[string]string{
			"key1": "column1",
			"key2": "column3",
		}
		files := []string{"test/missing.txt"}
		filetype := []string{"text/plain"}

		parseCSV(files, filetype, columnMap, false)
		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		expected := "file does not exist"
		matched, _ := regexp.MatchString(expected, string(logfile))
		Expect(matched).To(BeTrue())
	})

	It("should log error message when header column is missing", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/testfile1.txt",
			[]byte("column1\tcolumn2\tcolumn3\na\tb\tc\n"),
			0444,
		)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		// Mock HeaderMap
		fakeHeaderMap := func(columnMap map[string]string, header []string, ignoreMissing bool) (map[string]int, error) {
			return map[string]int{}, errors.New("missing header columns")
		}
		headerMapPatch := monkey.Patch(mapHeader, fakeHeaderMap)
		defer headerMapPatch.Unpatch()

		columnMap := map[string]string{
			"key1": "column1",
			"key2": "column3",
		}
		files := []string{"test/testfile1.txt"}
		filetype := []string{"text/plain"}

		parseCSV(files, filetype, columnMap, false)
		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		expected := "missing header columns"
		matched, _ := regexp.MatchString(expected, string(logfile))
		Expect(matched).To(BeTrue())
	})
})
