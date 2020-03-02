package heatmap

import (
	"image"

	"github.com/knightjdr/prohits-viz-analysis/pkg/font"
	"golang.org/x/image/math/fixed"
)

func drawAnnotations(img *image.RGBA, h *Heatmap) {
	if len(h.Annotations.List) > 0 {
		fontSize := h.Annotations.FontSize
		drawer := font.CreateDrawer(img, fontSize, h.FontPath)
		yAdjust := (fontSize / 2)

		for _, annotation := range h.Annotations.List {
			text := annotation.Text
			x := int(annotation.Position.X * float64(h.Width))
			y := int(annotation.Position.Y*float64(h.Height)) + yAdjust

			drawer.Dot = fixed.Point26_6{
				X: fixed.I(x) - (drawer.MeasureString(text) / 2),
				Y: fixed.I(y),
			}
			drawer.DrawString(text)
		}
	}
}
