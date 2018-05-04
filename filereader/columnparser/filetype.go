package columnparser

import (
	"net/http"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/fs"
)

// FileType returns the mime type of a file.
func FileType(filename string) (mimetype string, err error) {
	// Open file.
	file, err := fs.Instance.Open(filename)
	if err != nil {
		// Return unknown if file cannot be opened.
		return "unknown", err
	}
	defer file.Close()

	// Read file and determine type.
	buffer := make([]byte, 512)
	bufferLength, err := file.Read(buffer)
	if err != nil {
		// return unknown if file cannot be read.
		return "unknown", err
	}

	// Split mimetype at first semi colon.
	mimetype = strings.Split(http.DetectContentType(buffer[:bufferLength]), ";")[0]
	return
}
