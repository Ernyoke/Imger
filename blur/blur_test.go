package blur

import (
	"image"
	"os"
	"image/jpeg"
	"fmt"
	"testing"
	"image/draw"
	"github.com/ernyoke/imgur/grayscale"
)

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

func TestGrayBlur(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	blured, _ := BlurGray(gray, image.Point{15, 15}, image.Point{8, 8})
	tearDownTestCase(t, blured, "../res/blur/grayBlur.jpg")
}

func TestRGBABlur(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	blured, _ := BlurRGBA(rgba, image.Point{15, 15}, image.Point{8, 8})
	tearDownTestCase(t, blured, "../res/blur/rgbaBlur.jpg")
}
