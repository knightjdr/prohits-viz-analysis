package logmessage

import (
	"errors"
	"log"
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/fs"
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
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644)

	// TEST1: message logged to file
	Write("test/logfile.txt", "test message")
	logfile, _ := afero.ReadFile(fs.Instance, "test/logfile.txt")
	want := "test message"
	matched, _ := regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "message not being logged")

	// mock exit
	fakeExit := func(int) {
		panic("os.Exit called")
	}
	exitPatch := monkey.Patch(os.Exit, fakeExit)
	defer exitPatch.Unpatch()

	// TEST2: exit if no file specified
	assert.PanicsWithValue(
		t,
		"os.Exit called",
		func() { Write("", "test message") },
		"os.Exit was not called when missing file",
	)

	// mock fatal
	fakeFatal := func(string, ...interface{}) {
		panic("log.Fatalf called")
	}
	fatalPatch := monkey.Patch(log.Fatalf, fakeFatal)
	defer fatalPatch.Unpatch()

	// mock OpenFile
	file, _ := fs.Instance.Open("test/logfile.txt")
	fakeOpenFile := func(*afero.MemMapFs, string, int, os.FileMode) (afero.File, error) {
		return file, errors.New("Test open error")
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(fs.Instance), "OpenFile", fakeOpenFile)
	defer monkey.UnpatchInstanceMethod(reflect.TypeOf(fs.Instance), "OpenFile")

	// exit if file can't be opened
	assert.PanicsWithValue(
		t,
		"log.Fatalf called",
		func() { Write("test/logfile.txt", "test message") },
		"log.Fatalf called was not called when file doesn't exist",
	)
}
