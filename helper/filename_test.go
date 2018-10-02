package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilename(t *testing.T) {
	slice := []string{"file.txt", "some/path/file2.svg", "/path/file3.png"}

	// TEST1: returns slice of filenames
	want := []string{"file.txt", "file2.svg", "file3.png"}
	assert.Equal(t, want, Filename(slice), "Slice not converted to filenames")
}
