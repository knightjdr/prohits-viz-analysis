package uri_test

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"reflect"

	"github.com/bouk/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/png/uri"
)

var _ = Describe("Convert png to uri", func() {
	It("should convert a png file to a uri", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		pngImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1, 1}})
		c := color.RGBA{uint8(0), uint8(0), uint8(0), 255}
		pngImage.Set(0, 0, c)
		myfile, _ := fs.Instance.Create("test.png")
		png.Encode(myfile, pngImage)

		expected := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
		Expect(Convert("test.png")).To(Equal(expected))
	})

	It("should return empty uri when failing to open file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test.txt", []byte(""), 0644)

		file, _ := fs.Instance.Open("test.txt")
		fakeOpenFile := func(*afero.MemMapFs, string) (afero.File, error) {
			return file, errors.New("Test open error")
		}
		monkey.PatchInstanceMethod(reflect.TypeOf(fs.Instance), "Open", fakeOpenFile)
		defer monkey.UnpatchInstanceMethod(reflect.TypeOf(fs.Instance), "Open")

		expected := "data:image/png;base64,"
		Expect(Convert("test.png")).To(Equal(expected))
	})
})

/* func TestPngurl(t *testing.T) {
	// Mock filesystem.


	// Create png.


	// Create test directory and files.


	// TEST1: convert png.

	assert.Equal(t, want, Pngurl("test.png"), "PNG not converted correctly")

	// Mock OpenFile.


	// TEST2: error opening file.

} */
