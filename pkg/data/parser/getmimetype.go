package parser

import (
	"net/http"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
)

func getMimeTypes(files []string) []string {
	mimeTypes := make([]string, len(files))

	for i, filename := range files {
		var mimeTypeErr error
		mimeTypes[i], mimeTypeErr = getMimeType(filename)
		// If a filetype cannot be opened, log it but don't exit. Will just skip.
		log.CheckError(mimeTypeErr, false)
	}

	return mimeTypes
}

func getMimeType(filename string) (mimetype string, err error) {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		return "unknown", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	bufferLength, err := file.Read(buffer)
	if err != nil {
		return "unknown", err
	}

	mimetype = strings.Split(http.DetectContentType(buffer[:bufferLength]), ";")[0]
	return
}
