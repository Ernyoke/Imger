package transform

import (
	"github.com/ernyoke/imger/imgio"
	"github.com/ernyoke/imger/utils"
	"image"
	"math"
	"testing"
)

// ---------------------------------Unit tests--------------------------------------
func Test_AngleToRadians0(t *testing.T) {
	expected := 0.0
	actual := angleToRadians(0)
	if !utils.IsEqualFloat64(actual, expected) {
		t.Errorf("Expected value %f is not eqaul with actual value %f", expected, actual)
	}
}

func Test_AngleToRadians30(t *testing.T) {
	expected := math.Pi / 6
	actual := angleToRadians(30)
	if !utils.IsEqualFloat64(actual, expected) {
		t.Errorf("Expected value %f is not eqaul with actual value %f", expected, actual)
	}
}

func Test_AngleToRadians45(t *testing.T) {
	expected := math.Pi / 4
	actual := angleToRadians(45)
	if !utils.IsEqualFloat64(actual, expected) {
		t.Errorf("Expected value %f is not eqaul with actual value %f", expected, actual)
	}
}

func Test_AngleToRadians60(t *testing.T) {
	expected := math.Pi / 3
	actual := angleToRadians(60)
	if !utils.IsEqualFloat64(actual, expected) {
		t.Errorf("Expected value %f is not eqaul with actual value %f", expected, actual)
	}
}

func Test_AngleToRadians90(t *testing.T) {
	expected := math.Pi / 2
	actual := angleToRadians(90)
	if !utils.IsEqualFloat64(actual, expected) {
		t.Errorf("Expected value %f is not eqaul with actual value %f", expected, actual)
	}
}

func Test_AngleToRadians180(t *testing.T) {
	expected := math.Pi
	actual := angleToRadians(180)
	if !utils.IsEqualFloat64(actual, expected) {
		t.Errorf("Expected value %f is not eqaul with actual value %f", expected, actual)
	}
}

func Test_AngleToRadians360(t *testing.T) {
	expected := math.Pi * 2
	actual := angleToRadians(360)
	if !utils.IsEqualFloat64(actual, expected) {
		t.Errorf("Expected value %f is not eqaul with actual value %f", expected, actual)
	}
}

func Test_AngleToRadiansNeg30(t *testing.T) {
	expected := -math.Pi / 6
	actual := angleToRadians(-30)
	if !utils.IsEqualFloat64(actual, expected) {
		t.Errorf("Expected value %f is not eqaul with actual value %f", expected, actual)
	}
}

func Test_AngleToRadiansNeg60(t *testing.T) {
	expected := -math.Pi / 3
	actual := angleToRadians(-60)
	if !utils.IsEqualFloat64(actual, expected) {
		t.Errorf("Expected value %f is not eqaul with actual value %f", expected, actual)
	}
}

func Test_ComputeFitSize90(t *testing.T) {
	initial := image.Point{X: 1024, Y: 768}
	angle := 90.0
	expected := image.Point{X: 768, Y: 1024}
	actual := computeFitSize(initial, angleToRadians(angle))
	if expected != actual {
		t.Errorf("Expected value %s is not eqaul with actual value %s", expected, actual)
	}
}

func Test_ComputeFitSize45(t *testing.T) {
	initial := image.Point{X: 1024, Y: 768}
	angle := 45.0
	expected := image.Point{X: 1267, Y: 1267}
	actual := computeFitSize(initial, angleToRadians(angle))
	if expected != actual {
		t.Errorf("Expected value %s is not eqaul with actual value %s", expected, actual)
	}
}

func Test_ComputeFitSize22_5(t *testing.T) {
	initial := image.Point{X: 1024, Y: 768}
	angle := 22.5
	expected := image.Point{X: 1239, Y: 1101}
	actual := computeFitSize(initial, angleToRadians(angle))
	if expected != actual {
		t.Errorf("Expected value %s is not eqaul with actual value %s", expected, actual)
	}
}

func Test_ComputeFitSizeNeg22_5(t *testing.T) {
	initial := image.Point{X: 1024, Y: 768}
	angle := -22.5
	expected := image.Point{X: 1239, Y: 1101}
	actual := computeFitSize(initial, angleToRadians(angle))
	if expected != actual {
		t.Errorf("Expected value %s is not eqaul with actual value %s", expected, actual)
	}
}

func Test_ComputeFitSizeNeg45(t *testing.T) {
	initial := image.Point{X: 1024, Y: 768}
	angle := -45.0
	expected := image.Point{X: 1267, Y: 1267}
	actual := computeFitSize(initial, angleToRadians(angle))
	if expected != actual {
		t.Errorf("Expected value %s is not eqaul with actual value %s", expected, actual)
	}
}

func Test_ComputeOffset90(t *testing.T) {
	size := image.Point{X: 1024, Y: 768}
	radians := angleToRadians(90)
	expected := image.Point{X: -128, Y: 128}
	actual := computeOffset(size, computeFitSize(size, radians))
	if expected != actual {
		t.Errorf("Expected value %s is not eqaul with actual value %s", expected, actual)
	}
}

func Test_ComputeOffset180(t *testing.T) {
	size := image.Point{X: 1024, Y: 768}
	radians := angleToRadians(180)
	expected := image.Point{X: 0, Y: 0}
	actual := computeOffset(size, computeFitSize(size, radians))
	if expected != actual {
		t.Errorf("Expected value %s is not eqaul with actual value %s", expected, actual)
	}
}

func Test_GetOriginalPixelPosition90(t *testing.T) {
	x := 0
	y := 0
	size := image.Point{X: 1024, Y: 768}
	radians := angleToRadians(90)
	anchor := image.Point{X: 512, Y: 384}
	offset := computeOffset(size, computeFitSize(size, radians))
	expectedX := 1024
	expectedY := 0
	actualX, actualY := getOriginalPixelPosition(x, y, radians, anchor, offset)
	if actualX != expectedX && actualY != expectedY {
		t.Errorf("Expected value [%d %d] is not eqaul with actual value [%d %d]", expectedX, expectedY, actualX, actualY)
	}
}

func Test_GetOriginalPixelPosition45(t *testing.T) {
	x := 0
	y := 0
	size := image.Point{X: 1024, Y: 768}
	radians := angleToRadians(45)
	anchor := image.Point{X: 512, Y: 384}
	offset := computeOffset(size, computeFitSize(size, radians))
	expectedX := 512
	expectedY := -512
	actualX, actualY := getOriginalPixelPosition(x, y, radians, anchor, offset)
	if actualX != expectedX && actualY != expectedY {
		t.Errorf("Expected value [%d %d] is not eqaul with actual value [%d %d]", expectedX, expectedY, actualX, actualY)
	}
}

// -----------------------------Acceptance tests------------------------------------
func setupTestCaseGray(t *testing.T) *image.Gray {
	path := "../res/building.jpg"
	img, err := imgio.ImreadGray(path)
	if err != nil {
		t.Errorf("Could not read image from path: %s", path)
	}
	return img
}

func setupTestCaseRGBA(t *testing.T) *image.RGBA {
	path := "../res/building.jpg"
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
	actual, err := RotateGray(gray, 90, image.Point{X: 512, Y: 384}, true)
	if err != nil {
		t.Errorf("Should not throw error!")
	}
	tearDownTestCase(t, actual, "../res/transform/roateGray90.jpg")
}

func Test_Acceptance_RotateGray45(t *testing.T) {
	gray := setupTestCaseGray(t)
	actual, err := RotateGray(gray, 45, image.Point{X: 512, Y: 384}, true)
	if err != nil {
		t.Errorf("Should not throw error!")
	}
	tearDownTestCase(t, actual, "../res/transform/roateGray45.jpg")
}

func Test_Acceptance_RotateGray22(t *testing.T) {
	gray := setupTestCaseGray(t)
	actual, err := RotateGray(gray, 22, image.Point{X: 512, Y: 384}, true)
	if err != nil {
		t.Errorf("Should not throw error!")
	}
	tearDownTestCase(t, actual, "../res/transform/roateGray22.jpg")
}

func Test_Acceptance_RotateRGBA90(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, err := RotateRGBA(rgba, 90, image.Point{X: 512, Y: 384}, true)
	if err != nil {
		t.Errorf("Should not throw error!")
	}
	tearDownTestCase(t, actual, "../res/transform/roateRGBA90.jpg")
}

func Test_Acceptance_RotateRGBA45(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, err := RotateRGBA(rgba, 45, image.Point{X: 512, Y: 384}, true)
	if err != nil {
		t.Errorf("Should not throw error!")
	}
	tearDownTestCase(t, actual, "../res/transform/roateRGBA45.jpg")
}

func Test_Acceptance_RotateRGBA22(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, err := RotateRGBA(rgba, 22, image.Point{X: 512, Y: 384}, true)
	if err != nil {
		t.Errorf("Should not throw error!")
	}
	tearDownTestCase(t, actual, "../res/transform/roateRGBA22.jpg")
}

func Test_Acceptance_RotateRGBA45Noresize(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, err := RotateRGBA(rgba, 45, image.Point{X: 512, Y: 384}, false)
	if err != nil {
		t.Errorf("Should not throw error!")
	}
	tearDownTestCase(t, actual, "../res/transform/roateRGBA45Noresize.jpg")
}
