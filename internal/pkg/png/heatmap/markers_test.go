package heatmap

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Draw markers", func() {
	It("should not add markers", func() {
		img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{4, 4}})

		h := Initialize()
		h.Markers = types.Markers{
			Color: "#000000",
			List:  map[string]types.Marker{},
		}
		h.Height = 4
		h.Width = 4

		expected := "iVBORw0KGgoAAAANSUhEUgAAAAQAAAAECAYAAACp8Z5+AAAAEklEQVR4nGJiQANkCAACAAD//wGYAAkg3ttpAAAAAElFTkSuQmCC"

		drawMarkers(img, h)
		actualBuf := new(bytes.Buffer)
		png.Encode(actualBuf, img)

		Expect(base64.StdEncoding.EncodeToString(actualBuf.Bytes())).To(Equal(expected))
	})

	It("should add markers", func() {
		img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{10, 10}})

		h := Initialize()
		h.CellSize = 1
		h.Markers = types.Markers{
			Color: "#000000",
			List: map[string]types.Marker{
				"a": types.Marker{Height: 3, Width: 4, X: 0.25, Y: 0.5},
			},
		}
		h.Height = 10
		h.Width = 10

		expected := "iVBORw0KGgoAAAANSUhEUgAAAAoAAAAKCAYAAACNMs+9AAAAJ0lEQVR4nGJiIBIMBYUsSOz/WOQZsSlEkUDXSJb" +
			"VuKwnDQACAAD//6ExAxg0kTKOAAAAAElFTkSuQmCC"

		drawMarkers(img, h)
		actualBuf := new(bytes.Buffer)
		png.Encode(actualBuf, img)

		Expect(base64.StdEncoding.EncodeToString(actualBuf.Bytes())).To(Equal(expected))
	})
})
