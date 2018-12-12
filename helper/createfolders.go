package helper

import (
	"os"
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// CreateFolders creates all folders in the "list" argument
func CreateFolders(list []string) {
	for _, folder := range list {
		path := filepath.Join(".", folder)
		err := fs.Instance.MkdirAll(path, os.ModePerm)
		logmessage.CheckError(err, true)
	}
}
