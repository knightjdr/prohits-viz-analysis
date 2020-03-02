package color_test

import (
	"errors"
	"fmt"
	"image/color"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/color"
)

var _ = Describe("Convert hex to RGB", func() {
	It("should return an error for invalid format", func() {
		tests := []string{"000000", "#0000"}

		err := errors.New("invalid format")
		for _, test := range tests {
			_, actualErr := ConvertHexToRGB(test)
			Expect(actualErr).To(Equal(err), fmt.Sprintf("should return error for invalid color: %s", test))
		}
	})

	It("should convert colors", func() {
		tests := []string{"#0040ff", "#3366ff", "#ff3333"}

		expected := []color.RGBA{
			color.RGBA{A: 0xff, R: 0, G: 64, B: 255},
			color.RGBA{A: 0xff, R: 51, G: 102, B: 255},
			color.RGBA{A: 0xff, R: 255, G: 51, B: 51},
		}

		for i, test := range tests {
			actualRGB, _ := ConvertHexToRGB(test)
			Expect(actualRGB).To(Equal(expected[i]), fmt.Sprintf("should convert %s to %d,%d,%d", test, expected[i].R, expected[i].G, expected[i].B))
		}
	})
})
