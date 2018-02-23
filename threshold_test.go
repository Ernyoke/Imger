package imger

import (
	"testing"
	"fmt"
	"os"
	"image"
	"image/jpeg"
	"image/draw"
)

func setupThresholdTestCase(t *testing.T) image.Image {
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

func tearDownThresholdTestCase(t *testing.T, image image.Image, path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, image, nil)
}

func TestThresholdBinray(t *testing.T) {
	img := setupThresholdTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale(rgba)
	thrsh, _ := Threshold(gray, 100, ThreshBinary)
	tearDownThresholdTestCase(t, thrsh, "res/threshold/threshBin.jpg")
}

func TestThresholdBinrayInv(t *testing.T) {
	img := setupThresholdTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale(rgba)
	thrsh, _ := Threshold(gray, 100, ThreshBinaryInv)
	tearDownThresholdTestCase(t, thrsh, "res/threshold/threshBinInv.jpg")
}

func TestThresholdTrunc(t *testing.T) {
	img := setupThresholdTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale(rgba)
	thrsh, _ := Threshold(gray, 100, ThreshTrunc)
	tearDownThresholdTestCase(t, thrsh, "res/threshold/threshTrunc.jpg")
}

func TestThresholdToZero(t *testing.T) {
	img := setupThresholdTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale(rgba)
	thrsh, _ := Threshold(gray, 100, ThreshToZero)
	tearDownThresholdTestCase(t, thrsh, "res/threshold/threshToZero.jpg")
}

func TestThresholdToZeroInv(t *testing.T) {
	img := setupThresholdTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale(rgba)
	thrsh, _ := Threshold(gray, 100, ThreshToZeroInv)
	tearDownThresholdTestCase(t, thrsh, "res/threshold/threshBin.jpg")
}

func TestThreshold16Bin(t *testing.T) {
	img := setupThresholdTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale16(rgba)
	thrsh, _ := Threshold16(gray, 32000, ThreshBinary)
	tearDownThresholdTestCase(t, thrsh, "res/threshold/thresh16Bin.jpg")
}

func TestThreshold16BinInv(t *testing.T) {
	img := setupThresholdTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale16(rgba)
	thrsh, _ := Threshold16(gray, 32000, ThreshBinaryInv)
	tearDownThresholdTestCase(t, thrsh, "res/threshold/thresh16BinInv.jpg")
}

func TestThreshold16Trunc(t *testing.T) {
	img := setupThresholdTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale16(rgba)
	thrsh, _ := Threshold16(gray, 32000, ThreshTrunc)
	tearDownThresholdTestCase(t, thrsh, "res/threshold/thresh16Trunc.jpg")
}

func TestThreshold16ToZero(t *testing.T) {
	img := setupThresholdTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale16(rgba)
	thrsh, _ := Threshold16(gray, 32000, ThreshToZero)
	tearDownThresholdTestCase(t, thrsh, "res/threshold/thresh16ToZero.jpg")
}

func TestThreshold16ToZeroInv(t *testing.T) {
	img := setupThresholdTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := Grayscale16(rgba)
	thrsh, _ := Threshold16(gray, 32000, ThreshToZeroInv)
	tearDownThresholdTestCase(t, thrsh, "res/threshold/thresh16ToZeroInv.jpg")
}