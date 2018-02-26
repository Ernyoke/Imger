package edgedetection

import (
	"fmt"
	"github.com/ernyoke/imger/grayscale"
	"github.com/ernyoke/imger/padding"
	"image"
	"image/draw"
	"image/png"
	"os"
	"testing"
)

// -----------------------------Acceptance tests------------------------------------

func setupTestLaplacianCase(t *testing.T) image.Image {
	imagePath := "../res/engine.png"
	file, err := os.Open(imagePath)
	defer file.Close()
	if err != nil {
		t.Log(os.Stderr, "%v\n", err)
	}
	img, err := png.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return img
}

func tearDownLaplacianTestCase(t *testing.T, image image.Image, path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, image)
}

func Test_Acceptance_LaplacianGrayK4(t *testing.T) {
	img := setupTestLaplacianCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	laplacian, _ := LaplacianGray(gray, padding.BorderReflect, K4)
	tearDownLaplacianTestCase(t, laplacian, "../res/sobel/laplacianGrayK4.png")
}

func Test_Acceptance_LaplacianGrayK8(t *testing.T) {
	img := setupTestLaplacianCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	gray := grayscale.Grayscale(rgba)
	laplacian, _ := LaplacianGray(gray, padding.BorderReflect, K8)
	tearDownLaplacianTestCase(t, laplacian, "../res/sobel/laplacianGrayK8.png")
}

func Test_Acceptance_LaplacianRGBAK4(t *testing.T) {
	img := setupTestLaplacianCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	laplacian, _ := LaplacianRGBA(rgba, padding.BorderReflect, K4)
	tearDownLaplacianTestCase(t, laplacian, "../res/sobel/laplacianRGBAK4.png")
}

func Test_Acceptance_LaplacianRGBAK8(t *testing.T) {
	img := setupTestLaplacianCase(t)
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	laplacian, _ := LaplacianRGBA(rgba, padding.BorderReflect, K8)
	tearDownLaplacianTestCase(t, laplacian, "../res/sobel/laplacianRGBAK8.png")
}

// ---------------------------------------------------------------------------------
