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
