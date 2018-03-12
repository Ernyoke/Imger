package edgedetection

import (
	"github.com/Ernyoke/Imger/imgio"
	"github.com/Ernyoke/Imger/padding"
	"image"
	"testing"
)

// -----------------------------Acceptance tests------------------------------------

func setupTestCaseGraySobel(t *testing.T) *image.Gray {
	path := "../res/engine.png"
	img, err := imgio.ImreadGray(path)
	if err != nil {
		t.Errorf("Could not read image from path: %s", path)
	}
	return img
}

func setupTestCaseRGBASobel(t *testing.T) *image.RGBA {
	path := "../res/engine.png"
	img, err := imgio.ImreadRGBA(path)
	if err != nil {
		t.Errorf("Could not read image from path: %s", path)
	}
	return img
}

func tearDownTestCaseSobel(t *testing.T, img image.Image, path string) {
	err := imgio.Imwrite(img, path)
	if err != nil {
		t.Errorf("Could not write image to path: %s", path)
	}
}

func Test_Acceptance_HorizontalSobelGray(t *testing.T) {
	gray := setupTestCaseGraySobel(t)
	sobel, _ := HorizontalSobelGray(gray, padding.BorderReflect)
	tearDownTestCaseSobel(t, sobel, "../res/edge/horizontalSobelGray.png")
}

func Test_Acceptance_VerticalSobelGray(t *testing.T) {
	gray := setupTestCaseGraySobel(t)
	sobel, _ := VerticalSobelGray(gray, padding.BorderReflect)
	tearDownTestCaseSobel(t, sobel, "../res/edge/verticalSobelGray.png")
}

func Test_Acceptance_SobelGray(t *testing.T) {
	gray := setupTestCaseGraySobel(t)
	sobel, _ := SobelGray(gray, padding.BorderReflect)
	tearDownTestCaseSobel(t, sobel, "../res/edge/sobelGray.png")
}

func Test_Acceptance_HorizontalSobelRGBA(t *testing.T) {
	rgba := setupTestCaseRGBASobel(t)
	sobel, _ := HorizontalSobelRGBA(rgba, padding.BorderReflect)
	tearDownTestCaseSobel(t, sobel, "../res/edge/horizontalSobelRGBA.png")
}

func Test_Acceptance_VerticalSobelRGBA(t *testing.T) {
	rgba := setupTestCaseRGBASobel(t)
	sobel, _ := VerticalSobelRGBA(rgba, padding.BorderReflect)
	tearDownTestCaseSobel(t, sobel, "../res/edge/verticalSobelRGBA.png")
}

func Test_Acceptance_SobelRGBA(t *testing.T) {
	rgba := setupTestCaseRGBASobel(t)
	sobel, _ := SobelRGBA(rgba, padding.BorderReflect)
	tearDownTestCaseSobel(t, sobel, "../res/edge/sobelRGBA.png")
}

// ---------------------------------------------------------------------------------
