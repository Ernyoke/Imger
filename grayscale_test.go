package imger

import (
	"os"
	"fmt"
	"image/jpeg"
	"image"
	"image/draw"
	"testing"
)

func setupGrayscaleTestCase(t *testing.T) image.Image {
	imagePath := "res/girl.jpg"
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

func tearDownGrayscaleTestCase(t *testing.T, image image.Image, path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, image, nil)
}

func TestGrayScale(t *testing.T) {
	img := setupGrayscaleTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale(rgba)
	tearDownGrayscaleTestCase(t, gray, "res/grayscale/gray.jpg")
}

func TestGrayScale16(t *testing.T) {
	img := setupGrayscaleTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale16(rgba)
	tearDownGrayscaleTestCase(t, gray, "res/grayscale/gray16.jpg")
}


