package effects

import (
	"github.com/ernyoke/imger/imgio"
	"github.com/ernyoke/imger/utils"
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

func Test_InvertedGray(t *testing.T) {
	gray := image.Gray{
		Rect:   image.Rect(0, 0, 4, 1),
		Stride: 4,
		Pix: []uint8{
			0x00, 0xFF, 0x80, 0xAB,
		},
	}
	expected := image.Gray{
		Rect:   image.Rect(0, 0, 4, 1),
		Stride: 4,
		Pix: []uint8{
			0xFF, 0x00, 0x7F, 0x54,
		},
	}
	actual := InvertGray(&gray)
	utils.CompareGrayImages(t, &expected, actual)
}

func Test_InvertedRGBA(t *testing.T) {
	rgba := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 1),
		Stride: 4,
		Pix: []uint8{
			0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x7F, 0xAB, 0xFF,
		},
	}
	expected := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 1),
		Stride: 4,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x7F, 0x80, 0x54, 0xFF,
		},
	}
	actual := InvertRGBA(&rgba)
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

func Test_Acceptance_EmbossGray(t *testing.T) {
	gray := setupTestCaseGray(t)
	emboss, err := EmbossGray(gray)
	if err != nil {
		t.Fatalf("Should not reach this point!")
	}
	tearDownTestCase(t, emboss, "../res/effects/embossGray.jpg")
}

func Test_Acceptance_EmbossRGBA(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	emboss, err := EmbossRGBA(rgba)
	if err != nil {
		t.Fatalf("Should not reach this point!")
	}
	tearDownTestCase(t, emboss, "../res/effects/embossRGBA.jpg")
}

func Test_Acceptance_SharpenGray(t *testing.T) {
	gray := setupTestCaseGray(t)
	sharp, err := SharpenGray(gray)
	if err != nil {
		t.Fatalf("Should not reach this point!")
	}
	tearDownTestCase(t, sharp, "../res/effects/sharpenGray.jpg")
}

func Test_Acceptance_SharpenRGBA(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	sharp, err := SharpenRGBA(rgba)
	if err != nil {
		t.Fatalf("Should not reach this point!")
	}
	tearDownTestCase(t, sharp, "../res/effects/sharpenRGBA.jpg")
}

func Test_Acceptance_InvertGray(t *testing.T) {
	gray := setupTestCaseGray(t)
	inerted := InvertGray(gray)
	tearDownTestCase(t, inerted, "../res/effects/invertedGray.jpg")
}

func Test_Acceptance_InvertedRGBA(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	inverted := InvertRGBA(rgba)
	tearDownTestCase(t, inverted, "../res/effects/invertedRGBA.jpg")
}
