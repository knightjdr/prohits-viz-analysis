package file

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	// Mock fs.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory.
	fs.Instance.MkdirAll("test", 0755)

	// TEST: no path should return nil file and error
	file, err := Create("")
	assert.Nil(t, err, "Expected no error when no path specified")
	assert.Nil(t, file, "File should be nil when no path specified")

	// TEST: valid path should return file that exists and no error
	file, err = Create("test/test1.txt")
	assert.Nil(t, err, "Expected no error when no path specified")
	assert.NotNil(t, file, "File should not be nil when valid path specified")
}
