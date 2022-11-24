package generate

import (
	"github.com/ernyoke/imger/imgio"
	"image"
	"image/color"
	"testing"
)

// -----------------------------Acceptance tests------------------------------------
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

func Test_Acceptance_LinearGradientHorizontal(t *testing.T) {
	res := LinearGradient(image.Point{X: 500, Y: 200}, color.RGBA{R: 0, G: 0, B: 0, A: 255}, color.RGBA{R: 255, G: 255, B: 255, A: 255}, H)
	tearDownTestCase(t, res, "../res/generate/linearGradientHorizontal.jpg")
}

func Test_Acceptance_LinearGradientVertical(t *testing.T) {
	res := LinearGradient(image.Point{X: 500, Y: 200}, color.RGBA{R: 0, G: 0, B: 0, A: 255}, color.RGBA{R: 255, G: 255, B: 255, A: 255}, V)
	tearDownTestCase(t, res, "../res/generate/linearGradientVertical.jpg")
}

func Test_Acceptance_SigmoidalHorizontal(t *testing.T) {
	res := SigmoidalGradient(image.Point{X: 500, Y: 200}, color.RGBA{R: 0, G: 0, B: 0, A: 255}, color.RGBA{R: 255, G: 255, B: 255, A: 255}, H)
	tearDownTestCase(t, res, "../res/generate/sigmoidalGradientHorizontal.jpg")
}

func Test_Acceptance_SigmoidalGradientVertical(t *testing.T) {
	res := SigmoidalGradient(image.Point{X: 500, Y: 200}, color.RGBA{R: 0, G: 0, B: 0, A: 255}, color.RGBA{R: 255, G: 255, B: 255, A: 255}, V)
	tearDownTestCase(t, res, "../res/generate/sigmoidalGradientVertical.jpg")
}

// ---------------------------------------------------------------------------------
