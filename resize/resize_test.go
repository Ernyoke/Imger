package resize

import (
	"testing"
	"image"
	"os"
	"image/jpeg"
	"fmt"
	"image/draw"
	"github.com/ernyoke/imger/grayscale"
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

func Test_Acceptance_GrayResize_NN_2X(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
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
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
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
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
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
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
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
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
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
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
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
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
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
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
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
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
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
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
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
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
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
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
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
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
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
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
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