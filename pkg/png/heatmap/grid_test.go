package heatmap

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Draw grid", func() {
	It("should draw a heatmap to a file", func() {
		img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{4, 4}})

		h := Initialize()
		h.AbundanceCap = 50
		h.CellSize = 2
		h.Height = 4
		h.MinAbundance = 0
		h.NumColors = 11
		h.Width = 4

		matrix := [][]float64{
			{0, 50},
			{50, 0},
		}

		expected := "iVBORw0KGgoAAAANSUhEUgAAAAQAAAAECAIAAAAmkwkpAAAAHElEQVR4nGL6DwYMYMDEgAQYIRREEkUGEAAA//9Y1Aj/wbgG8AAAAABJRU5ErkJggg=="

		drawGrid(img, h, matrix)
		buf := new(bytes.Buffer)
		png.Encode(buf, img)

		Expect(base64.StdEncoding.EncodeToString(buf.Bytes())).To(Equal(expected))
	})
})

var _ = Describe("Get gradient index", func() {
	It("should return a function for computing the slice index", func() {
		h := &Heatmap{
			AbundanceCap: 50,
			MinAbundance: 0,
			NumColors:    11,
		}

		testFunc := getGradientIndex(h)
		tests := []float64{-1, 60, 25, 10}

		expected := []int{0, 10, 5, 2}
		for i, test := range tests {
			Expect(testFunc(test)).To(Equal(expected[i]))
		}
	})
})
