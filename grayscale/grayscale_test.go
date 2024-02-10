package grayscale

import (
	"image"
	"testing"

	"github.com/ernyoke/imger/imgio"
)

// -----------------------------Acceptance tests------------------------------------
func setupTestCaseRGBA(t *testing.T) *image.RGBA {
	path := "../res/girl.jpg"
	img, err := imgio.ImreadRGBA(path)
	if err != nil {
		t.Errorf("Could not read image from path: %s", path)
	}
	return img
}

func tearDownTestCase(t *testing.T, img image.Image, path string) {
	err := imgio.Imwrite(img, path)
	if err != nil {
		t.Errorf("Could not write image to path: %s", path)
	}
}

func Test_Acceptance_GrayScale(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	gray := Grayscale(rgba)
	tearDownTestCase(t, gray, "../res/grayscale/gray.jpg")
}

func Test_Acceptance_GrayScaleCropped(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	cropped := rgba.SubImage(image.Rect(100, 100, rgba.Bounds().Max.X-100, rgba.Bounds().Max.Y-100))
	gray := Grayscale(cropped)
	tearDownTestCase(t, gray, "../res/grayscale/cropped_gray.jpg")
}

func Test_Acceptance_GrayScale16(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	gray := Grayscale16(rgba)
	tearDownTestCase(t, gray, "../res/grayscale/gray16.jpg")
}

func Test_Acceptance_GrayScale16Cropped(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	cropped := rgba.SubImage(image.Rect(100, 100, rgba.Bounds().Max.X-100, rgba.Bounds().Max.Y-100))
	gray := Grayscale16(cropped)
	tearDownTestCase(t, gray, "../res/grayscale/cropped_gray16.jpg")
}

// ---------------------------------------------------------------------------------
