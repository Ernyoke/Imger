package grayscale

import (
	"image"
	"image/color"

	"github.com/ernyoke/imger/utils"
)

// Grayscale takes an image on any type and returns the equivalent grayscale image represented on 8 bits.
func Grayscale(img image.Image) *image.Gray {
	gray := image.NewGray(img.Bounds())
	size := img.Bounds().Size()
	offset := img.Bounds().Min
	utils.ParallelForEachPixel(size, func(x, y int) {
		gray.Set(x+offset.X, y+offset.Y, color.GrayModel.Convert(img.At(x+offset.X, y+offset.Y)))
	})
	return gray
}

// Grayscale16 takes an image on any type and returns the equivalent grayscale image represented on 16 bits.
func Grayscale16(img image.Image) *image.Gray16 {
	gray := image.NewGray16(img.Bounds())
	size := img.Bounds().Size()
	offset := img.Bounds().Min
	utils.ParallelForEachPixel(size, func(x, y int) {
		gray.Set(x+offset.X, y+offset.Y, color.Gray16Model.Convert(img.At(x+offset.X, y+offset.Y)))
	})
	return gray
}
