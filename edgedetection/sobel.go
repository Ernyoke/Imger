package edgedetection

import (
	"github.com/ernyoke/imger/blend"
	"github.com/ernyoke/imger/convolution"
	"github.com/ernyoke/imger/grayscale"
	"github.com/ernyoke/imger/padding"
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

// HorizontalSobelGray applies the horizontal Sobel operator (horizontal kernel) to a grayscale image. The result
// of the Sobel operator is a 2-dimensional map of the gradient at each point.
// More information on the Sobel operator: https://en.wikipedia.org/wiki/Sobel_operator
func HorizontalSobelGray(gray *image.Gray, border padding.Border) (*image.Gray, error) {
	return convolution.ConvolveGray(gray, &horizontalKernel, image.Point{X: 1, Y: 1}, border)
}

// VerticalSobelGray applies the vertical Sobel operator (vertical kernel) to a grayscale image. The result
// of the Sobel operator is a 2-dimensional map of the gradient at each point.
// More information on the Sobel operator: https://en.wikipedia.org/wiki/Sobel_operator
func VerticalSobelGray(gray *image.Gray, border padding.Border) (*image.Gray, error) {
	return convolution.ConvolveGray(gray, &verticalKernel, image.Point{X: 1, Y: 1}, border)
}

// SobelGray combines the horizontal and the vertical gradients of a grayscale image. The result is grayscale image
// which contains the high gradients ("edges") marked as white.
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

// HorizontalSobelRGBA applies the horizontal Sobel operator (horizontal kernel) to an RGGBA image. The result
// of the Sobel operator is a 2-dimensional map of the gradient at each point.
// More information on the Sobel operator: https://en.wikipedia.org/wiki/Sobel_operator
func HorizontalSobelRGBA(img *image.RGBA, border padding.Border) (*image.Gray, error) {
	gray := grayscale.Grayscale(img)
	return convolution.ConvolveGray(gray, &horizontalKernel, image.Point{X: 1, Y: 1}, border)
}

// VerticalSobelRGBA applies the vertical Sobel operator (vertical kernel) to an RGBA image. The result
// of the Sobel operator is a 2-dimensional map of the gradient at each point.
// More information on the Sobel operator: https://en.wikipedia.org/wiki/Sobel_operator
func VerticalSobelRGBA(img *image.RGBA, border padding.Border) (*image.Gray, error) {
	gray := grayscale.Grayscale(img)
	return convolution.ConvolveGray(gray, &verticalKernel, image.Point{X: 1, Y: 1}, border)
}

// SobelRGBA combines the horizontal and the vertical gradients of an RGBA image. The result is grayscale image
// which contains the high gradients ("edges") marked as white.
func SobelRGBA(img *image.RGBA, border padding.Border) (*image.Gray, error) {
	gray := grayscale.Grayscale(img)
	return SobelGray(gray, border)
}
