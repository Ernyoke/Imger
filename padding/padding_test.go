package padding

import (
	"fmt"
	"github.com/ernyoke/imgur/grayscale"
	"github.com/ernyoke/imgur/utils"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
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

func Test_Acceptance_GrayPaddingBorderConstant(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	padded, _ := PaddingGray(gray, image.Point{15, 15}, image.Point{8, 8}, BorderConstant)
	tearDownTestCase(t, padded, "../res/padding/grayPaddingBorderConstant.jpg")
}

func Test_Acceptance_GrayPaddingBorderConstantDistortedAnchor(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	padded, _ := PaddingGray(gray, image.Point{50, 50}, image.Point{8, 8}, BorderConstant)
	tearDownTestCase(t, padded, "../res/padding/grayPaddingBorderConstantDistortedAnchor.jpg")
}

func Test_Acceptance_GrayPaddingBorderReplicate(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	padded, _ := PaddingGray(gray, image.Point{15, 15}, image.Point{8, 8}, BorderReplicate)
	tearDownTestCase(t, padded, "../res/padding/grayPaddingBorderReplicate.jpg")
}

func Test_Acceptance_GrayPaddingBorderReflect(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	padded, _ := PaddingGray(gray, image.Point{15, 15}, image.Point{8, 8}, BorderReflect)
	tearDownTestCase(t, padded, "../res/padding/grayPaddingBorderReflect.jpg")
}

func Test_Acceptance_RGBAPaddingBorderConstant(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	padded, _ := PaddingRGBA(rgba, image.Point{15, 15}, image.Point{8, 8}, BorderConstant)
	tearDownTestCase(t, padded, "../res/padding/rgbaPaddedBorderConstant.jpg")
}

func Test_Acceptance_RGBAPaddingBorderReplicate(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	padded, _ := PaddingRGBA(rgba, image.Point{15, 15}, image.Point{8, 8}, BorderReplicate)
	tearDownTestCase(t, padded, "../res/padding/rgbaPaddedBorderReplicate.jpg")
}

func Test_Acceptance_RGBAPaddingBorderReflect(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	padded, _ := PaddingRGBA(rgba, image.Point{15, 15}, image.Point{8, 8}, BorderReflect)
	tearDownTestCase(t, padded, "../res/padding/rgbaPaddedBorderReflect.jpg")
}

// ---------------------------------------------------------------------------------
