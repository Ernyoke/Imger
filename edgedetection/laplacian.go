package edgedetection

import (
	"errors"
	"github.com/ernyoke/imger/convolution"
	"github.com/ernyoke/imger/grayscale"
	"github.com/ernyoke/imger/padding"
	"image"
)

var kernel4 = convolution.Kernel{Content: [][]float64{
	{0, 1, 0},
	{1, -4, 1},
	{0, 1, 0},
}, Width: 3, Height: 3}

var kernel8 = convolution.Kernel{Content: [][]float64{
	{1, 1, 1},
	{1, -8, 1},
	{1, 1, 1},
}, Width: 3, Height: 3}

// LaplacianKernel - constant type for differentiating Laplacian kernels
type LaplacianKernel int

const (
	// K4 Laplacian kernel:
	//	{0, 1, 0},
	//	{1, -4, 1},
	//	{0, 1, 0},
	K4 LaplacianKernel = iota
	// K8 Laplacian kernel:
	//	{0, 1, 0},
	//	{1, -8, 1},
	//	{0, 1, 0},
	K8
)

// LaplacianGray applies Laplacian filter to a grayscale image. The kernel types are: K4 and K8 (see LaplacianKernel)
// Example of usage:
//
//		 res, err := edgedetection.LaplacianGray(img, paddding.BorderReflect, edgedetection.K8)
//
func LaplacianGray(gray *image.Gray, border padding.Border, kernel LaplacianKernel) (*image.Gray, error) {
	var laplacianKernel convolution.Kernel
	switch kernel {
	case K4:
		laplacianKernel = kernel4
	case K8:
		laplacianKernel = kernel8
	default:
		return nil, errors.New("invalid kernel")
	}
	return convolution.ConvolveGray(gray, &laplacianKernel, image.Point{X: 1, Y: 1}, border)
}

// LaplacianRGBA applies Laplacian filter to an RGBA image. The kernel types are: K4 and K8 (see LaplacianKernel)
// Example of usage:
//
//		 res, err := edgedetection.LaplacianRGBA(img, paddding.BorderReflect, edgedetection.K8)
//
func LaplacianRGBA(img *image.RGBA, border padding.Border, kernel LaplacianKernel) (*image.Gray, error) {
	gray := grayscale.Grayscale(img)
	return LaplacianGray(gray, border, kernel)
}
