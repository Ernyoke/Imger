package imger

import (
	"os"
	"fmt"
	"image/jpeg"
	"image"
	"image/draw"
	"testing"
)

func TestGrayPaddingBorderConstant(t *testing.T) {
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
	padded, _ := PaddingGray(gray, image.Point{15, 15}, image.Point{8, 8}, BorderConstant)
	f, err := os.Create("res/grayPaddedBorderConstant.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, padded, nil)
}

func TestGrayPaddingBorderConstantDistortedAnchor(t *testing.T) {
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
	padded, _ := PaddingGray(gray, image.Point{50, 50}, image.Point{8, 8}, BorderConstant)
	f, err := os.Create("res/grayPaddedBorderConstantDistorted.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, padded, nil)
}

func TestGrayPaddingBorderReplicate(t *testing.T) {
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
	padded, _ := PaddingGray(gray, image.Point{15, 15}, image.Point{8, 8}, BorderReplicate)
	f, err := os.Create("res/grayPaddedBorderReplicate.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, padded, nil)
}

func TestGrayPaddingBorderReflect(t *testing.T) {
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
	padded, _ := PaddingGray(gray, image.Point{15, 15}, image.Point{8, 8}, BorderReflect)
	f, err := os.Create("res/grayPaddedBorderReflect.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, padded, nil)
}

func TestRGBAPaddingBorderConstant(t *testing.T) {
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
	padded, _ := PaddingRGBA(rgba, image.Point{15, 15}, image.Point{8, 8}, BorderConstant)
	f, err := os.Create("res/rgbaPaddedBorderConstant.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, padded, nil)
}

func TestRGBAPaddingBorderReplicate(t *testing.T) {
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
	padded, _ := PaddingRGBA(rgba, image.Point{15, 15}, image.Point{8, 8}, BorderReplicate)
	f, err := os.Create("res/rgbaPaddedBorderReplicate.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, padded, nil)
}

func TestRGBAPaddingBorderReflect(t *testing.T) {
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
	padded, _ := PaddingRGBA(rgba, image.Point{15, 15}, image.Point{8, 8}, BorderReflect)
	f, err := os.Create("res/rgbaPaddedBorderReflect.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, padded, nil)
}
