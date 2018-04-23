package columnparser

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestFileType(t *testing.T) {
	// mock fs
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// create test directory and files
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(fs.Instance, "test/testfile.txt", []byte("a\tb\tc\n"), 0444)
	afero.WriteFile(fs.Instance, "test/unreadable.txt", []byte(""), 0444) // unreadable because empty

	// TEST1: mimetype is returned
	mimetype, err := FileType("test/testfile.txt", "test/logfile.txt")
	want := "text/plain"
	assert.Nil(t, err, "Expected no error when reading valid file")
	assert.Equal(t, want, mimetype, "Mimetype not correct")

	// TEST2: unknown is returned when file can't be opened
	mimetype, err = FileType("test/missing", "test/logfile.txt")
	want = "unknown"
	assert.NotNil(t, err, "Expected an error when the file cannot be found")
	assert.Equal(t, want, mimetype, "Mimetype not correct")

	// TEST3: unreadable file returns unknown file type
	mimetype, err = FileType("test/unreadable.txt", "test/logfile.txt")
	want = "unknown"
	assert.NotNil(t, err, "Expected an error when the file is unreadable")
	assert.Equal(t, want, mimetype, "Mimetype not correct")
}
