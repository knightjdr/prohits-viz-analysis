package logmessage

import (
	"errors"
	"regexp"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestCheckError(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

	// TEST1: message logged to file.
	err := errors.New("test message")
	CheckError(err, false)
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	want := "test message"
	matched, _ := regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "Message not being logged")

	// TEST2: should panic
	assert.Panics(t, func() { CheckError(err, true) }, "Did not panic")
}
