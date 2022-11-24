package edgedetection

import (
	"github.com/ernyoke/imger/imgio"
	"github.com/ernyoke/imger/padding"
	"image"
	"testing"
)

// -----------------------------Acceptance tests------------------------------------

func setupTestCaseGrayLapl(t *testing.T) *image.Gray {
	path := "../res/engine.png"
	img, err := imgio.ImreadGray(path)
	if err != nil {
		t.Errorf("Could not read image from path: %s", path)
	}
	return img
}

func setupTestCaseRGBALapl(t *testing.T) *image.RGBA {
	path := "../res/engine.png"
	img, err := imgio.ImreadRGBA(path)
	if err != nil {
		t.Errorf("Could not read image from path: %s", path)
	}
	return img
}

func tearDownTestCaseLapl(t *testing.T, img image.Image, path string) {
	err := imgio.Imwrite(img, path)
	if err != nil {
		t.Errorf("Could not write image to path: %s", path)
	}
}

func Test_Acceptance_LaplacianGrayK4(t *testing.T) {
	gray := setupTestCaseGrayLapl(t)
	laplacian, _ := LaplacianGray(gray, padding.BorderReflect, K4)
	tearDownTestCaseLapl(t, laplacian, "../res/edge/laplacianGrayK4.png")
}

func Test_Acceptance_LaplacianGrayK8(t *testing.T) {
	gray := setupTestCaseGrayLapl(t)
	laplacian, _ := LaplacianGray(gray, padding.BorderReflect, K8)
	tearDownTestCaseLapl(t, laplacian, "../res/edge/laplacianGrayK8.png")
}

func Test_Acceptance_LaplacianRGBAK4(t *testing.T) {
	rgba := setupTestCaseRGBALapl(t)
	laplacian, _ := LaplacianRGBA(rgba, padding.BorderReflect, K4)
	tearDownTestCaseLapl(t, laplacian, "../res/edge/laplacianRGBAK4.png")
}

func Test_Acceptance_LaplacianRGBAK8(t *testing.T) {
	rgba := setupTestCaseRGBALapl(t)
	laplacian, _ := LaplacianRGBA(rgba, padding.BorderReflect, K8)
	tearDownTestCaseLapl(t, laplacian, "../res/edge/laplacianRGBAK8.png")
}

// ---------------------------------------------------------------------------------
