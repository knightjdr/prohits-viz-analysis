package columnparser

import (
	"errors"
	"regexp"
	"testing"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestParsecsv(t *testing.T) {
	// mock fs
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// create test directory and files
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
	afero.WriteFile(fs.Instance, "test/unreadable.txt", []byte(""), 0444)
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644)

	// reusable vars
	columnMap := map[string]string{
		"key1": "column1",
		"key2": "column3",
	}

	// TEST1: a single file return the expected parsed array
	files := []string{"test/testfile1.txt"}
	filetype := []string{"text/plain"}
	data := Parsecsv(files, filetype, columnMap, "test/logfile.txt")
	want := []map[string]string{
		{"key1": "a", "key2": "c"},
	}
	assert.Equal(
		t,
		want,
		data,
		"Processed file does not return correct data array",
	)

	// TEST2: two files return the expected parsed array
	files = []string{"test/testfile1.txt", "test/testfile2.txt"}
	filetype = []string{"text/plain", "text/csv"}
	data = Parsecsv(files, filetype, columnMap, "test/logfile.txt")
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

	// TEST3: file with missing header column is skipped
	files = []string{"test/testfile3.txt", "test/testfile1.txt"}
	filetype = []string{"text/plain", "text/plain"}
	data = Parsecsv(files, filetype, columnMap, "test/logfile.txt")
	want = []map[string]string{
		{"key1": "a", "key2": "c"},
	}
	assert.Equal(
		t,
		want,
		data,
		"Processed files do not return correct data array",
	)

	// TEST4: missing file logs message (intergration with logger)
	files = []string{"test/missing.txt"}
	filetype = []string{"text/plain"}
	Parsecsv(files, filetype, columnMap, "test/logfile.txt")
	logfile, _ := afero.ReadFile(fs.Instance, "test/logfile.txt")
	wantMessage := "could not be opened"
	matched, _ := regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "message not being logged")

	// mock Headermap
	fakeHeadermap := func(columnMap map[string]string, header []string) (map[string]int, error) {
		return map[string]int{}, errors.New("Headermap error")
	}
	headerMapPatch := monkey.Patch(Headermap, fakeHeadermap)

	// TEST5: header error logs (intergration with logger)
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644) // clear log
	files = []string{"test/testfile1.txt"}
	filetype = []string{"text/plain"}
	Parsecsv(files, filetype, columnMap, "test/logfile.txt")
	logfile, _ = afero.ReadFile(fs.Instance, "test/logfile.txt")
	wantMessage = "missing header columns"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "message not being logged")
	headerMapPatch.Unpatch() // unmock Headermap
}
