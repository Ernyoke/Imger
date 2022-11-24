package effects

import (
	"errors"
	"github.com/ernyoke/imger/blend"
	"github.com/ernyoke/imger/convolution"
	"github.com/ernyoke/imger/grayscale"
	"github.com/ernyoke/imger/padding"
	"github.com/ernyoke/imger/resize"
	"github.com/ernyoke/imger/utils"
	"image"
	"image/color"
)

var sharpenKernel = convolution.Kernel{Content: [][]float64{
	{0, -1, 0},
	{-1, 5, -1},
	{0, -1, 0},
}, Width: 3, Height: 3}

// PixelateGray enlarges the pixels of a grayscale image. The factor value specifies how much should be the pixels
// enlarged.
// Example of usage:
//
//		 res, err := effects.PixelateGray(img, 5.0)
//
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

// PixelateRGBA enlarges the pixels of a RGBA image. The factor value specifies how much should be the pixels enlarged.
// Example of usage:
//
//		 res, err := effects.PixelateRGBA(img, 5.0)
//
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

// Sepia applies Sepia tone to an RGBA image.
func Sepia(img *image.RGBA) *image.RGBA {
	res := image.NewRGBA(img.Rect)
	utils.ParallelForEachPixel(img.Bounds().Size(), func(x, y int) {
		pixel := img.RGBAAt(x, y)
		r := float64(pixel.R)
		g := float64(pixel.G)
		b := float64(pixel.B)

		resR := r*0.393 + g*0.769 + b*0.189
		resG := r*0.349 + g*0.686 + b*0.168
		resB := r*0.272 + g*0.534 + b*0.131
		resPixel := color.RGBA{R: uint8(utils.ClampF64(resR, utils.MinUint8, float64(utils.MaxUint8))),
			G: uint8(utils.ClampF64(resG, utils.MinUint8, float64(utils.MaxUint8))),
			B: uint8(utils.ClampF64(resB, utils.MinUint8, float64(utils.MaxUint8))), A: pixel.A}

		res.SetRGBA(x, y, resPixel)
	})
	return res
}

// EmbossGray takes a grayscale image and returns a copy of the image in which each pixel has been replaced either by a
// highlight or a shadow representation.
func EmbossGray(img *image.Gray) (*image.Gray, error) {
	var kernel = convolution.Kernel{Content: [][]float64{
		{-1, -1, 0},
		{-1, 0, 1},
		{0, 1, 1},
	}, Width: 3, Height: 3}

	conv, err := convolution.ConvolveGray(img, &kernel, image.Point{X: 1, Y: 1}, padding.BorderReflect)
	if err != nil {
		return nil, err
	}
	return blend.AddScalarToGray(conv, 128), nil
}

// EmbossRGBA takes an RGBA image and returns a grayscale image in which each pixel has been replaced either by a
// highlight or a shadow representation.
func EmbossRGBA(img *image.RGBA) (*image.Gray, error) {
	gray := grayscale.Grayscale(img)
	return EmbossGray(gray)
}

// SharpenGray takes a grayscale image and returns another grayscale image where each edge is added to the original
// image.
func SharpenGray(img *image.Gray) (*image.Gray, error) {
	return convolution.ConvolveGray(img, &sharpenKernel, image.Point{X: 1, Y: 1}, padding.BorderReflect)
}

// SharpenRGBA takes an RGBA image and returns another RGBA image where each edge is added to the original image.
func SharpenRGBA(img *image.RGBA) (*image.RGBA, error) {
	return convolution.ConvolveRGBA(img, &sharpenKernel, image.Point{X: 1, Y: 1}, padding.BorderReflect)
}

// InvertGray takes a grayscale image and return its inverted grayscale image.
func InvertGray(img *image.Gray) *image.Gray {
	size := img.Bounds().Size()
	inverted := image.NewGray(img.Rect)
	utils.ParallelForEachPixel(size, func(x, y int) {
		original := img.GrayAt(x, y).Y
		inverted.SetGray(x, y, color.Gray{Y: utils.MaxUint8 - original})
	})
	return inverted
}

// InvertRGBA takes an RGBA image and return its inverted RGBA image.
func InvertRGBA(img *image.RGBA) *image.RGBA {
	size := img.Bounds().Size()
	inverted := image.NewRGBA(img.Rect)
	utils.ParallelForEachPixel(size, func(x, y int) {
		originalColor := img.RGBAAt(x, y)
		invertedColor := color.RGBA{R: utils.MaxUint8 - originalColor.R,
			G: utils.MaxUint8 - originalColor.G,
			B: utils.MaxUint8 - originalColor.B,
			A: originalColor.A}
		inverted.SetRGBA(x, y, invertedColor)
	})
	return inverted
}
