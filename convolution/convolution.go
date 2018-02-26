package convolution

import (
	"github.com/ernyoke/imger/padding"
	"image"
	"image/color"
)

func ConvolveGray(img *image.Gray, kernel *Kernel, anchor image.Point, border padding.Border) (*image.Gray, error) {
	kernelSize := kernel.Size()
	padded, error := padding.PaddingGray(img, kernelSize, anchor, border)
	if error != nil {
		return nil, error
	}
	originalSize := img.Bounds().Size()
	resultImage := image.NewGray(img.Bounds())
	for y := 0; y < originalSize.Y; y++ {
		for x := 0; x < originalSize.X; x++ {
			sum := float64(0)
			for ky := 0; ky < kernelSize.Y; ky++ {
				for kx := 0; kx < kernelSize.X; kx++ {
					pixel := padded.GrayAt(x+kx, y+ky)
					kE := kernel.At(kx, ky)
					sum += float64(pixel.Y) * kE
				}
			}
			if sum < 0 {
				sum = 0
			}
			if sum > 255 {
				sum = 255
			}
			resultImage.Set(x, y, color.Gray{uint8(sum)})
		}
	}
	return resultImage, nil
}

func ConvolveRGBA(img *image.RGBA, kernel *Kernel, anchor image.Point, border padding.Border) (*image.RGBA, error) {
	kernelSize := kernel.Size()
	padded, error := padding.PaddingRGBA(img, kernelSize, anchor, border)
	if error != nil {
		return nil, error
	}
	originalSize := img.Bounds().Size()
	resultImage := image.NewRGBA(img.Bounds())
	for x := 0; x < originalSize.X; x++ {
		for y := 0; y < originalSize.Y; y++ {
			sumR := float64(0)
			sumG := float64(0)
			sumB := float64(0)
			for kx := 0; kx < kernelSize.X; kx++ {
				for ky := 0; ky < kernelSize.Y; ky++ {
					pixel := padded.RGBAAt(x+kx, y+ky)
					sumR += float64(pixel.R) * kernel.At(kx, ky)
					sumG += float64(pixel.G) * kernel.At(kx, ky)
					sumB += float64(pixel.B) * kernel.At(kx, ky)
				}
			}
			rgba := img.RGBAAt(x, y)
			resultImage.Set(x, y, color.RGBA{uint8(sumR), uint8(sumG), uint8(sumB), rgba.A})
		}
	}
	return resultImage, nil
}
