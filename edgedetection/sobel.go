package edgedetection

import (
	"github.com/Ernyoke/Imger/blend"
	"github.com/Ernyoke/Imger/convolution"
	"github.com/Ernyoke/Imger/grayscale"
	"github.com/Ernyoke/Imger/padding"
	"image"
)

var horizontalKernel = convolution.Kernel{Content: [][]float64{
	{-1, 0, 1},
	{-2, 0, 2},
	{-1, 0, 1},
}, Width: 3, Height: 3}

var verticalKernel = convolution.Kernel{Content: [][]float64{
	{-1, -2, -1},
	{0, 0, 0},
	{1, 2, 1},
}, Width: 3, Height: 3}

func HorizontalSobelGray(gray *image.Gray, border padding.Border) (*image.Gray, error) {
	return convolution.ConvolveGray(gray, &horizontalKernel, image.Point{X: 1, Y: 1}, border)
}

func VerticalSobelGray(gray *image.Gray, border padding.Border) (*image.Gray, error) {
	return convolution.ConvolveGray(gray, &verticalKernel, image.Point{X: 1, Y: 1}, border)
}

func SobelGray(img *image.Gray, border padding.Border) (*image.Gray, error) {
	horizontal, error := HorizontalSobelGray(img, border)
	if error != nil {
		return nil, error
	}
	vertical, error := VerticalSobelGray(img, border)
	if error != nil {
		return nil, error
	}
	res, error := blend.AddGrayWeighted(horizontal, 0.5, vertical, 0.5)
	if error != nil {
		return nil, error
	}
	return res, nil
}

func HorizontalSobelRGBA(img *image.RGBA, border padding.Border) (*image.Gray, error) {
	gray := grayscale.Grayscale(img)
	return convolution.ConvolveGray(gray, &horizontalKernel, image.Point{X: 1, Y: 1}, border)
}

func VerticalSobelRGBA(img *image.RGBA, border padding.Border) (*image.Gray, error) {
	gray := grayscale.Grayscale(img)
	return convolution.ConvolveGray(gray, &verticalKernel, image.Point{X: 1, Y: 1}, border)
}

func SobelRGBA(img *image.RGBA, border padding.Border) (*image.Gray, error) {
	gray := grayscale.Grayscale(img)
	return SobelGray(gray, border)
}
