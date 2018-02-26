package edgedetection

import (
	"errors"
	"github.com/ernyoke/imger/convolution"
	"github.com/ernyoke/imger/grayscale"
	"github.com/ernyoke/imger/padding"
	"image"
)

var kernel4 = convolution.Kernel{[][]float64{
	{0, 1, 0},
	{1, -4, 1},
	{0, 1, 0},
}, 3, 3}

var kernel8 = convolution.Kernel{[][]float64{
	{1, 1, 1},
	{1, -8, 1},
	{1, 1, 1},
}, 3, 3}

type LaplacianKernel int

const (
	K4 LaplacianKernel = iota
	K8
)

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

func LaplacianRGBA(img *image.RGBA, border padding.Border, kernel LaplacianKernel) (*image.Gray, error) {
	gray := grayscale.Grayscale(img)
	return LaplacianGray(gray, border, kernel)
}
