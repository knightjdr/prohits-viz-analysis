// Packages files has functions for interacting with the file system.
package files

import (
	"os"
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
)

// CreateFolders creates all folders in the "list".
func CreateFolders(list []string) {
	for _, folder := range list {
		path := filepath.Join(".", folder)
		err := fs.Instance.MkdirAll(path, os.ModePerm)
		log.CheckError(err, true)
	}
}
