package columnparser

import (
	"regexp"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	// Mock fs.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/testfile.txt",
		[]byte("column1\tcolumn2\tcolumn3\na\tb\tc\nd\te\tf\n"),
		0444,
	)
	afero.WriteFile(
		fs.Instance,
		"test/empty.txt",
		[]byte("column1,column2,column3\n"),
		0444,
	)
	afero.WriteFile(fs.Instance, "test/error.txt", []byte(""), 0644)

	// TEST1: file gets parsed to a slice of maps.
	files := []string{"test/testfile.txt"}
	columnMap := map[string]string{
		"condition": "column1",
		"readout":   "column3",
	}
	want := []map[string]string{
		{"condition": "a", "readout": "c"},
		{"condition": "d", "readout": "f"},
	}
	parsed := ReadFile(files, columnMap, false)
	assert.Equal(t, want, parsed, "File not parsed as expected")

	// TEST2: No parsed results should panic.
	files = []string{"test/empty.txt"}
	assert.Panics(t, func() { ReadFile(files, columnMap, false) }, "File with no parsed results should panic")

	// TEST3: Invalid file type logs an error.
	files = []string{"test/testfile.txt", "test/missing.txt"}
	parsed = ReadFile(files, columnMap, false)
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	wantMessage := "file does not exist"
	matched, _ := regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged")
}
