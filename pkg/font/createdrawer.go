// Package font loads and creates a context for rendering text in Arial.
package font

import (
	"image"

	"golang.org/x/image/font"
)

// CreateDrawer for drawing text.
func CreateDrawer(img *image.RGBA, fontSize int, fontPath string) *font.Drawer {
	fontFace := loadFont(fontPath, fontSize)

	fg := image.Black
	return &font.Drawer{
		Dst:  img,
		Src:  fg,
		Face: fontFace,
	}
}
