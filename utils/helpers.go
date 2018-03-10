package utils

import (
	"image"
	"image/color"
)

// ForEachPixel loops through the image and calls f functions for each [x, y] position.
func ForEachPixel(size image.Point, f func(x int, y int)) {
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			f(x, y)
		}
	}
}

// ForEachGrayPixel loops through the image and calls f functions for each gray pixel.
func ForEachGrayPixel(img *image.Gray, f func(pixel color.Gray)) {
	ForEachPixel(img.Bounds().Size(), func(x, y int) {
		pixel := img.GrayAt(x, y)
		f(pixel)
	})
}

// ForEachRGBAPixel loops through the image and calls f functions for each RGBA pixel.
func ForEachRGBAPixel(img *image.RGBA, f func(pixel color.RGBA)) {
	ForEachPixel(img.Bounds().Size(), func(x, y int) {
		pixel := img.RGBAAt(x, y)
		f(pixel)
	})
}

// ForEachRGBARedPixel loops through the image and calls f functions for red component of each RGBA pixel.
func ForEachRGBARedPixel(img *image.RGBA, f func(r uint8)) {
	ForEachRGBAPixel(img, func(pixel color.RGBA) {
		f(pixel.R)
	})
}

// ForEachRGBAGreenPixel loops through the image and calls f functions for green component of each RGBA pixel.
func ForEachRGBAGreenPixel(img *image.RGBA, f func(r uint8)) {
	ForEachRGBAPixel(img, func(pixel color.RGBA) {
		f(pixel.G)
	})
}

// ForEachRGBABluePixel loops through the image and calls f functions for blue component of each RGBA pixel.
func ForEachRGBABluePixel(img *image.RGBA, f func(r uint8)) {
	ForEachRGBAPixel(img, func(pixel color.RGBA) {
		f(pixel.B)
	})
}

// ForEachRGBAAlphaPixel loops through the image and calls f functions for alpha component of each RGBA pixel
func ForEachRGBAAlphaPixel(img *image.RGBA, f func(r uint8)) {
	ForEachRGBAPixel(img, func(pixel color.RGBA) {
		f(pixel.A)
	})
}

// ClampInt returns min if value is lesser then min, max if value is greater them max or value if the input value is
// between min and max.
func ClampInt(value int, min int, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

// ClampF64 returns min if value is lesser then min, max if value is greater them max or value if the input value is
// between min and max.
func ClampF64(value float64, min float64, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

// GetMax returns the maximum value from a slice
func GetMax(v []uint64) uint64 {
	max := v[0]
	for _, value := range v {
		if max < value {
			max = value
		}
	}
	return max
}
