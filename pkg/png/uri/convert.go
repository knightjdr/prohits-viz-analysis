// Package uri converts a png to a data uri.
package uri

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
)

// Convert a png image to a data uri.
func Convert(filename string) (url string) {
	url = "data:image/png;base64,"

	// Open png. If it can't be opened, return an empty url.
	pngfile, err := fs.Instance.Open(filename)
	if err != nil {
		log.CheckError(err, false)
	} else {
		// Convert png to url.
		pngsrc, _, decodeerr := image.Decode(pngfile)
		log.CheckError(decodeerr, false)
		if decodeerr != nil {
			return
		}

		// Encode image to buffer.
		buf := new(bytes.Buffer)
		encodeerr := png.Encode(buf, pngsrc)
		log.CheckError(encodeerr, false)
		if encodeerr != nil {
			return
		}

		// Create url.
		base64Img := base64.StdEncoding.EncodeToString(buf.Bytes())
		url = fmt.Sprintf("data:image/png;base64,%s", base64Img)
	}
	defer pngfile.Close()
	return
}
