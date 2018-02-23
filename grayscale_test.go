package imger

import (
	"os"
	"fmt"
	"image/jpeg"
	"image"
	"image/draw"
	"testing"
)

func TestGrayScale(t *testing.T) {
	imagePath := "res/girl.jpg"

	file, err := os.Open(imagePath)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale(rgba)
	f, err := os.Create("res/gray.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, gray, nil)
}

func TestGrayScale16(t *testing.T) {
	imagePath := "res/girl.jpg"

	file, err := os.Open(imagePath)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale16(rgba)
	f, err := os.Create("res/gray16.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, gray, nil)
}


