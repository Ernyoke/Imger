package threshold

import (
	"fmt"
	"github.com/ernyoke/imgur/grayscale"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"testing"
)

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

func Test_Acceptance_ThresholdBinray(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	thrsh, _ := Threshold(gray, 100, ThreshBinary)
	tearDownTestCase(t, thrsh, "../res/threshold/threshBin.jpg")
}

func Test_Acceptance_ThresholdBinrayInv(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	thrsh, _ := Threshold(gray, 100, ThreshBinaryInv)
	tearDownTestCase(t, thrsh, "../res/threshold/threshBinInv.jpg")
}

func Test_Acceptance_ThresholdTrunc(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	thrsh, _ := Threshold(gray, 100, ThreshTrunc)
	tearDownTestCase(t, thrsh, "../res/threshold/threshTrunc.jpg")
}

func Test_Acceptance_ThresholdToZero(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	thrsh, _ := Threshold(gray, 100, ThreshToZero)
	tearDownTestCase(t, thrsh, "../res/threshold/threshToZero.jpg")
}

func Test_Acceptance_ThresholdToZeroInv(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	thrsh, _ := Threshold(gray, 100, ThreshToZeroInv)
	tearDownTestCase(t, thrsh, "../res/threshold/threshBin.jpg")
}

func Test_Acceptance_Threshold16Bin(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale16(rgba)
	thrsh, _ := Threshold16(gray, 32000, ThreshBinary)
	tearDownTestCase(t, thrsh, "../res/threshold/thresh16Bin.jpg")
}

func Test_Acceptance_Threshold16BinInv(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale16(rgba)
	thrsh, _ := Threshold16(gray, 32000, ThreshBinaryInv)
	tearDownTestCase(t, thrsh, "../res/threshold/thresh16BinInv.jpg")
}

func Test_Acceptance_Threshold16Trunc(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale16(rgba)
	thrsh, _ := Threshold16(gray, 32000, ThreshTrunc)
	tearDownTestCase(t, thrsh, "../res/threshold/thresh16Trunc.jpg")
}

func Test_Acceptance_Threshold16ToZero(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale16(rgba)
	thrsh, _ := Threshold16(gray, 32000, ThreshToZero)
	tearDownTestCase(t, thrsh, "../res/threshold/thresh16ToZero.jpg")
}

func Test_Acceptance_Threshold16ToZeroInv(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale16(rgba)
	thrsh, _ := Threshold16(gray, 32000, ThreshToZeroInv)
	tearDownTestCase(t, thrsh, "../res/threshold/thresh16ToZeroInv.jpg")
}

//---------------------------------------------------------------------------------
