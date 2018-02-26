package blur

import (
	"errors"
	"github.com/ernyoke/imger/convolution"
	"github.com/ernyoke/imger/padding"
	"image"
	"math"
)

func BoxGray(img *image.Gray, kernelSize image.Point, anchor image.Point, border padding.Border) (*image.Gray, error) {
	kernel := generateBoxKernel(&kernelSize)
	result, error := convolution.ConvolveGray(img, kernel.Normalize(), anchor, border)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func BoxRGBA(img *image.RGBA, kernelSize image.Point, anchor image.Point, border padding.Border) (*image.RGBA, error) {
	kernel := generateBoxKernel(&kernelSize)
	result, error := convolution.ConvolveRGBA(img, kernel.Normalize(), anchor, border)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func GaussianBlurGray(img *image.Gray, radius float64, sigma float64, border padding.Border) (*image.Gray, error) {
	if radius <= 0 {
		return nil, errors.New("radius must be bigger then 0")
	}
	return convolution.ConvolveGray(img, generateGaussianKernel(radius, sigma).Normalize(), image.Point{X: int(math.Ceil(radius)), Y: int(math.Ceil(radius))}, border)
}

func GaussianBlurRGBA(img *image.RGBA, radius float64, sigma float64, border padding.Border) (*image.RGBA, error) {
	if radius <= 0 {
		return nil, errors.New("radius must be bigger then 0")
	}
	return convolution.ConvolveRGBA(img, generateGaussianKernel(radius, sigma).Normalize(), image.Point{X: int(math.Ceil(radius)), Y: int(math.Ceil(radius))}, border)
}

// -------------------------------------------------------------------------------------------------------
func generateBoxKernel(kernelSize *image.Point) *convolution.Kernel {
	kernel, _ := convolution.NewKernel(kernelSize.X, kernelSize.Y)
	for x := 0; x < kernelSize.X; x++ {
		for y := 0; y < kernelSize.Y; y++ {
			kernel.Set(x, y, 1.0/float64(kernelSize.X*kernelSize.Y))
		}
	}
	return kernel
}

func generateGaussianKernel(radius float64, sigma float64) *convolution.Kernel {
	length := int(math.Ceil(2*radius + 1))
	kernel, _ := convolution.NewKernel(length, length)
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			kernel.Set(x, y, gaussianFunc(float64(x)-radius, float64(y)-radius, sigma))
		}
	}
	return kernel
}

func gaussianFunc(x, y, sigma float64) float64 {
	sigSqr := sigma * sigma
	return (1.0 / (2 * math.Pi * sigSqr)) * math.Exp(-(x*x+y*y)/(2*sigSqr))
}
