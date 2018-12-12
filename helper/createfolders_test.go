package helper

import (
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestCreateFolders(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory.
	fs.Instance.MkdirAll("test", 0755)

	// TEST: create two folder in a parent folder that exists
	CreateFolders([]string{"test/folder1", "test/folder2"})
	exists, _ := afero.DirExists(fs.Instance, "test/folder1")
	assert.True(t, exists, "Test folder 1 not created")
	exists, _ = afero.DirExists(fs.Instance, "test/folder2")
	assert.True(t, exists, "Test folder 2 not created")
}
