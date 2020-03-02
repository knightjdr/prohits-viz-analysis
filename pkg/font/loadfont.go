package font

import (
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

func loadFont(fontPath string, fontSize int) font.Face {
	if fontPath == "" {
		return basicfont.Face7x13
	}

	fontBytes, err := ioutil.ReadFile(fontPath)
	log.CheckError(err, true)

	parsedFont, err := freetype.ParseFont(fontBytes)
	log.CheckError(err, true)

	// Scale font size for 300 dpi (72 / 300 = 0.24).
	scaledFontSize := float64(fontSize) * 0.24

	return truetype.NewFace(parsedFont, &truetype.Options{
		DPI:  300,
		Size: scaledFontSize,
	})
}
