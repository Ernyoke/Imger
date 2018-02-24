package blur

import (
	"image"
	"errors"
	"github.com/ernyoke/imgur/convolution"
	"github.com/ernyoke/imgur/padding"
)

func BlurGray(img *image.Gray, kernelSize image.Point, anchor image.Point) (*image.Gray, error) {
	if kernelSize.X % 2 == 0 || kernelSize.Y == 0 {
		return nil, errors.New("kernel size must contain odd numbers only")
	}
	kernel := generateKernel(&kernelSize)
	result, error := convolution.ConvolveGray(img, kernel, anchor, padding.BorderConstant)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func BlurRGBA(img *image.RGBA, kernelSize image.Point, anchor image.Point) (*image.RGBA, error) {
	if kernelSize.X % 2 == 0 || kernelSize.Y == 0 {
		return nil, errors.New("kernel size must contain odd numbers only")
	}
	kernel := generateKernel(&kernelSize)
	result, error := convolution.ConvolveRGBA(img, kernel, anchor, padding.BorderConstant)
	if error != nil {
		return nil, error
	}
	return result, nil
}

//----------------------------------------------------------------
func generateKernel(kernelSize *image.Point) *convolution.Kernel {
	kernel,_ := convolution.NewKernel(kernelSize.X, kernelSize.Y)
	for x := 0; x < kernelSize.X; x++ {
		for y := 0; y < kernelSize.Y; y++ {
			kernel.Set(x, y, 1.0 / float64(kernelSize.X * kernelSize.Y))
		}
	}
	return kernel
}
