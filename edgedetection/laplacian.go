package edgedetection

import (
	"errors"
	"github.com/Ernyoke/Imger/convolution"
	"github.com/Ernyoke/Imger/grayscale"
	"github.com/Ernyoke/Imger/padding"
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
