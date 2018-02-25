package grayscale

import (
	"image"
	"image/color"
)

func Grayscale(img image.Image) *image.Gray {
	gray := image.NewGray(img.Bounds())
	size := img.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			gray.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}
	return gray
}

func Grayscale16(img image.Image) *image.Gray16 {
	gray := image.NewGray16(img.Bounds())
	size := img.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			gray.Set(x, y, color.Gray16Model.Convert(img.At(x, y)))
		}
	}
	return gray
}
