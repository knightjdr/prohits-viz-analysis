package logmessage

import (
	"errors"
	"log"
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestFiletype(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644)

	// TEST1: message logged to file.
	Write("test/logfile.txt", "test message")
	logfile, _ := afero.ReadFile(fs.Instance, "test/logfile.txt")
	want := "test message"
	matched, _ := regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "message not being logged")

	// Mock exit.
	fakeExit := func(int) {
		panic("os.Exit called")
	}
	exitPatch := monkey.Patch(os.Exit, fakeExit)
	defer exitPatch.Unpatch()

	// Mock fatal.
	fakeFatal := func(string, ...interface{}) {
		panic("log.Fatalf called")
	}
	fatalPatch := monkey.Patch(log.Fatalf, fakeFatal)
	defer fatalPatch.Unpatch()

	// Mock OpenFile.
	file, _ := fs.Instance.Open("test/logfile.txt")
	fakeOpenFile := func(*afero.MemMapFs, string, int, os.FileMode) (afero.File, error) {
		return file, errors.New("Test open error")
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(fs.Instance), "OpenFile", fakeOpenFile)
	defer monkey.UnpatchInstanceMethod(reflect.TypeOf(fs.Instance), "OpenFile")

	// TEST2: exit if file can't be opened.
	assert.PanicsWithValue(
		t,
		"log.Fatalf called",
		func() { Write("test/logfile.txt", "test message") },
		"log.Fatalf called was not called when file doesn't exist",
	)
}
