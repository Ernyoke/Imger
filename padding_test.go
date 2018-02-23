package imger

import (
	"os"
	"fmt"
	"image/jpeg"
	"image"
	"image/draw"
	"testing"
)

func setupPaddingTestCase(t *testing.T) image.Image {
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

func tearDownPaddingTestCase(t *testing.T, image image.Image, path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, image, nil)
}

func TestGrayPaddingBorderConstant(t *testing.T) {
	img := setupPaddingTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale(rgba)
	padded, _ := PaddingGray(gray, image.Point{15, 15}, image.Point{8, 8}, BorderConstant)
	tearDownPaddingTestCase(t, padded, "res/padding/grayPaddingBorderConstant.jpg")
}

func TestGrayPaddingBorderConstantDistortedAnchor(t *testing.T) {
	img := setupPaddingTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale(rgba)
	padded, _ := PaddingGray(gray, image.Point{50, 50}, image.Point{8, 8}, BorderConstant)
	tearDownPaddingTestCase(t, padded, "res/padding/grayPaddingBorderConstantDistortedAnchor.jpg")
}

func TestGrayPaddingBorderReplicate(t *testing.T) {
	img := setupPaddingTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale(rgba)
	padded, _ := PaddingGray(gray, image.Point{15, 15}, image.Point{8, 8}, BorderReplicate)
	tearDownPaddingTestCase(t, padded, "res/padding/grayPaddingBorderReplicate.jpg")
}

func TestGrayPaddingBorderReflect(t *testing.T) {
	img := setupPaddingTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale(rgba)
	padded, _ := PaddingGray(gray, image.Point{15, 15}, image.Point{8, 8}, BorderReflect)
	tearDownPaddingTestCase(t, padded, "res/padding/grayPaddingBorderReflect.jpg")
}

func TestRGBAPaddingBorderConstant(t *testing.T) {
	img := setupPaddingTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	padded, _ := PaddingRGBA(rgba, image.Point{15, 15}, image.Point{8, 8}, BorderConstant)
	tearDownPaddingTestCase(t, padded, "res/padding/rgbaPaddedBorderConstant.jpg")
}

func TestRGBAPaddingBorderReplicate(t *testing.T) {
	img := setupPaddingTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	padded, _ := PaddingRGBA(rgba, image.Point{15, 15}, image.Point{8, 8}, BorderReplicate)
	tearDownPaddingTestCase(t, padded, "res/padding/rgbaPaddedBorderReplicate.jpg")
}

func TestRGBAPaddingBorderReflect(t *testing.T) {
	img := setupPaddingTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	padded, _ := PaddingRGBA(rgba, image.Point{15, 15}, image.Point{8, 8}, BorderReflect)
	tearDownPaddingTestCase(t, padded, "res/padding/rgbaPaddedBorderReflect.jpg")
}
