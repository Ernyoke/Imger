package effects

import (
	"errors"
	"github.com/Ernyoke/Imger/resize"
	"github.com/Ernyoke/Imger/utils"
	"image"
	"image/color"
)

func PixelateGray(img *image.Gray, factor float64) (*image.Gray, error) {
	if factor < 1.0 {
		return nil, errors.New("invalid factor, should be greater then 1.0")
	}
	fdown := 1.0 / factor
	downScaled, downscaleError := resize.ResizeGray(img, fdown, fdown, resize.InterNearest)
	if downscaleError != nil {
		return nil, downscaleError
	}
	upscaled, upscaleError := resize.ResizeGray(downScaled, factor, factor, resize.InterNearest)
	if upscaleError != nil {
		return nil, upscaleError
	}
	return upscaled, nil
}

func PixelateRGBA(img *image.RGBA, factor float64) (*image.RGBA, error) {
	if factor < 1.0 {
		return nil, errors.New("invalid factor, should be greater then 1.0")
	}
	fdown := 1.0 / factor
	downScaled, downscaleError := resize.ResizeRGBA(img, fdown, fdown, resize.InterNearest)
	if downscaleError != nil {
		return nil, downscaleError
	}
	upscaled, upscaleError := resize.ResizeRGBA(downScaled, factor, factor, resize.InterNearest)
	if upscaleError != nil {
		return nil, upscaleError
	}
	return upscaled, nil
}

func Sepia(img *image.RGBA) *image.RGBA {
	res := image.NewRGBA(img.Rect)
	utils.ForEachPixel(img.Bounds().Size(), func(x, y int) {
		pixel := img.RGBAAt(x, y)
		r := float64(pixel.R)
		g := float64(pixel.G)
		b := float64(pixel.B)

		resR := r*0.393 + g*0.769 + b*0.189
		resG := r*0.349 + g*0.686 + b*0.168
		resB := r*0.272 + g*0.534 + b*0.131
		resPixel := color.RGBA{R: uint8(utils.ClampF64(resR, 0, 255)), G: uint8(utils.ClampF64(resG, 0, 255)), B: uint8(utils.ClampF64(resB, 0, 255)), A: pixel.A}

		res.SetRGBA(x, y, resPixel)
	})
	return res
}

func Emboss(img *image.RGBA) *image.RGBA {
	return nil
}

func Sharpen(img *image.RGBA) *image.RGBA {
	return nil
}
