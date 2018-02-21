package imger

import (
	"testing"
	//"github.com/stretchr/testify/assert"
	"fmt"
	"os"
	"image"
	"image/jpeg"
	"image/draw"
)

func TestThresholdBinray(t *testing.T) {
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
	thrsh, _ := Threshold(gray, 100, ThreshBinary)
	f, err := os.Create("res/threshBin.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, thrsh, nil)

}

func TestThresholdBinrayInv(t *testing.T) {
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
	thrsh, _ := Threshold(gray, 100, ThreshBinaryInv)
	f, err := os.Create("res/threshBinInv.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, thrsh, nil)

}

func TestThresholdTrunc(t *testing.T) {
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
	thrsh, _ := Threshold(gray, 100, ThreshTrunc)
	f, err := os.Create("res/threshTrunc.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, thrsh, nil)

}

func TestThresholdToZero(t *testing.T) {
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
	thrsh, _ := Threshold(gray, 100, ThreshToZero)
	f, err := os.Create("res/threshToZero.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, thrsh, nil)

}

func TestThresholdToZeroInv(t *testing.T) {
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
	thrsh, _ := Threshold(gray, 100, ThreshToZeroInv)
	f, err := os.Create("res/threshToZeroInv.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, thrsh, nil)

}

func TestThreshold16Bin(t *testing.T) {
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
	thrsh, _ := Threshold16(gray, 32000, ThreshBinary)
	f, err := os.Create("res/thresh16Bin.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, thrsh, nil)

}

func TestThreshold16BinInv(t *testing.T) {
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
	thrsh, _ := Threshold16(gray, 32000, ThreshBinaryInv)
	f, err := os.Create("res/thresh16BinInv.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, thrsh, nil)

}

func TestThreshold16Trunc(t *testing.T) {
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
	thrsh, _ := Threshold16(gray, 32000, ThreshTrunc)
	f, err := os.Create("res/thresh16Trunc.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, thrsh, nil)

}

func TestThreshold16ToZero(t *testing.T) {
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
	thrsh, _ := Threshold16(gray, 32000, ThreshToZero)
	f, err := os.Create("res/thresh16ToZero.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, thrsh, nil)

}

func TestThreshold16ToZeroInv(t *testing.T) {
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
	thrsh, _ := Threshold16(gray, 32000, ThreshToZeroInv)
	f, err := os.Create("res/thresh16ToZeroInv.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, thrsh, nil)

}