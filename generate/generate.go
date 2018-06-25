package generate

import (
	"image"
	"image/color"
	"math"
)

// Direction constant -  direction of the gradient from first color to the second color.
type Direction int

const (
	// H - horizontal direction
	H Direction = iota
	// V - vertical direction
	V
)

func normalize(value float64, min float64, max float64) float64 {
	lower := -6.0
	upper := 6.0
	norm := (value - min) / (max - min)
	return norm*(upper-lower) + lower
}

// LinearGradient generates a gradient image using a linear function.
func LinearGradient(size image.Point, startColor color.RGBA, endColor color.RGBA, direction Direction) *image.RGBA {
	gradFunc := func(colorChannel uint8, percent float64) uint8 {
		return uint8(math.Floor(float64(colorChannel) * percent))
	}
	res := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
	switch direction {
	case V:
		step := 1.0 / float64(size.Y)
		percent := 0.0
		for y := 0; y < size.Y; y++ {
			c := color.RGBA{
				R: gradFunc(startColor.R, 1.0-percent) + gradFunc(endColor.R, percent),
				G: gradFunc(startColor.G, 1.0-percent) + gradFunc(endColor.G, percent),
				B: gradFunc(startColor.B, 1.0-percent) + gradFunc(endColor.B, percent),
				A: 255,
			}
			percent += step
			for x := 0; x < size.X; x++ {
				res.SetRGBA(x, y, c)
			}
		}
	case H:
		step := 1.0 / float64(size.X)
		percent := 0.0
		for x := 0; x < size.X; x++ {
			c := color.RGBA{
				R: gradFunc(startColor.R, -percent) + gradFunc(endColor.R, percent),
				G: gradFunc(startColor.G, -percent) + gradFunc(endColor.G, percent),
				B: gradFunc(startColor.B, -percent) + gradFunc(endColor.B, percent),
				A: 255,
			}
			percent += step
			for y := 0; y < size.Y; y++ {
				res.SetRGBA(x, y, c)
			}
		}
	}
	return res
}

// SigmoidalGradient generates a gradient image using the sigmoid ( f(x) = 1 / (1 + exp(-x)) )  function.
func SigmoidalGradient(size image.Point, startColor color.RGBA, endColor color.RGBA, direction Direction) *image.RGBA {
	sigmoid := func(val float64) float64 {
		return 1.0 / (1.0 + math.Exp(-val))
	}
	res := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
	switch direction {
	case V:
		for y := 0; y < size.Y; y++ {
			percent := sigmoid(normalize(float64(y), 0, float64(size.Y)))
			c := color.RGBA{
				R: uint8((1.0-percent)*float64(startColor.R)) + uint8(percent*float64(endColor.R)),
				G: uint8((1.0-percent)*float64(startColor.G)) + uint8(percent*float64(endColor.G)),
				B: uint8((1.0-percent)*float64(startColor.B)) + uint8(percent*float64(endColor.B)),
				A: 255,
			}
			for x := 0; x < size.X; x++ {
				res.SetRGBA(x, y, c)
			}
		}
	case H:
		for x := 0; x < size.X; x++ {
			percent := sigmoid(normalize(float64(x), 0, float64(size.X)))
			c := color.RGBA{
				R: uint8((1.0-percent)*float64(startColor.R)) + uint8(percent*float64(endColor.R)),
				G: uint8((1.0-percent)*float64(startColor.G)) + uint8(percent*float64(endColor.G)),
				B: uint8((1.0-percent)*float64(startColor.B)) + uint8(percent*float64(endColor.B)),
				A: 255,
			}
			for y := 0; y < size.Y; y++ {
				res.SetRGBA(x, y, c)
			}
		}
	}
	return res
}
