package convolution

import (
	"github.com/Ernyoke/Imger/padding"
	"github.com/Ernyoke/Imger/utils"
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
	utils.ParallelForEachPixel(originalSize, func(x int, y int) {
		sum := float64(0)
		for ky := 0; ky < kernelSize.Y; ky++ {
			for kx := 0; kx < kernelSize.X; kx++ {
				pixel := padded.GrayAt(x+kx, y+ky)
				kE := kernel.At(kx, ky)
				sum += float64(pixel.Y) * kE
			}
		}
		sum = utils.ClampF64(sum, utils.MinUint8, float64(utils.MaxUint8))
		resultImage.Set(x, y, color.Gray{uint8(sum)})
	})
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
	utils.ParallelForEachPixel(originalSize, func(x int, y int) {
		sumR, sumG, sumB := 0.0, 0.0, 0.0
		for kx := 0; kx < kernelSize.X; kx++ {
			for ky := 0; ky < kernelSize.Y; ky++ {
				pixel := padded.RGBAAt(x+kx, y+ky)
				sumR += float64(pixel.R) * kernel.At(kx, ky)
				sumG += float64(pixel.G) * kernel.At(kx, ky)
				sumB += float64(pixel.B) * kernel.At(kx, ky)
			}
		}
		sumR = utils.ClampF64(sumR, utils.MinUint8, float64(utils.MaxUint8))
		sumG = utils.ClampF64(sumG, utils.MinUint8, float64(utils.MaxUint8))
		sumB = utils.ClampF64(sumB, utils.MinUint8, float64(utils.MaxUint8))
		rgba := img.RGBAAt(x, y)
		resultImage.Set(x, y, color.RGBA{uint8(sumR), uint8(sumG), uint8(sumB), rgba.A})
	})
	return resultImage, nil
}
