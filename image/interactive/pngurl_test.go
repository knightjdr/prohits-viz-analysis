package interactive

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestPngurl(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create png.
	pngImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1, 1}})
	c := color.RGBA{uint8(0), uint8(0), uint8(0), 255}
	pngImage.Set(0, 0, c)
	myfile, _ := fs.Instance.Create("test.png")
	png.Encode(myfile, pngImage)

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(fs.Instance, "test.txt", []byte(""), 0644)

	// TEST1: convert png.
	want := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
	assert.Equal(t, want, Pngurl("test.png"), "PNG not converted correctly")

	// Mock OpenFile.
	file, _ := fs.Instance.Open("test.txt")
	fakeOpenFile := func(*afero.MemMapFs, string) (afero.File, error) {
		return file, errors.New("Test open error")
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(fs.Instance), "Open", fakeOpenFile)
	defer monkey.UnpatchInstanceMethod(reflect.TypeOf(fs.Instance), "Open")

	// TEST2: error opening file.
	want = "data:image/png;base64,"
	assert.Equal(t, want, Pngurl("test.png"), "Error opening file should return empty url")
}
