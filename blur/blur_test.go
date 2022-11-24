package blur

import (
	"github.com/ernyoke/imger/imgio"
	"github.com/ernyoke/imger/padding"
	"github.com/ernyoke/imger/utils"
	"image"
	"testing"
)

// ---------------------------------Unit tests------------------------------------
func TestGrayGaussianBlurZeroRadius(t *testing.T) {
	input := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
			0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
		},
	}
	_, err := GaussianBlurRGBA(&input, 0, 6, padding.BorderReflect)
	if err != nil {
		//ok
	} else {
		t.Fatal("no error thrown")
	}
}

func TestGrayGaussianBlurOneRadius(t *testing.T) {
	input := image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0x80, 0x56,
			0x56, 0x80, 0x69,
			0xEE, 0x29, 0xBB,
		},
	}
	expected := &image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0x47, 0x5A, 0x33,
			0x64, 0x88, 0x4D,
			0x3B, 0x59, 0x36,
		},
	}
	result, err := GaussianBlurGray(&input, 1, 2, padding.BorderConstant)
	if err != nil {
		t.Fatal(err)
	}
	utils.CompareGrayImagesWithOffset(t, expected, result, 1)
}

func TestRGBAGaussianBlurOneRadius(t *testing.T) {
	input := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0x56, 0x56, 0x56, 0xFF,
			0x56, 0x56, 0x56, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0x69, 0x69, 0x69, 0xFF,
			0xEE, 0xEE, 0xEE, 0xFF, 0x29, 0x29, 0x29, 0xFF, 0xBB, 0xBB, 0xBB, 0xFF,
		},
	}
	expected := &image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0x47, 0x47, 0x47, 0xFF, 0x5A, 0x5A, 0x5A, 0xFF, 0x33, 0x33, 0x33, 0xFF,
			0x64, 0x64, 0x64, 0xFF, 0x88, 0x88, 0x88, 0xFF, 0x4D, 0x4D, 0x4D, 0xFF,
			0x3B, 0x3B, 0x3B, 0xFF, 0x59, 0x59, 0x59, 0xFF, 0x36, 0x36, 0x36, 0xFF,
		},
	}
	actual, err := GaussianBlurRGBA(&input, 1, 2, padding.BorderConstant)
	if err != nil {
		t.Fatal(err)
	}
	//utils.PrintRGBA(t, actual)
	utils.CompareRGBAImagesWithOffset(t, expected, actual, 1)
}

// ---------------------------------------------------------------------------------

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

func Test_Acceptance_GrayBlurInt(t *testing.T) {
	gray := setupTestCaseGray(t)
	blured, _ := BoxGray(gray, image.Point{X: 15, Y: 15}, image.Point{X: 8, Y: 8}, padding.BorderReflect)
	tearDownTestCase(t, blured, "../res/blur/grayBlur.jpg")
}

func Test_Acceptance_RGBABlurInt(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	blured, _ := BoxRGBA(rgba, image.Point{X: 15, Y: 15}, image.Point{X: 8, Y: 8}, padding.BorderReflect)
	tearDownTestCase(t, blured, "../res/blur/rgbaBlur.jpg")
}

func Test_Acceptance_GrayGaussianBlurInt(t *testing.T) {
	gray := setupTestCaseGray(t)
	blured, _ := GaussianBlurGray(gray, 7, 6, padding.BorderReflect)
	tearDownTestCase(t, blured, "../res/blur/grayGaussianBlur.jpg")
}

func Test_Acceptance_RGBAGaussianBlurInt(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	blured, _ := GaussianBlurRGBA(rgba, 5, 500, padding.BorderReflect)
	tearDownTestCase(t, blured, "../res/blur/rgbaGaussianBlur.jpg")
}

// ----------------------------------------------------------------------------------
