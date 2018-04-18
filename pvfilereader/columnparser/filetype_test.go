package columnparser

import (
	"regexp"
	"testing"

	"github.com/spf13/afero"
)

// mock filesystem
var appFS = afero.NewOsFs()

func TestFiletype(t *testing.T) {
	// create test directory and files
	appFS.MkdirAll("test", 0755)
	afero.WriteFile(appFS, "test/testfile.txt", []byte("a\tb\tc\n"), 0444)
	afero.WriteFile(appFS, "test/logfile.txt", []byte(""), 0644)

	// mimetype is returned
	want := "text/plain"
	mimetype := Filetype("test/testfile.txt", "test/logfile.txt")
	if mimetype != want {
		t.Errorf("mimetype == %v, want %v", mimetype, want)
	}

	// unknown is returned when file can't be opened and message is logged
	want = "unknown"
	mimetype = Filetype("test/missing", "test/logfile.txt")
	if mimetype != want {
		t.Errorf("mimetype == %v, want %v", mimetype, want)
	}
	logfile, _ := afero.ReadFile(appFS, "test/logfile.txt")
	want = "test/missing: could not be opened"
	matched, _ := regexp.MatchString(want, string(logfile))
	if !matched {
		t.Errorf("not logging correct message == %v, want %v", string(logfile), want)
	}
}
