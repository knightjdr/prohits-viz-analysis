package heatmap

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Draw annotations", func() {
	It("should not add annotations", func() {
		img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{4, 4}})

		h := Initialize()
		h.Annotations = types.Annotations{
			FontSize: 16,
			List:     map[string]types.Annotation{},
		}
		h.Height = 4
		h.Width = 4

		expected := "iVBORw0KGgoAAAANSUhEUgAAAAQAAAAECAYAAACp8Z5+AAAAEklEQVR4nGJiQANkCAACAAD//wGYAAkg3ttpAAAAAElFTkSuQmCC"

		drawAnnotations(img, h)
		actualBuf := new(bytes.Buffer)
		png.Encode(actualBuf, img)

		Expect(base64.StdEncoding.EncodeToString(actualBuf.Bytes())).To(Equal(expected))
	})

	It("should add annotations", func() {
		img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{40, 40}})

		h := Initialize()
		h.Annotations = types.Annotations{
			FontSize: 16,
			List: map[string]types.Annotation{
				"a": {Text: "a", Position: types.AnnotationPosition{X: 0.5, Y: 0.25}},
				"b": {Text: "b", Position: types.AnnotationPosition{X: 0.5, Y: 0.75}},
			},
		}
		h.Height = 40
		h.Width = 40

		expected := "iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAYAAACM/rhtAAAAiUlEQVR4nOyXMQ7AIAwDoer/v0z3KFNtSxbxb" +
			"Szo5CZpeJY5EUSJIEoEUSKIEkEUe8GXeNcp5028W0IV/oUyQSs6OYqwqkloabI+8S5S9Xwv9nMwgigRRIkgSgRRxgjK/rtjEp" +
			"TB3KhXs3LZcJo6tNuoJYnZ1yBTUDJqmE3SvUvux74GvwAAAP//DYUPSX+vhOYAAAAASUVORK5CYII="

		drawAnnotations(img, h)
		actualBuf := new(bytes.Buffer)
		png.Encode(actualBuf, img)

		Expect(base64.StdEncoding.EncodeToString(actualBuf.Bytes())).To(Equal(expected))
	})
})
