package columnparser

import (
	"testing"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/logmessage"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestFiletype(t *testing.T) {
	// mock fs
	oldFs := appFs
	defer func() { appFs = oldFs }()
	appFs = afero.NewMemMapFs()

	// create test directory and files
	appFs.MkdirAll("test", 0755)
	afero.WriteFile(appFs, "test/testfile.txt", []byte("a\tb\tc\n"), 0444)
	afero.WriteFile(appFs, "test/unreadable.txt", []byte(""), 0444) // unreadable because empty
	afero.WriteFile(appFs, "test/logfile.txt", []byte(""), 0644)

	// TEST1: mimetype is returned
	want := "text/plain"
	mimetype, err := Filetype("test/testfile.txt", "test/logfile.txt")
	assert.Nil(t, err, "expected no error, got %s", err)
	assert.Equal(t, want, mimetype, "mimetype == %v, want %v", mimetype, want)

	// mock logger
	fakeLog := func(string, string) {}
	write := logmessage.Write
	writePatch := monkey.Patch(write, fakeLog)
	defer writePatch.Unpatch()

	// TEST2: unknown is returned when file can't be opened
	want = "unknown"
	mimetype, err = Filetype("test/missing", "test/logfile.txt")
	assert.NotNil(t, err, "expected an error when the file cannot be found")
	assert.Equal(t, want, mimetype, "mimetype == %v, want %v", mimetype, want)

	// TEST3: unreadable file returns unknown file type
	want = "unknown"
	mimetype, err = Filetype("test/unreadable.txt", "test/logfile.txt")
	assert.NotNil(t, err, "expected an error when the file cannot be found")
	assert.Equal(t, want, mimetype, "mimetype == %v, want %v", mimetype, want)
}
