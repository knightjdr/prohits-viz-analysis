// Package fs creates a filesystem to use (for easy mocking)
//
// This replaces most of the filesystem and io methods from os and io
package fs

import "github.com/spf13/afero"

// Instance contains the file system instance
var Instance = afero.NewOsFs()
