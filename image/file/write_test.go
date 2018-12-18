package file

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	// Mock fs.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and file.
	fs.Instance.MkdirAll("test", 0755)
	file, _ := fs.Instance.Create("test/test1.txt")

	// TEST: non nil file should cause string to be written to it
	outstring := make([]string, 0)
	writestring := "test string\n"
	Write(&outstring, writestring, file)
	bytes, _ := afero.ReadFile(fs.Instance, "test/test1.txt")
	assert.Equal(t, writestring, string(bytes), "String not written to file")

	// TEST: nil file should append string to outstring
	Write(&outstring, writestring, nil)
	newstring := helper.StringConcat(outstring)
	assert.Equal(t, writestring, newstring, "String not appended to slice")
}
