package convolution

import (
	"image"
	"errors"
	"image/color"
	"github.com/ernyoke/imgur/padding"
)

func Convolution(img *image.Gray, kernel [][]uint8, anchor image.Point, border padding.Border) (*image.Gray, error) {
	n := len(kernel)
	if n <= 0 {
		return nil, errors.New("invalid kernel size")
	}
	m := len(kernel[0])
	kernelSize := image.Point{m, n}
	padded, error := padding.PaddingGray(img, kernelSize, anchor, border)
	if error != nil {
		return nil, error
	}
	originalSize := img.Bounds().Size()
	result := image.NewGray(img.Bounds())
	_, p, error := padding.GetPaddings(img.Bounds().Size(), kernelSize, anchor)
	for x := 0; x < originalSize.X; x++ {
		for y := 0; y < originalSize.Y; y++ {
			sum := uint32(0)
			for kx := 0; kx < kernelSize.X; kx++ {
				for ky := 0; ky < kernelSize.Y; ky++ {
					pixel := padded.At(x - p.PaddingLeft + kx, y - p.PaddingTop + ky)
					value, _, _, _ := pixel.RGBA()
					sum += value * uint32(kernel[kx][ky])
				}
			}
			result.Set(x, y, color.Gray{uint8(sum / uint32(n * m))})
		}
	}


	return nil, nil
}
