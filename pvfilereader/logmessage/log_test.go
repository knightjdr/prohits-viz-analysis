package logmessage

import (
	"errors"
	"log"
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/bouk/monkey"
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
	afero.WriteFile(appFs, "test/logfile.txt", []byte(""), 0644)

	// TEST1: message logged to file
	Write("test/logfile.txt", "test message")
	logfile, _ := afero.ReadFile(appFs, "test/logfile.txt")
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
	file, _ := appFs.Open("test/logfile.txt")
	fakeOpenFile := func(*afero.MemMapFs, string, int, os.FileMode) (afero.File, error) {
		return file, errors.New("Test open error")
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(appFs), "OpenFile", fakeOpenFile)
	defer monkey.UnpatchInstanceMethod(reflect.TypeOf(appFs), "OpenFile")

	// exit if file can't be opened
	assert.PanicsWithValue(
		t,
		"log.Fatalf called",
		func() { Write("test/logfile.txt", "test message") },
		"log.Fatalf called was not called when file doesn't exist",
	)
}
