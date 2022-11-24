package edgedetection

import (
	"github.com/ernyoke/imger/imgio"
	"image"
	"testing"
)

// -----------------------------Acceptance tests------------------------------------
func setupTestCaseGray(t *testing.T) *image.Gray {
	path := "../res/engine.png"
	img, err := imgio.ImreadGray(path)
	if err != nil {
		t.Errorf("Could not read image from path: %s", path)
	}
	return img
}

func setupTestCaseRGBA(t *testing.T) *image.RGBA {
	path := "../res/engine.png"
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

func Test_Acceptance_CannyGray(t *testing.T) {
	gray := setupTestCaseGray(t)
	cny, err := CannyGray(gray, 15, 45, 5)
	if err != nil {
		t.Fatalf("Should not reach this point!")
	}
	tearDownTestCase(t, cny, "../res/edge/cannygray.jpg")
}

func Test_Acceptance_CannyRGBA(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	cny, err := CannyRGBA(rgba, 15, 45, 5)
	if err != nil {
		t.Fatalf("Should not reach this point!")
	}
	tearDownTestCase(t, cny, "../res/edge/cannyrgba.jpg")
}
