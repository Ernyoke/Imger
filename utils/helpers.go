package utils

import (
	"image"
)

func ForEchPixel(size image.Point, f func(x int, y int)) {
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			f(x, y)
		}
	}
}

func ClampInt(value int, min int, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

func ClampF64(value float64, min float64, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}
