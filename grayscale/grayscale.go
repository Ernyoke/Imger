package grayscale

import (
	"github.com/ernyoke/imger/utils"
	"image"
	"image/color"
)

// Grayscale takes an image on any type and returns the equivalent grayscale image represented on 8 bits.
func Grayscale(img image.Image) *image.Gray {
	gray := image.NewGray(img.Bounds())
	size := img.Bounds().Size()
	utils.ParallelForEachPixel(size, func(x, y int) {
		gray.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
	})
	return gray
}

// Grayscale16 takes an image on any type and returns the equivalent grayscale image represented on 16 bits.
func Grayscale16(img image.Image) *image.Gray16 {
	gray := image.NewGray16(img.Bounds())
	size := img.Bounds().Size()
	utils.ParallelForEachPixel(size, func(x, y int) {
		gray.Set(x, y, color.Gray16Model.Convert(img.At(x, y)))
	})
	return gray
}
