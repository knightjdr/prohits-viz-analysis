package columnparser

import (
	"regexp"
	"testing"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/fs"
	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/logmessage"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestFiletype(t *testing.T) {
	// mock fs
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// create test directory and files
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(fs.Instance, "test/testfile.txt", []byte("a\tb\tc\n"), 0444)
	afero.WriteFile(fs.Instance, "test/unreadable.txt", []byte(""), 0444) // unreadable because empty
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644)

	// TEST1: mimetype is returned
	mimetype, err := Filetype("test/testfile.txt", "test/logfile.txt")
	want := "text/plain"
	assert.Nil(t, err, "expected no error, got %s", err)
	assert.Equal(t, want, mimetype, "mimetype == %v, want %v", mimetype, want)

	// mock logger (unmock before TEST4)
	fakeLog := func(string, string) {}
	write := logmessage.Write
	writePatch := monkey.Patch(write, fakeLog)

	// TEST2: unknown is returned when file can't be opened
	mimetype, err = Filetype("test/missing", "test/logfile.txt")
	want = "unknown"
	assert.NotNil(t, err, "expected an error when the file cannot be found")
	assert.Equal(t, want, mimetype, "mimetype == %v, want %v", mimetype, want)

	// TEST3: unreadable file returns unknown file type
	mimetype, err = Filetype("test/unreadable.txt", "test/logfile.txt")
	want = "unknown"
	assert.NotNil(t, err, "expected an error when the file is unreadable")
	assert.Equal(t, want, mimetype, "mimetype == %v, want %v", mimetype, want)

	// TEST4: intergration with logger
	writePatch.Unpatch()
	mimetype, err = Filetype("test/missing", "test/logfile.txt")
	logfile, _ := afero.ReadFile(fs.Instance, "test/logfile.txt")
	want = "could not be opened"
	matched, _ := regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "message not being logged")
}
