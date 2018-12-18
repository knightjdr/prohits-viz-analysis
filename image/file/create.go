/*Package file has methods for opening and writing to a file for creating images*/
package file

import (
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
)

// Create a file for writing
func Create(path string) (file afero.File, err error) {
	if path != "" {
		file, err = fs.Instance.Create(path)
	}
	return file, err
}
