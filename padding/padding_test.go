package padding

import (
	"github.com/Ernyoke/Imger/imgio"
	"github.com/Ernyoke/Imger/utils"
	"image"
	"testing"
)

// ---------------------------------Unit tests--------------------------------------
func Test_GrayPaddingBorderConstant_1pxPadding(t *testing.T) {
	gray := image.Gray{
		Rect:   image.Rect(0, 0, 5, 3),
		Stride: 5,
		Pix: []uint8{
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
		},
	}
	expected := image.Gray{
		Rect:   image.Rect(0, 0, 7, 5),
		Stride: 7,
		Pix: []uint8{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0x00,
			0x00, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0x00,
			0x00, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
	}
	paddingSize := image.Point{X: 3, Y: 3}
	anchor := image.Point{X: 1, Y: 1}
	actual, _ := PaddingGray(&gray, paddingSize, anchor, BorderConstant)
	utils.CompareGrayImages(t, &expected, actual)
}

func Test_GrayPaddingBorderConstant_2pxPadding(t *testing.T) {
	gray := image.Gray{
		Rect:   image.Rect(0, 0, 5, 3),
		Stride: 5,
		Pix: []uint8{
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
		},
	}
	expected := image.Gray{
		Rect:   image.Rect(0, 0, 9, 7),
		Stride: 9,
		Pix: []uint8{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0x00, 0x00,
			0x00, 0x00, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0x00, 0x00,
			0x00, 0x00, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
	}
	paddingSize := image.Point{X: 5, Y: 5}
	anchor := image.Point{X: 2, Y: 2}
	actual, _ := PaddingGray(&gray, paddingSize, anchor, BorderConstant)
	//utils.PrintGray(t, actual)
	utils.CompareGrayImages(t, &expected, actual)
}

func Test_GrayPaddingBorderConstant_1_3pxPadding(t *testing.T) {
	gray := image.Gray{
		Rect:   image.Rect(0, 0, 5, 3),
		Stride: 5,
		Pix: []uint8{
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
		},
	}
	expected := image.Gray{
		Rect:   image.Rect(0, 0, 9, 7),
		Stride: 9,
		Pix: []uint8{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0x00, 0x00, 0x00,
			0x00, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0x00, 0x00, 0x00,
			0x00, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
	}
	paddingSize := image.Point{X: 5, Y: 5}
	anchor := image.Point{X: 1, Y: 1}
	actual, _ := PaddingGray(&gray, paddingSize, anchor, BorderConstant)
	//utils.PrintGray(t, actual)
	utils.CompareGrayImages(t, &expected, actual)
}

func Test_GrayPaddingBorderReplicate_1pxPadding(t *testing.T) {
	gray := image.Gray{
		Rect:   image.Rect(0, 0, 5, 3),
		Stride: 5,
		Pix: []uint8{
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
		},
	}
	expected := image.Gray{
		Rect:   image.Rect(0, 0, 7, 5),
		Stride: 7,
		Pix: []uint8{
			0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE,
			0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE,
			0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE,
			0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE,
			0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE,
		},
	}
	paddingSize := image.Point{X: 3, Y: 3}
	anchor := image.Point{X: 1, Y: 1}
	actual, _ := PaddingGray(&gray, paddingSize, anchor, BorderReplicate)
	utils.CompareGrayImages(t, &expected, actual)
}

func Test_GrayPaddingBorderReplicate_2pxPadding(t *testing.T) {
	gray := image.Gray{
		Rect:   image.Rect(0, 0, 5, 3),
		Stride: 5,
		Pix: []uint8{
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
		},
	}
	expected := image.Gray{
		Rect:   image.Rect(0, 0, 9, 7),
		Stride: 9,
		Pix: []uint8{
			0xAA, 0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE,
			0xAA, 0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE,
			0xAA, 0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE,
			0xAA, 0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE,
			0xAA, 0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE,
			0xAA, 0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE,
			0xAA, 0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE,
		},
	}
	paddingSize := image.Point{X: 5, Y: 5}
	anchor := image.Point{X: 2, Y: 2}
	actual, _ := PaddingGray(&gray, paddingSize, anchor, BorderReplicate)
	//utils.PrintGray(t, actual)
	utils.CompareGrayImages(t, &expected, actual)
}

func Test_GrayPaddingBorderReplicate_1_3pxPadding(t *testing.T) {
	gray := image.Gray{
		Rect:   image.Rect(0, 0, 5, 3),
		Stride: 5,
		Pix: []uint8{
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
		},
	}
	expected := image.Gray{
		Rect:   image.Rect(0, 0, 9, 7),
		Stride: 9,
		Pix: []uint8{
			0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE, 0xEE,
			0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE, 0xEE,
			0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE, 0xEE,
			0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE, 0xEE,
			0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE, 0xEE,
			0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE, 0xEE,
			0xAA, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xEE, 0xEE, 0xEE,
		},
	}
	paddingSize := image.Point{X: 5, Y: 5}
	anchor := image.Point{X: 1, Y: 1}
	actual, _ := PaddingGray(&gray, paddingSize, anchor, BorderReplicate)
	//utils.PrintGray(t, actual)
	utils.CompareGrayImages(t, &expected, actual)
}

func Test_GrayPaddingBorderReflect_1_3pxPadding(t *testing.T) {
	gray := image.Gray{
		Rect:   image.Rect(0, 0, 5, 3),
		Stride: 5,
		Pix: []uint8{
			0xAA, 0xBB, 0xCC, 0xDD, 0xEE,
			0x11, 0xBB, 0xCC, 0xDD, 0xEE,
			0x22, 0xBB, 0xCC, 0xDD, 0xEE,
		},
	}
	expected := image.Gray{
		Rect:   image.Rect(0, 0, 9, 6),
		Stride: 9,
		Pix: []uint8{
			0xBB, 0x11, 0xBB, 0xCC, 0xDD, 0xEE, 0xDD, 0xCC, 0xBB,
			0xBB, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xDD, 0xCC, 0xBB,
			0xBB, 0x11, 0xBB, 0xCC, 0xDD, 0xEE, 0xDD, 0xCC, 0xBB,
			0xBB, 0x22, 0xBB, 0xCC, 0xDD, 0xEE, 0xDD, 0xCC, 0xBB,
			0xBB, 0x11, 0xBB, 0xCC, 0xDD, 0xEE, 0xDD, 0xCC, 0xBB,
			0xBB, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xDD, 0xCC, 0xBB,
		},
	}
	paddingSize := image.Point{X: 5, Y: 4}
	anchor := image.Point{X: 1, Y: 1}
	actual, _ := PaddingGray(&gray, paddingSize, anchor, BorderReflect)
	//utils.PrintGray(t, actual)
	utils.CompareGrayImages(t, &expected, actual)
}

func Test_GrayPaddingBorderReflect_2pxPadding(t *testing.T) {
	gray := image.Gray{
		Rect:   image.Rect(0, 0, 4, 4),
		Stride: 4,
		Pix: []uint8{
			0xAA, 0xBB, 0xCC, 0xDD,
			0x11, 0x11, 0x11, 0x11,
			0x22, 0x22, 0x22, 0x22,
			0x33, 0x33, 0x33, 0x33,
		},
	}
	expected := image.Gray{
		Rect:   image.Rect(0, 0, 8, 8),
		Stride: 8,
		Pix: []uint8{
			0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22,
			0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11,
			0xCC, 0xBB, 0xAA, 0xBB, 0xCC, 0xDD, 0xCC, 0xBB,
			0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11,
			0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22,
			0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33,
			0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22,
			0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11,
		},
	}
	paddingSize := image.Point{X: 5, Y: 5}
	anchor := image.Point{X: 2, Y: 2}
	actual, _ := PaddingGray(&gray, paddingSize, anchor, BorderReflect)
	//utils.PrintGray(t, actual)
	utils.CompareGrayImages(t, &expected, actual)
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

func Test_Acceptance_GrayPaddingBorderConstant(t *testing.T) {
	gray := setupTestCaseGray(t)
	padded, _ := PaddingGray(gray, image.Point{X: 15, Y: 15}, image.Point{X: 8, Y: 8}, BorderConstant)
	tearDownTestCase(t, padded, "../res/padding/grayPaddingBorderConstant.jpg")
}

func Test_Acceptance_GrayPaddingBorderConstantDistortedAnchor(t *testing.T) {
	gray := setupTestCaseGray(t)
	padded, _ := PaddingGray(gray, image.Point{X: 50, Y: 50}, image.Point{X: 8, Y: 8}, BorderConstant)
	tearDownTestCase(t, padded, "../res/padding/grayPaddingBorderConstantDistortedAnchor.jpg")
}

func Test_Acceptance_GrayPaddingBorderReplicate(t *testing.T) {
	gray := setupTestCaseGray(t)
	padded, _ := PaddingGray(gray, image.Point{X: 15, Y: 15}, image.Point{X: 8, Y: 8}, BorderReplicate)
	tearDownTestCase(t, padded, "../res/padding/grayPaddingBorderReplicate.jpg")
}

func Test_Acceptance_GrayPaddingBorderReflect(t *testing.T) {
	gray := setupTestCaseGray(t)
	padded, _ := PaddingGray(gray, image.Point{X: 15, Y: 15}, image.Point{X: 8, Y: 8}, BorderReflect)
	tearDownTestCase(t, padded, "../res/padding/grayPaddingBorderReflect.jpg")
}

func Test_Acceptance_RGBAPaddingBorderConstant(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	padded, _ := PaddingRGBA(rgba, image.Point{X: 15, Y: 15}, image.Point{X: 8, Y: 8}, BorderConstant)
	tearDownTestCase(t, padded, "../res/padding/rgbaPaddedBorderConstant.jpg")
}

func Test_Acceptance_RGBAPaddingBorderReplicate(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	padded, _ := PaddingRGBA(rgba, image.Point{X: 15, Y: 15}, image.Point{X: 8, Y: 8}, BorderReplicate)
	tearDownTestCase(t, padded, "../res/padding/rgbaPaddedBorderReplicate.jpg")
}

func Test_Acceptance_RGBAPaddingBorderReflect(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	padded, _ := PaddingRGBA(rgba, image.Point{X: 15, Y: 15}, image.Point{X: 8, Y: 8}, BorderReflect)
	tearDownTestCase(t, padded, "../res/padding/rgbaPaddedBorderReflect.jpg")
}

// ---------------------------------------------------------------------------------