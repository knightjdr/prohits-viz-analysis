package interactive

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// Pngurl converts a png image to a data url.
func Pngurl(filename string) (url string) {
	url = "data:image/png;base64,"

	// Open png. If it can't be opend, return an empty url.
	pngfile, err := fs.Instance.Open(filename)
	if err != nil {
		logmessage.CheckError(err, false)
	} else {
		// Convert png to url.
		pngsrc, _, decodeerr := image.Decode(pngfile)
		logmessage.CheckError(decodeerr, false)
		if decodeerr != nil {
			return
		}

		// Encode image to buffer.
		buf := new(bytes.Buffer)
		encodeerr := png.Encode(buf, pngsrc)
		logmessage.CheckError(encodeerr, false)
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
