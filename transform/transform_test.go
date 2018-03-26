package transform

import (
	"testing"
	"image"
	"github.com/Ernyoke/Imger/imgio"
)

// -----------------------------Acceptance tests------------------------------------
func setupTestCaseGray(t *testing.T) *image.Gray {
	path := "../res/girl.jpg"
	img, err := imgio.ImreadGray(path)
	if err != nil {
		t.Errorf("Could not read image from path: %s", path)
	}
	return img
}

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

func Test_Acceptance_RotateGray90(t *testing.T) {
	gray := setupTestCaseGray(t)
	actual, err:= RotateGray(gray, 90, image.Point{X: 201, Y: 201}, true)
	if err != nil {
		t.Errorf("Should not throw error!")
	}
	tearDownTestCase(t, actual, "../res/transform/roateGray90.jpg")
}

func Test_Acceptance_RotateGray45(t *testing.T) {
	gray := setupTestCaseGray(t)
	actual, err:= RotateGray(gray, 45, image.Point{X: 201, Y: 201}, true)
	if err != nil {
		t.Errorf("Should not throw error!")
	}
	tearDownTestCase(t, actual, "../res/transform/roateGray45.jpg")
}

func Test_Acceptance_RotateGray22(t *testing.T) {
	gray := setupTestCaseGray(t)
	actual, err:= RotateGray(gray, 22, image.Point{X: 201, Y: 201}, true)
	if err != nil {
		t.Errorf("Should not throw error!")
	}
	tearDownTestCase(t, actual, "../res/transform/roateGray22.jpg")
}

func Test_Acceptance_RotateRGBA90(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, err:= RotateRGBA(rgba, 90, image.Point{X: 201, Y: 201}, true)
	if err != nil {
		t.Errorf("Should not throw error!")
	}
	tearDownTestCase(t, actual, "../res/transform/roateRGBA90.jpg")
}

func Test_Acceptance_RotateRGBA45(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, err:= RotateRGBA(rgba, 45, image.Point{X: 201, Y: 201}, true)
	if err != nil {
		t.Errorf("Should not throw error!")
	}
	tearDownTestCase(t, actual, "../res/transform/roateRGBA45.jpg")
}

func Test_Acceptance_RotateRGBA22(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, err:= RotateRGBA(rgba, 22, image.Point{X: 201, Y: 201}, true)
	if err != nil {
		t.Errorf("Should not throw error!")
	}
	tearDownTestCase(t, actual, "../res/transform/roateRGBA22.jpg")
}

