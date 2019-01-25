package columnparser

import (
	"errors"
	"regexp"
	"testing"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestParseCsv(t *testing.T) {
	// Mock fs.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// create test directory and files.
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
	afero.WriteFile(
		fs.Instance,
		"test/testfile3.txt",
		[]byte("column1\tcolumn2\ng\th\n"),
		0444,
	)
	afero.WriteFile(
		fs.Instance,
		"test/testfile4.txt",
		[]byte("column1\tcolumn2\tcolumn3\tcolumn4\na\tb\tc\n"),
		0444,
	)
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

	// Reusable vars.
	columnMap := map[string]string{
		"key1": "column1",
		"key2": "column3",
	}

	// TEST: a single file returns the expected parsed array.
	files := []string{"test/testfile1.txt"}
	filetype := []string{"text/plain"}
	data := ParseCsv(files, filetype, columnMap, false)
	want := []map[string]string{
		{"key1": "a", "key2": "c"},
	}
	assert.Equal(
		t,
		want,
		data,
		"Processed file does not return correct data array",
	)

	// TEST: two files return the expected parsed array.
	files = []string{"test/testfile1.txt", "test/testfile2.txt"}
	filetype = []string{"text/plain", "text/csv"}
	data = ParseCsv(files, filetype, columnMap, false)
	want = []map[string]string{
		{"key1": "a", "key2": "c"},
		{"key1": "d", "key2": "f"},
	}
	assert.Equal(
		t,
		want,
		data,
		"Processed files do not return correct data array",
	)

	// TEST: a file with an extra header column is still parsed.
	files = []string{"test/testfile4.txt"}
	filetype = []string{"text/plain"}
	data = ParseCsv(files, filetype, columnMap, false)
	want = []map[string]string{
		{"key1": "a", "key2": "c"},
	}
	assert.Equal(
		t,
		want,
		data,
		"Processed file does not return correct data array",
	)

	// TEST: file with missing header column is skipped.
	files = []string{"test/testfile3.txt", "test/testfile1.txt"}
	filetype = []string{"text/plain", "text/plain"}
	data = ParseCsv(files, filetype, columnMap, false)
	want = []map[string]string{
		{"key1": "a", "key2": "c"},
	}
	assert.Equal(
		t,
		want,
		data,
		"Processed files do not return correct data array",
	)

	// TEST: missing file logs message (intergration with logger).
	files = []string{"test/missing.txt"}
	filetype = []string{"text/plain"}
	ParseCsv(files, filetype, columnMap, false)
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	wantMessage := "file does not exist"
	matched, _ := regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "message not being logged")

	// Mock HeaderMap
	fakeHeaderMap := func(columnMap map[string]string, header []string, ignoreMissing bool) (map[string]int, error) {
		return map[string]int{}, errors.New("Missing header columns")
	}
	headerMapPatch := monkey.Patch(HeaderMap, fakeHeaderMap)

	// TEST: header error logs (intergration with logger).
	afero.WriteFile(fs.Instance, "test/error.txt", []byte(""), 0644) // clear log
	files = []string{"test/testfile1.txt"}
	filetype = []string{"text/plain"}
	ParseCsv(files, filetype, columnMap, false)
	logfile, _ = afero.ReadFile(fs.Instance, "error.txt")
	wantMessage = "Missing header columns"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "message not being logged")
	headerMapPatch.Unpatch() // Unmock HeaderMap.
}
