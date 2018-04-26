package imageproc_test

import (
	"image"
	"image/color"
	"image/png"
	"path/filepath"
	"testing"

	"github.com/translucens/gopher-dojo/imgconv/imageproc"
	"github.com/translucens/gopher-dojo/imgconv/options"
)

const Separator = string(filepath.Separator)
const TestdataDir = ".." + Separator + "testdata" + Separator

// TestHelper: check image width, height
func imageWidthHeightIs(t *testing.T, img image.Image, width int, height int) {
	t.Helper()

	size := img.Bounds().Size()
	t.Logf("Image size is %d x %d\n", size.X, size.Y)
	if size.X != width || size.Y != height {
		t.Fatalf("Image size is not match with expected %d x %d", width, height)
	}
}

// TestHelper: check top-left corner image color
func cornerColorIs(t *testing.T, img image.Image, expected color.Color) {
	t.Helper()

	colorIs(t, img, 0, 0, expected)
}

// TestHelper: check top-left corner image color
func colorIs(t *testing.T, img image.Image, x int, y int, expected color.Color) {
	t.Helper()

	eR, eG, eB, _ := expected.RGBA()
	r, g, b, _ := img.At(x, y).RGBA()

	t.Logf("color @(%d, %d) (R,G,B) is %d, %d, %d\n", x, y, r, g, b)
	if r != eR || g != eG || b != eB {
		t.Fatalf("color is not match with expected %d, %d, %d", eR, eG, eB)
	}
}

type TestImageProperty struct {
	filepath    string
	filetype    string
	width       int
	height      int
	cornercolor color.Color
}

func TestLoadImages(t *testing.T) {
	cases := []TestImageProperty{
		{filepath: TestdataDir + "black16x32px.jpg", filetype: "jpg", width: 16, height: 32, cornercolor: color.NRGBA{0, 0, 0, 0}},
		{filepath: TestdataDir + "black16px-prograssive.jpeg", filetype: "jpg", width: 16, height: 16, cornercolor: color.NRGBA{0, 0, 0, 0}},
		{filepath: TestdataDir + "black16px.png", filetype: "png", width: 16, height: 16, cornercolor: color.NRGBA{0, 0, 0, 0}},
		{filepath: TestdataDir + "white32px.JPG", filetype: "jpg", width: 32, height: 32, cornercolor: color.NRGBA{255, 255, 255, 255}},
		{filepath: TestdataDir + "childdir" + Separator + "white32px.png", filetype: "png", width: 32, height: 32, cornercolor: color.NRGBA{255, 255, 255, 255}},
		{filepath: TestdataDir + Separator + "alphach.png", filetype: "png", width: 16, height: 16, cornercolor: color.NRGBA{255, 0, 0, 255}}}

	for _, c := range cases {
		img, err := imageproc.OpenImageFile(c.filepath, c.filetype)
		if err != nil {
			t.Fatalf("Open file error: %s", err.Error())
		}
		imageWidthHeightIs(t, img, c.width, c.height)
		cornerColorIs(t, img, c.cornercolor)
	}
}

func TestLoadAlphaChPNGFile(t *testing.T) {
	img, err := imageproc.OpenImageFile(TestdataDir+"alphach.png", "png")
	if err != nil {
		t.Fatalf("Open file error: %s", err.Error())
	}
	imageWidthHeightIs(t, img, 16, 16)
	colorIs(t, img, 15, 0, color.NRGBA{0, 255, 0, 255})
	colorIs(t, img, 0, 15, color.NRGBA{0, 0, 255, 255})
	colorIs(t, img, 15, 15, color.NRGBA{0, 0, 0, 0})
}

func TestBrokenFile(t *testing.T) {
	_, err := imageproc.OpenImageFile(TestdataDir+"invalid.jpg", "jpg")
	if err == nil {
		t.Fatalf("This file should be broken!")
	}
	t.Log(err.Error())
}

func TestConvertJPEG2PNG(t *testing.T) {
	if err := imageproc.ConvertImage(TestdataDir+"black16x32px.jpg", options.Options{Overwrite: true, PNGLevel: int(png.DefaultCompression), InputFileType: "jpg", OutputFileType: "png"}); err != nil {
		t.Fatalf("convert failed %s", err.Error())
	}

	// skip if destination file already exists
	if err := imageproc.ConvertImage(TestdataDir+"black16x32px.jpg", options.Options{Overwrite: false, PNGLevel: int(png.DefaultCompression), InputFileType: "jpg", OutputFileType: "png"}); err != nil {
		t.Fatalf("convert failed %s", err.Error())
	}
}

func TestConvertPNG2JPEG(t *testing.T) {
	if err := imageproc.ConvertImages([]string{TestdataDir + "alphach.png"}, options.Options{Overwrite: true, JPEGQuality: 100, InputFileType: "png", OutputFileType: "jpg"}); err != nil {
		t.Fatalf("convert failed %s", err.Error())
	}
}
