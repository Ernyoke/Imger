package edgedetection

import (
	"fmt"
	"github.com/ernyoke/imger/grayscale"
	"github.com/ernyoke/imger/padding"
	"image"
	"image/draw"
	"image/png"
	"os"
	"testing"
)

// -----------------------------Acceptance tests------------------------------------

func setupTestCase(t *testing.T) image.Image {
	imagePath := "../res/engine.png"
	file, err := os.Open(imagePath)
	defer file.Close()
	if err != nil {
		t.Log(os.Stderr, "%v\n", err)
	}
	img, err := png.Decode(file)
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
	png.Encode(f, image)
}

func Test_Acceptance_HorizontalSobelGray(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	sobel, _ := HorizontalSobelGray(gray, padding.BorderReflect)
	tearDownTestCase(t, sobel, "../res/sobel/horizontalSobelGray.png")
}

func Test_Acceptance_VerticalSobelGray(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	sobel, _ := VerticalSobelGray(gray, padding.BorderReflect)
	tearDownTestCase(t, sobel, "../res/sobel/verticalSobelGray.png")
}

func Test_Acceptance_SobelGray(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	sobel, _ := SobelGray(gray, padding.BorderReflect)
	tearDownTestCase(t, sobel, "../res/sobel/sobelGray.png")
}

func Test_Acceptance_SobelRGBA(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	sobel, _ := SobelRGBA(rgba, padding.BorderReflect)
	tearDownTestCase(t, sobel, "../res/sobel/sobelRGBA.png")
}

// ---------------------------------------------------------------------------------
