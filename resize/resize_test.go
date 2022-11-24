package resize

import (
	"github.com/ernyoke/imger/imgio"
	"image"
	"testing"
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

func Test_Acceptance_GrayResize_NN_2X(t *testing.T) {
	fxy := 2.0
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, fxy, fxy, InterNearest)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize2x.jpg")
}

func Test_Acceptance_GrayResize_NN_1_5X(t *testing.T) {
	fxy := 1.5
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, fxy, fxy, InterNearest)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize1_5x.jpg")
}

func Test_Acceptance_GrayResize_NN_2_5_AND_3_5X(t *testing.T) {
	fx := 2.5
	fy := 3.5
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, fx, fy, InterNearest)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fx), Y: int(float64(originalSize.Y) * fy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize2_5_and_3_5x.jpg")
}

func Test_Acceptance_GrayResize_NN_0_5X(t *testing.T) {
	fxy := 0.5
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, fxy, fxy, InterNearest)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize0_5x.jpg")
}

func Test_Acceptance_GrayResize_Linear_2X(t *testing.T) {
	fxy := 2.0
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, fxy, fxy, InterLinear)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize_Linear_2x.jpg")
}

func Test_Acceptance_GrayResize_CatmullRom_2X(t *testing.T) {
	fxy := 2.0
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, fxy, fxy, InterCatmullRom)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize_CatmullRom_2x.jpg")
}

func Test_Acceptance_GrayResize_CatmullRom_0_5X(t *testing.T) {
	fxy := 0.5
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, fxy, fxy, InterCatmullRom)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize_CatmullRom_0_5x.jpg")
}

func Test_Acceptance_GrayResize_Lanczos_2X(t *testing.T) {
	fxy := 2.0
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, fxy, fxy, InterLanczos)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize_Lanczos_2x.jpg")
}

func Test_Acceptance_GrayResize_Lanczos_0_5X(t *testing.T) {
	fxy := 0.5
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, fxy, fxy, InterLanczos)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize_Lanczos_0_5x.jpg")
}

func Test_Acceptance_RGBAResize_NN_2X(t *testing.T) {
	fxy := 2.0
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, fxy, fxy, InterNearest)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize2x.jpg")
}

func Test_Acceptance_RGBAResize_NN_1_5X(t *testing.T) {
	fxy := 1.5
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, fxy, fxy, InterNearest)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize1_5x.jpg")
}

func Test_Acceptance_RGBAResize_NN_2_5_AND_3_5X(t *testing.T) {
	fx := 2.5
	fy := 3.5
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, fx, fy, InterNearest)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fx), Y: int(float64(originalSize.Y) * fy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize2_5_and_3_5x.jpg")
}

func Test_Acceptance_RGBAResize_NN_0_5X(t *testing.T) {
	fxy := 0.5
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, fxy, fxy, InterNearest)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize0_5x.jpg")
}

func Test_Acceptance_RGBAResize_Linear_2X(t *testing.T) {
	fxy := 2.0
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, fxy, fxy, InterLinear)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize_Linear_2x.jpg")
}

func Test_Acceptance_RGBAResize_CatmullRom_2X(t *testing.T) {
	fxy := 2.0
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, fxy, fxy, InterCatmullRom)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize_CatmullRom_2x.jpg")
}

func Test_Acceptance_RGBAResize_CatmullRom_0_5X(t *testing.T) {
	fxy := 0.5
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, fxy, fxy, InterCatmullRom)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize_CatmullRom_0_5x.jpg")
}

func Test_Acceptance_RGBAResize_Lanczos_2X(t *testing.T) {
	fxy := 2.0
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, fxy, fxy, InterLanczos)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize_Lanczos_2x.jpg")
}

func Test_Acceptance_RGBAResize_Lanczos_0_5X(t *testing.T) {
	fxy := 0.5
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, fxy, fxy, InterLanczos)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * fxy), Y: int(float64(originalSize.Y) * fxy)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize_Lanczos_0_5x.jpg")
}

// ----------------------------------------------------------------------------------
