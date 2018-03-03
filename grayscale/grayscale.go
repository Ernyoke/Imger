package grayscale

import (
	"image"
	"image/color"
	"github.com/ernyoke/imger/utils"
)

func Grayscale(img image.Image) *image.Gray {
	gray := image.NewGray(img.Bounds())
	size := img.Bounds().Size()
	utils.ForEachPixel(size, func(x, y int) {
		gray.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
	})
	return gray
}

func Grayscale16(img image.Image) *image.Gray16 {
	gray := image.NewGray16(img.Bounds())
	size := img.Bounds().Size()
	utils.ForEachPixel(size, func(x, y int) {
		gray.Set(x, y, color.Gray16Model.Convert(img.At(x, y)))
	})
	return gray
}
