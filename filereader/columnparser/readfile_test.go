package columnparser

import (
	"regexp"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	// mock fs
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// create test directory and files
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
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644)

	// TEST1: file gets parsed to a slice of maps
	files := []string{"test/testfile.txt"}
	columnMap := map[string]string{
		"bait": "column1",
		"prey": "column3",
	}
	want := []map[string]string{
		{"bait": "a", "prey": "c"},
		{"bait": "d", "prey": "f"},
	}
	parsed, err := ReadFile(files, columnMap, "test/logfile.txt")
	assert.Nil(t, err, "Reading file should not produce an error")
	assert.Equal(t, want, parsed, "File not parsed as expected")

	// TEST2: No parsed results returns an error
	files = []string{"test/empty.txt"}
	parsed, err = ReadFile(files, columnMap, "test/logfile.txt")
	assert.NotNil(t, err, "File with no parsed results should return error")

	// TEST3: Invalid file type logs an error
	files = []string{"test/missing.txt"}
	parsed, _ = ReadFile(files, columnMap, "test/logfile.txt")
	logfile, _ := afero.ReadFile(fs.Instance, "test/logfile.txt")
	wantMessage := "could not be opened"
	matched, _ := regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged")
}
