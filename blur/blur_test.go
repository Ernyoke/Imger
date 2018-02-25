package blur

import (
	"fmt"
	"github.com/ernyoke/imgur/grayscale"
	"github.com/ernyoke/imgur/padding"
	"github.com/ernyoke/imgur/utils"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"testing"
)

// ---------------------------------Unit tests------------------------------------
func TestRGBAGaussianBlurZeroRadius(t *testing.T) {
	input := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
			0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
		},
	}
	_, error := GaussianBlurRGBA(&input, 0, 6, padding.BorderReflect)
	if error != nil {
		//ok
	} else {
		t.Fatal("no error thrown")
	}
}

func TestRGBAGaussianBlurOneRadius(t *testing.T) {
	input := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
			0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
		},
	}
	expected := &image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0xae, 0xae, 0xae, 0xff, 0x56, 0x56, 0x56, 0xff, 0x19, 0x19, 0x19, 0xff,
			0x4d, 0x4d, 0x4d, 0xff, 0x68, 0x68, 0x68, 0xff, 0x8b, 0x8b, 0x8b, 0xff,
			0x0, 0x0, 0x0, 0xff, 0x26, 0x26, 0x26, 0xff, 0x59, 0x59, 0x59, 0xff,
		},
	}
	f, err := os.Create("../res/asd.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, &input, nil)

	result, error := GaussianBlurRGBA(&input, 1, 2, padding.BorderConstant)
	if error != nil {
		t.Fatal(error)
	}
	utils.CompareRGBAImages(t, expected, result)
}

// ---------------------------------------------------------------------------------

// -----------------------------Acceptance tests------------------------------------
func TestReadImg(t *testing.T) {
	imagePath := "../res/image.jpg"
	file, err := os.Open(imagePath)
	defer file.Close()
	if err != nil {
		t.Log(os.Stderr, "%v\n", err)
	}
	img, err := jpeg.Decode(file)
	rgba := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	imgSize := rgba.Bounds().Size()
	for x := 0; x < imgSize.X; x++ {
		for y := 0; y < imgSize.Y; y++ {
			c := rgba.RGBAAt(x, y)
			fmt.Printf("%x %x %x %x\n", c.R, c.G, c.B, c.A)
		}
	}
}

// ----------------------------------------------------------------------------------
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

func Test_Acceptance_GrayBlurInt(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	blured, _ := BlurGray(gray, image.Point{15, 15}, image.Point{8, 8}, padding.BorderReflect)
	tearDownTestCase(t, blured, "../res/blur/grayBlur.jpg")
}

func Test_Acceptance_RGBABlurInt(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	blured, _ := BlurRGBA(rgba, image.Point{15, 15}, image.Point{8, 8}, padding.BorderReflect)
	tearDownTestCase(t, blured, "../res/blur/rgbaBlur.jpg")
}

func Test_Acceptance_GrayGaussianBlurInt(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	blured, _ := GaussianBlurGray(gray, 7, 6, padding.BorderReflect)
	tearDownTestCase(t, blured, "../res/blur/grayGaussianBlur.jpg")
}

func Test_Acceptance_RGBAGaussianBlurInt(t *testing.T) {
	img := setupTestCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	blured, _ := GaussianBlurRGBA(rgba, 5, 500, padding.BorderReflect)
	tearDownTestCase(t, blured, "../res/blur/rgbaGaussianBlur.jpg")
}
