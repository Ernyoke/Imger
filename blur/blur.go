package blur

import (
	"image"
	"errors"
)

func generateKernel(kernelSize *image.Point) [][]uint8 {
	var kernel [][]uint8
	for x := 0; x < kernelSize.X; x++ {
		var row []uint8
		for y := 0; y < kernelSize.Y; y++ {
			row = append(row, 1)
		}
		kernel = append(kernel, row)
	}
	return kernel
}

func Blur(img image.Image, kernelSize image.Point, anchor image.Point) (*image.Image, error) {
	if kernelSize.X %2 == 0 || kernelSize.Y == 0 {
		return nil, errors.New("kernel size must contain odd numbers only")
	}
	//kernel := generateKernel(&kernelSize)


	return nil, nil
}
