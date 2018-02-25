package blur

import (
	"fmt"
	"github.com/ernyoke/imgur/grayscale"
	"github.com/ernyoke/imgur/padding"
	"github.com/ernyoke/imgur/utils"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"testing"
)

// ---------------------------------Unit tests------------------------------------
func TestGrayGaussianBlurZeroRadius(t *testing.T) {
	input := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
			0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
		},
	}
	_, error := GaussianBlurRGBA(&input, 0, 6, padding.BorderReflect)
	if error != nil {
		//ok
	} else {
		t.Fatal("no error thrown")
	}
}

func TestGrayGaussianBlurOneRadius(t *testing.T) {
	input := image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0x80, 0x56,
			0x56, 0x80, 0x69,
			0xEE, 0x29, 0xBB,
		},
	}
	expected := &image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0x47, 0x5A, 0x33,
			0x64, 0x88, 0x4D,
			0x3B, 0x59, 0x36,
		},
	}
	result, error := GaussianBlurGray(&input, 1, 2, padding.BorderConstant)
	if error != nil {
		t.Fatal(error)
	}
	utils.CompareGrayImagesWithOffset(t, expected, result, 1)
}

func TestRGBAGaussianBlurOneRadius(t *testing.T) {
	input := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0x56, 0x56, 0x56, 0xFF,
			0x56, 0x56, 0x56, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0x69, 0x69, 0x69, 0xFF,
			0xEE, 0xEE, 0xEE, 0xFF, 0x29, 0x29, 0x29, 0xFF, 0xBB, 0xBB, 0xBB, 0xFF,
		},
	}
	expected := &image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0x47, 0x47, 0x47, 0xFF, 0x5A, 0x5A, 0x5A, 0xFF, 0x33, 0x33, 0x33, 0xFF,
			0x64, 0x64, 0x64, 0xFF, 0x88, 0x88, 0x88, 0xFF, 0x4D, 0x4D, 0x4D, 0xFF,
			0x3B, 0x3B, 0x3B, 0xFF, 0x59, 0x59, 0x59, 0xFF, 0x36, 0x36, 0x36, 0xFF,
		},
	}
	actual, error := GaussianBlurRGBA(&input, 1, 2, padding.BorderConstant)
	if error != nil {
		t.Fatal(error)
	}
	//utils.PrintRGBA(t, actual)
	utils.CompareRGBAImagesWithOffset(t, expected, actual, 1)
}

// ---------------------------------------------------------------------------------

// -----------------------------Acceptance tests------------------------------------
func setupTestCase(t *testing.T) image.Image {
	imagePath := "../res/girl.jpg"
	file, err := os.Open(imagePath)
	defer file.Close()
	if err != nil {
		t.Log(os.Stderr, "%v\n", err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return img
}

func tearDownTestCase(t *testing.T, image image.Image, path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, image, nil)
}

func Test_Acceptance_GrayBlurInt(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	blured, _ := BoxGray(gray, image.Point{15, 15}, image.Point{8, 8}, padding.BorderReflect)
	tearDownTestCase(t, blured, "../res/blur/grayBlur.jpg")
}

func Test_Acceptance_RGBABlurInt(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	blured, _ := BoxRGBA(rgba, image.Point{15, 15}, image.Point{8, 8}, padding.BorderReflect)
	tearDownTestCase(t, blured, "../res/blur/rgbaBlur.jpg")
}

func Test_Acceptance_GrayGaussianBlurInt(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	blured, _ := GaussianBlurGray(gray, 7, 6, padding.BorderReflect)
	tearDownTestCase(t, blured, "../res/blur/grayGaussianBlur.jpg")
}

func Test_Acceptance_RGBAGaussianBlurInt(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	blured, _ := GaussianBlurRGBA(rgba, 5, 500, padding.BorderReflect)
	tearDownTestCase(t, blured, "../res/blur/rgbaGaussianBlur.jpg")
}

// ----------------------------------------------------------------------------------
