package effects

import (
	"github.com/Ernyoke/Imger/imgio"
	"github.com/Ernyoke/Imger/utils"
	"image"
	"testing"
)

// --------------------------------Unit tests---------------------------------------
func Test_Sepia(t *testing.T) {
	rgba := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 1),
		Stride: 4,
		Pix: []uint8{
			0x01, 0x01, 0x01, 0xFF, 0x01, 0x01, 0x01, 0xFF, 0x01, 0x01, 0x01, 0xFF,
		},
	}
	expected := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 1),
		Stride: 4,
		Pix: []uint8{
			0x01, 0x01, 0x00, 0xFF, 0x01, 0x01, 0x00, 0xFF, 0x01, 0x01, 0x00, 0xFF,
		},
	}
	actual := Sepia(&rgba)
	utils.CompareRGBAImages(t, &expected, actual)
}

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

func Test_Acceptance_PixelateGray(t *testing.T) {
	rgba := setupTestCaseGray(t)
	sepia, err := PixelateGray(rgba, 5)
	if err != nil {
		t.Fatalf("Should not reach this point!")
	}
	tearDownTestCase(t, sepia, "../res/effects/pixelateGray.jpg")
}

func Test_Acceptance_PixelateRGBA(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	sepia, err := PixelateRGBA(rgba, 5)
	if err != nil {
		t.Fatalf("Should not reach this point!")
	}
	tearDownTestCase(t, sepia, "../res/effects/pixelateRGBA.jpg")
}

func Test_Acceptance_Sepia(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	sepia := Sepia(rgba)
	tearDownTestCase(t, sepia, "../res/effects/sepia.jpg")
}
