package resize

import (
	"image"
	"testing"
	"github.com/ernyoke/imger/imgio"
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
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, 2, 2, InterNearest)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: originalSize.X * 2, Y: originalSize.Y * 2}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize2x.jpg")
}

func Test_Acceptance_GrayResize_NN_1_5X(t *testing.T) {
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, 1.5, 1.5, InterNearest)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * 1.5), Y: int(float64(originalSize.Y) * 1.5)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize1_5x.jpg")
}

func Test_Acceptance_GrayResize_NN_2_5_AND_3_5X(t *testing.T) {
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, 2.5, 3.5, InterNearest)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * 2.5), Y: int(float64(originalSize.Y) * 3.5)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize2_5_and_3_5x.jpg")
}

func Test_Acceptance_GrayResize_NN_0_5X(t *testing.T) {
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, 0.5, 0.5, InterNearest)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * 0.5), Y: int(float64(originalSize.Y) * 0.5)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize0_5x.jpg")
}

func Test_Acceptance_GrayResize_Linear_2X(t *testing.T) {
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, 2, 2, InterLinear)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: originalSize.X * 2, Y: originalSize.Y * 2}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize_Linear_2x.jpg")
}

func Test_Acceptance_GrayResize_CatmullRom_2X(t *testing.T) {
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, 2, 2, InterCatmullRom)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: originalSize.X * 2, Y: originalSize.Y * 2}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize_CatmullRom_2x.jpg")
}

func Test_Acceptance_GrayResize_Lanczos_2X(t *testing.T) {
	gray := setupTestCaseGray(t)
	actual, _ := ResizeGray(gray, 2, 2, InterLanczos)
	originalSize := gray.Bounds().Size()
	expectedSize := image.Point{X: originalSize.X * 2, Y: originalSize.Y * 2}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/grayResize_Lanczos_2x.jpg")
}

func Test_Acceptance_RGBAResize_NN_2X(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, 2, 2, InterNearest)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: originalSize.X * 2, Y: originalSize.Y * 2}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize2x.jpg")
}

func Test_Acceptance_RGBAResize_NN_1_5X(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, 1.5, 1.5, InterNearest)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * 1.5), Y: int(float64(originalSize.Y) * 1.5)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize1_5x.jpg")
}

func Test_Acceptance_RGBAResize_NN_2_5_AND_3_5X(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, 2.5, 3.5, InterNearest)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * 2.5), Y: int(float64(originalSize.Y) * 3.5)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize2_5_and_3_5x.jpg")
}

func Test_Acceptance_RGBAResize_NN_0_5X(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, 0.5, 0.5, InterNearest)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: int(float64(originalSize.X) * 0.5), Y: int(float64(originalSize.Y) * 0.5)}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize0_5x.jpg")
}

func Test_Acceptance_RGBAResize_Linear_2X(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, 2, 2, InterLinear)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: originalSize.X * 2, Y: originalSize.Y * 2}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize_Linear_2x.jpg")
}

func Test_Acceptance_RGBAResize_CatmullRom_2X(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, 2, 2, InterCatmullRom)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: originalSize.X * 2, Y: originalSize.Y * 2}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize_CatmullRom_2x.jpg")
}

func Test_Acceptance_RGBAResize_Lanczos_2X(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	actual, _ := ResizeRGBA(rgba, 2, 2, InterLanczos)
	originalSize := rgba.Bounds().Size()
	expectedSize := image.Point{X: originalSize.X * 2, Y: originalSize.Y * 2}
	actualSize := actual.Bounds().Size()
	if expectedSize.X != actualSize.X || expectedSize.Y != actualSize.Y {
		t.Errorf("Expected size of [%d, %d] does not match actual size of [%d, %d]!", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, actual, "../res/resize/rgbaResize_Lanczos_2x.jpg")
}

// ----------------------------------------------------------------------------------
