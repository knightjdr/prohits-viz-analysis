package heatmap

import (
	"fmt"
	"image"
	"image/color"
	"image/png"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Minimap", func() {
	It("should parse minimap", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		pngImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1, 1}})
		c := color.RGBA{uint8(0), uint8(0), uint8(0), 255}
		pngImage.Set(0, 0, c)
		myfile, _ := fs.Instance.Create("test.png")
		png.Encode(myfile, pngImage)

		uri := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
		expected := fmt.Sprintf("\"minimap\": {\"main\":{\"image\":\"%s\"}}", uri)
		Expect(parseMinimap("test.png")).To(Equal(expected))
	})
})
